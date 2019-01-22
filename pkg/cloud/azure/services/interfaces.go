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

package services

import (
	providerv1 "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1alpha1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
)

/*
// Getter is a unified interfaces that includes all the getters.
type Getter interface {
	SDKSessionGetter
	AzureGetter
	LBGetter
}

// SDKSessionGetter has a single method that returns an Azure session.
type SDKSessionGetter interface {
	Session(*providerv1.AzureClusterProviderSpec) *session.Session
}

// AzureGetter has a single method that returns an Azure service interface.
type AzureGetter interface {
	Azure(*session.Session) AzureInterface
}

// LBGetter has a single method that returns a LB service interface.
type LBGetter interface {
	LB(*session.Session) LBInterface
}
*/

// AzureInterface encapsulates the methods exposed by the Azure service.
type AzureInterface interface {
	AzureClusterInterface
	AzureMachineInterface
}

// AzureClusterInterface encapsulates the methods exposed to the cluster
// actuator
type AzureClusterInterface interface {
	ReconcileNetwork() error
	ReconcileBastion() error
	DeleteNetwork() error
	DeleteBastion() error
}

// AzureMachineInterface encapsulates the methods exposed to the machine
// actuator
type AzureMachineInterface interface {
	InstanceIfExists(id string) (*providerv1.Instance, error)
	TerminateInstance(id string) error
	CreateOrGetMachine(machine *actuators.MachineScope, token string) (*providerv1.Instance, error)
	UpdateInstanceSecurityGroups(id string, securityGroups []string) error
	UpdateResourceTags(resourceID *string, create map[string]string, remove map[string]string) error
}

// LBInterface encapsulates the methods exposed by the elb service.
type LBInterface interface {
	ReconcileLoadBalancers() error
	DeleteLoadBalancers() error
	RegisterInstanceWithAPIServerLB(vmID string) error
	GetAPIServerDNSName() (string, error)
}

/*
// interface for all azure services clients
type AzureClients struct {
	Compute            AzureComputeClient
	Network            AzureNetworkClient
	Resourcemanagement AzureResourceManagementClient
}

// AzureComputeClient defines the operations that will interact with the Azure compute API
type AzureComputeClient interface {
	// Virtual Machines Operations
	RunCommand(resoureGroup string, name string, cmd string) (compute.VirtualMachinesRunCommandFuture, error)
	VMIfExists(resourceGroup string, name string) (*compute.VirtualMachine, error)
	DeleteVM(resourceGroup string, name string) (compute.VirtualMachinesDeleteFuture, error)
	WaitForVMRunCommandFuture(future compute.VirtualMachinesRunCommandFuture) error
	WaitForVMDeletionFuture(future compute.VirtualMachinesDeleteFuture) error

	// Disk Operations
	DeleteManagedDisk(resourceGroup string, name string) (compute.DisksDeleteFuture, error)
	WaitForDisksDeleteFuture(future compute.DisksDeleteFuture) error
}

// AzureNetworkClient defines the operations that will interact with the Azure networking API
type AzureNetworkClient interface {
	// Network Interfaces Operations
	DeleteNetworkInterface(resourceGroupName string, networkInterfaceName string) (network.InterfacesDeleteFuture, error)
	WaitForNetworkInterfacesDeleteFuture(future network.InterfacesDeleteFuture) error

	// Network Security Groups Operations
	CreateOrUpdateNetworkSecurityGroup(resourceGroupName string, networkSecurityGroupName string, location string) (*network.SecurityGroupsCreateOrUpdateFuture, error)
	NetworkSGIfExists(resourceGroupName string, networkSecurityGroupName string) (*network.SecurityGroup, error)
	WaitForNetworkSGsCreateOrUpdateFuture(future network.SecurityGroupsCreateOrUpdateFuture) error

	// Public Ip Address Operations
	GetPublicIPAddress(resourceGroupName string, IPName string) (network.PublicIPAddress, error)
	DeletePublicIPAddress(resourceGroup string, IPName string) (network.PublicIPAddressesDeleteFuture, error)
	WaitForPublicIPAddressDeleteFuture(future network.PublicIPAddressesDeleteFuture) error

	// Virtual Networks Operations
	CreateOrUpdateVnet(resourceGroupName string, virtualNetworkName string, location string) (*network.VirtualNetworksCreateOrUpdateFuture, error)
	WaitForVnetCreateOrUpdateFuture(future network.VirtualNetworksCreateOrUpdateFuture) error
}

// AzureResourceManagementClient defines the operations that will interact with the Azure resources API
type AzureResourceManagementClient interface {
	// Resource Groups Operations
	CreateOrUpdateGroup(resourceGroupName string, location string) (resources.Group, error)
	DeleteGroup(resourceGroupName string) (resources.GroupsDeleteFuture, error)
	CheckGroupExistence(rgName string) (autorest.Response, error)
	WaitForGroupsDeleteFuture(future resources.GroupsDeleteFuture) error

	// Deployment Operations
	CreateOrUpdateDeployment(machine *clusterv1.Machine, clusterConfig *azureconfigv1.AzureClusterProviderSpec, machineConfig *azureconfigv1.AzureMachineProviderSpec) (*resources.DeploymentsCreateOrUpdateFuture, error)
	GetDeploymentResult(future resources.DeploymentsCreateOrUpdateFuture) (de resources.DeploymentExtended, err error)
	ValidateDeployment(machine *clusterv1.Machine, clusterConfig *azureconfigv1.AzureClusterProviderSpec, machineConfig *azureconfigv1.AzureMachineProviderSpec) error
	WaitForDeploymentsCreateOrUpdateFuture(future resources.DeploymentsCreateOrUpdateFuture) error
}
*/
