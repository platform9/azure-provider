/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scope

import (
	"context"
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"

	"k8s.io/utils/net"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"k8s.io/klog/klogr"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	azure "sigs.k8s.io/cluster-api-provider-azure/cloud"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClusterScopeParams defines the input parameters used to create a new Scope.
type ClusterScopeParams struct {
	AzureClients
	Client       client.Client
	Logger       logr.Logger
	Cluster      *clusterv1.Cluster
	AzureCluster *infrav1.AzureCluster
}

// NewClusterScope creates a new Scope from the supplied parameters.
// This is meant to be called for each reconcile iteration.
func NewClusterScope(params ClusterScopeParams) (*ClusterScope, error) {
	if params.Cluster == nil {
		return nil, errors.New("failed to generate new scope from nil Cluster")
	}
	if params.AzureCluster == nil {
		return nil, errors.New("failed to generate new scope from nil AzureCluster")
	}

	if params.Logger == nil {
		params.Logger = klogr.New()
	}

	err := params.AzureClients.setCredentials(params.AzureCluster.Spec.SubscriptionID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to configure azure settings and credentials from environment")
	}

	helper, err := patch.NewHelper(params.AzureCluster, params.Client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init patch helper")
	}

	return &ClusterScope{
		Logger:       params.Logger,
		Client:       params.Client,
		AzureClients: params.AzureClients,
		Cluster:      params.Cluster,
		AzureCluster: params.AzureCluster,
		patchHelper:  helper,
	}, nil
}

// ClusterScope defines the basic context for an actuator to operate upon.
type ClusterScope struct {
	logr.Logger
	Client      client.Client
	patchHelper *patch.Helper

	AzureClients
	Cluster      *clusterv1.Cluster
	AzureCluster *infrav1.AzureCluster
}

// BaseURI returns the Azure ResourceManagerEndpoint.
func (s *ClusterScope) BaseURI() string {
	return s.ResourceManagerEndpoint
}

// Authorizer returns the Azure client Authorizer.
func (s *ClusterScope) Authorizer() autorest.Authorizer {
	return s.AzureClients.Authorizer
}

// PublicIPSpecs returns the public IP specs.
func (s *ClusterScope) PublicIPSpecs() []azure.PublicIPSpec {
	specs := []azure.PublicIPSpec{
		{
			Name: azure.GenerateNodeOutboundIPName(s.NodeOutboundLBName()),
		},
	}
	if s.APIServerLB().Type == infrav1.Public {
		specs = append(specs, azure.PublicIPSpec{
			Name:    s.APIServerPublicIP().Name,
			DNSName: s.APIServerPublicIP().DNSName,
			IsIPv6:  false, // currently azure requires a ipv4 lb rule to enable ipv6
		})
	}
	return specs
}

// LBSpecs returns the load balancer specs.
func (s *ClusterScope) LBSpecs() []azure.LBSpec {
	return []azure.LBSpec{
		{
			// Control Plane LB
			Name:              s.APIServerLB().Name,
			SubnetName:        s.ControlPlaneSubnet().Name,
			SubnetCidrs:       s.ControlPlaneSubnet().CIDRBlocks,
			FrontendIPConfigs: s.APIServerLB().FrontendIPs,
			APIServerPort:     s.APIServerPort(),
			Type:              s.APIServerLB().Type,
			Role:              infrav1.APIServerRole,
			BackendPoolName:   azure.GenerateBackendAddressPoolName(s.APIServerLB().Name),
		},
		{
			// Public Node outbound LB
			Name: s.NodeOutboundLBName(),
			FrontendIPConfigs: []infrav1.FrontendIP{
				{
					PublicIP: &infrav1.PublicIPSpec{
						Name: azure.GenerateNodeOutboundIPName(s.NodeOutboundLBName()),
					},
				},
			},
			Type:            infrav1.Public,
			BackendPoolName: azure.GenerateOutboundBackendddressPoolName(s.NodeOutboundLBName()),
			Role:            infrav1.NodeOutboundRole,
		},
	}
}

// RouteTableSpecs returns the node route table(s)
func (s *ClusterScope) RouteTableSpecs() []azure.RouteTableSpec {
	return []azure.RouteTableSpec{{
		Name: s.RouteTable().Name,
	}}
}

// NSGSpecs returns the security group specs.
func (s *ClusterScope) NSGSpecs() []azure.NSGSpec {
	return []azure.NSGSpec{
		{
			Name:         s.ControlPlaneSubnet().SecurityGroup.Name,
			IngressRules: s.ControlPlaneSubnet().SecurityGroup.IngressRules,
		},
		{
			Name:         s.NodeSubnet().SecurityGroup.Name,
			IngressRules: s.NodeSubnet().SecurityGroup.IngressRules,
		},
	}
}

// SubnetSpecs returns the subnets specs.
func (s *ClusterScope) SubnetSpecs() []azure.SubnetSpec {
	return []azure.SubnetSpec{
		{
			Name:              s.ControlPlaneSubnet().Name,
			CIDRs:             s.ControlPlaneSubnet().CIDRBlocks,
			VNetName:          s.Vnet().Name,
			SecurityGroupName: s.ControlPlaneSubnet().SecurityGroup.Name,
			Role:              s.ControlPlaneSubnet().Role,
			RouteTableName:    s.ControlPlaneSubnet().RouteTable.Name,
		},
		{
			Name:              s.NodeSubnet().Name,
			CIDRs:             s.NodeSubnet().CIDRBlocks,
			VNetName:          s.Vnet().Name,
			SecurityGroupName: s.NodeSubnet().SecurityGroup.Name,
			RouteTableName:    s.NodeSubnet().RouteTable.Name,
			Role:              s.NodeSubnet().Role,
		},
	}
}

/// VNetSpecs returns the virtual network specs.
func (s *ClusterScope) VNetSpecs() []azure.VNetSpec {
	return []azure.VNetSpec{
		{
			ResourceGroup: s.Vnet().ResourceGroup,
			Name:          s.Vnet().Name,
			CIDRs:         s.Vnet().CIDRBlocks,
		},
	}
}

// Vnet returns the cluster Vnet.
func (s *ClusterScope) Vnet() *infrav1.VnetSpec {
	return &s.AzureCluster.Spec.NetworkSpec.Vnet
}

// IsVnetManaged returns true if the vnet is managed.
func (s *ClusterScope) IsVnetManaged() bool {
	return s.Vnet().ID == "" || s.Vnet().Tags.HasOwned(s.ClusterName())
}

// IsIPv6Enabled returns true if IPv6 is enabled.
func (s *ClusterScope) IsIPv6Enabled() bool {
	for _, cidr := range s.AzureCluster.Spec.NetworkSpec.Vnet.CIDRBlocks {
		if net.IsIPv6CIDRString(cidr) {
			return true
		}
	}
	return false
}

// Subnets returns the cluster subnets.
func (s *ClusterScope) Subnets() infrav1.Subnets {
	return s.AzureCluster.Spec.NetworkSpec.Subnets
}

// ControlPlaneSubnet returns the cluster control plane subnet.
func (s *ClusterScope) ControlPlaneSubnet() *infrav1.SubnetSpec {
	return s.AzureCluster.Spec.NetworkSpec.GetControlPlaneSubnet()
}

// NodeSubnet returns the cluster node subnet.
func (s *ClusterScope) NodeSubnet() *infrav1.SubnetSpec {
	return s.AzureCluster.Spec.NetworkSpec.GetNodeSubnet()
}

// APIServerLB returns the cluster API Server load balancer.
func (s *ClusterScope) APIServerLB() *infrav1.LoadBalancerSpec {
	return &s.AzureCluster.Spec.NetworkSpec.APIServerLB
}

// RouteTable returns the cluster node routetable.
func (s *ClusterScope) RouteTable() *infrav1.RouteTable {
	return &s.AzureCluster.Spec.NetworkSpec.GetNodeSubnet().RouteTable
}

// APIServerLBName returns the API Server LB name.
func (s *ClusterScope) APIServerLBName() string {
	return s.APIServerLB().Name
}

// APIServerPublicIP returns the API Server public IP.
func (s *ClusterScope) APIServerPublicIP() *infrav1.PublicIPSpec {
	return s.APIServerLB().FrontendIPs[0].PublicIP
}

// APIServerPrivateIP returns the API Server private IP.
func (s *ClusterScope) APIServerPrivateIP() string {
	return s.APIServerLB().FrontendIPs[0].PrivateIPAddress
}

// NodeOutboundLBName returns the name of the node outbound LB.
func (s *ClusterScope) NodeOutboundLBName() string {
	return s.ClusterName()
}

// ResourceGroup returns the cluster resource group.
func (s *ClusterScope) ResourceGroup() string {
	return s.AzureCluster.Spec.ResourceGroup
}

// ClusterName returns the cluster name.
func (s *ClusterScope) ClusterName() string {
	return s.Cluster.Name
}

// Namespace returns the cluster namespace.
func (s *ClusterScope) Namespace() string {
	return s.Cluster.Namespace
}

// Location returns the cluster location.
func (s *ClusterScope) Location() string {
	return s.AzureCluster.Spec.Location
}

// GenerateFQDN generates a fully qualified domain name, based on a hash, cluster name and cluster location.
func (s *ClusterScope) GenerateFQDN(ipName string) string {
	h := fnv.New32a()
	if _, err := h.Write([]byte(fmt.Sprintf("%s/%s/%s", s.SubscriptionID(), s.ResourceGroup(), ipName))); err != nil {
		return ""
	}
	hash := fmt.Sprintf("%x", h.Sum32())
	return strings.ToLower(fmt.Sprintf("%s-%s.%s.%s", s.ClusterName(), hash, s.Location(), s.AzureClients.ResourceManagerVMDNSSuffix))
}

// ListOptionsLabelSelector returns a ListOptions with a label selector for clusterName.
func (s *ClusterScope) ListOptionsLabelSelector() client.ListOption {
	return client.MatchingLabels(map[string]string{
		clusterv1.ClusterLabelName: s.Cluster.Name,
	})
}

// PatchObject persists the cluster configuration and status.
func (s *ClusterScope) PatchObject(ctx context.Context) error {
	return s.patchHelper.Patch(
		ctx,
		s.AzureCluster,
		patch.WithOwnedConditions{Conditions: []clusterv1.ConditionType{
			clusterv1.ReadyCondition,
			infrav1.NetworkInfrastructureReadyCondition,
		}})
}

// Close closes the current scope persisting the cluster configuration and status.
func (s *ClusterScope) Close(ctx context.Context) error {
	return s.patchHelper.Patch(ctx, s.AzureCluster)
}

// AdditionalTags returns AdditionalTags from the scope's AzureCluster.
func (s *ClusterScope) AdditionalTags() infrav1.Tags {
	tags := make(infrav1.Tags)
	if s.AzureCluster.Spec.AdditionalTags != nil {
		tags = s.AzureCluster.Spec.AdditionalTags.DeepCopy()
	}
	return tags
}

// APIServerPort returns the APIServerPort to use when creating the load balancer.
func (s *ClusterScope) APIServerPort() int32 {
	if s.Cluster.Spec.ClusterNetwork != nil && s.Cluster.Spec.ClusterNetwork.APIServerPort != nil {
		return *s.Cluster.Spec.ClusterNetwork.APIServerPort
	}
	return 6443
}

// APIServerPort returns the APIServerPort to use when creating the load balancer.
func (s *ClusterScope) APIServerHost() string {
	if s.APIServerLB().Type == infrav1.Internal {
		return s.APIServerPrivateIP()
	}
	return s.APIServerPublicIP().DNSName
}

// SetFailureDomain will set the spec for a for a given key
func (s *ClusterScope) SetFailureDomain(id string, spec clusterv1.FailureDomainSpec) {
	if s.AzureCluster.Status.FailureDomains == nil {
		s.AzureCluster.Status.FailureDomains = make(clusterv1.FailureDomains, 0)
	}
	s.AzureCluster.Status.FailureDomains[id] = spec
}

func (s *ClusterScope) SetControlPlaneIngressRules() {
	if s.ControlPlaneSubnet().SecurityGroup.IngressRules == nil {
		s.ControlPlaneSubnet().SecurityGroup.IngressRules = infrav1.IngressRules{
			&infrav1.IngressRule{
				Name:             "allow_ssh",
				Description:      "Allow SSH",
				Priority:         2200,
				Protocol:         infrav1.SecurityGroupProtocolTCP,
				Source:           to.StringPtr("*"),
				SourcePorts:      to.StringPtr("*"),
				Destination:      to.StringPtr("*"),
				DestinationPorts: to.StringPtr("22"),
			},
			&infrav1.IngressRule{
				Name:             "allow_apiserver",
				Description:      "Allow K8s API Server",
				Priority:         2201,
				Protocol:         infrav1.SecurityGroupProtocolTCP,
				Source:           to.StringPtr("*"),
				SourcePorts:      to.StringPtr("*"),
				Destination:      to.StringPtr("*"),
				DestinationPorts: to.StringPtr(strconv.Itoa(int(s.APIServerPort()))),
			},
		}
	}
}

// SetDNSName sets the API Server public IP DNS name.
func (s *ClusterScope) SetDNSName() {
	if s.APIServerLB().Type == infrav1.Public {
		if s.APIServerPublicIP().DNSName == "" {
			s.APIServerPublicIP().DNSName = s.GenerateFQDN(s.APIServerPublicIP().Name)
		}
	}
}
