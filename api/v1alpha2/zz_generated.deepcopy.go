// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIEndpoint) DeepCopyInto(out *APIEndpoint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIEndpoint.
func (in *APIEndpoint) DeepCopy() *APIEndpoint {
	if in == nil {
		return nil
	}
	out := new(APIEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildParams) DeepCopyInto(out *BuildParams) {
	*out = *in
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.Additional != nil {
		in, out := &in.Additional, &out.Additional
		*out = make(Labels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildParams.
func (in *BuildParams) DeepCopy() *BuildParams {
	if in == nil {
		return nil
	}
	out := new(BuildParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCluster) DeepCopyInto(out *AzureCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCluster.
func (in *AzureCluster) DeepCopy() *AzureCluster {
	if in == nil {
		return nil
	}
	out := new(AzureCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterList) DeepCopyInto(out *AzureClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterList.
func (in *AzureClusterList) DeepCopy() *AzureClusterList {
	if in == nil {
		return nil
	}
	out := new(AzureClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterSpec) DeepCopyInto(out *AzureClusterSpec) {
	*out = *in
	in.NetworkSpec.DeepCopyInto(&out.NetworkSpec)
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(Labels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterSpec.
func (in *AzureClusterSpec) DeepCopy() *AzureClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AzureClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureClusterStatus) DeepCopyInto(out *AzureClusterStatus) {
	*out = *in
	in.Network.DeepCopyInto(&out.Network)
	if in.APIEndpoints != nil {
		in, out := &in.APIEndpoints, &out.APIEndpoints
		*out = make([]APIEndpoint, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureClusterStatus.
func (in *AzureClusterStatus) DeepCopy() *AzureClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AzureClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachine) DeepCopyInto(out *AzureMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachine.
func (in *AzureMachine) DeepCopy() *AzureMachine {
	if in == nil {
		return nil
	}
	out := new(AzureMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineList) DeepCopyInto(out *AzureMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineList.
func (in *AzureMachineList) DeepCopy() *AzureMachineList {
	if in == nil {
		return nil
	}
	out := new(AzureMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineSpec) DeepCopyInto(out *AzureMachineSpec) {
	*out = *in
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(string)
		**out = **in
	}
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.ImageFamily != nil {
		in, out := &in.ImageFamily, &out.ImageFamily
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(Labels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(bool)
		**out = **in
	}
	if in.AdditionalNetworkTags != nil {
		in, out := &in.AdditionalNetworkTags, &out.AdditionalNetworkTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccount)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineSpec.
func (in *AzureMachineSpec) DeepCopy() *AzureMachineSpec {
	if in == nil {
		return nil
	}
	out := new(AzureMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineStatus) DeepCopyInto(out *AzureMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.InstanceStatus != nil {
		in, out := &in.InstanceStatus, &out.InstanceStatus
		*out = new(InstanceStatus)
		**out = **in
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineStatus.
func (in *AzureMachineStatus) DeepCopy() *AzureMachineStatus {
	if in == nil {
		return nil
	}
	out := new(AzureMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplate) DeepCopyInto(out *AzureMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplate.
func (in *AzureMachineTemplate) DeepCopy() *AzureMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplateList) DeepCopyInto(out *AzureMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplateList.
func (in *AzureMachineTemplateList) DeepCopy() *AzureMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplateResource) DeepCopyInto(out *AzureMachineTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplateResource.
func (in *AzureMachineTemplateResource) DeepCopy() *AzureMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineTemplateSpec) DeepCopyInto(out *AzureMachineTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineTemplateSpec.
func (in *AzureMachineTemplateSpec) DeepCopy() *AzureMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(AzureMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Labels) DeepCopyInto(out *Labels) {
	{
		in := &in
		*out = make(Labels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Labels.
func (in Labels) DeepCopy() Labels {
	if in == nil {
		return nil
	}
	out := new(Labels)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.FirewallRules != nil {
		in, out := &in.FirewallRules, &out.FirewallRules
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIServerAddress != nil {
		in, out := &in.APIServerAddress, &out.APIServerAddress
		*out = new(string)
		**out = **in
	}
	if in.APIServerHealthCheck != nil {
		in, out := &in.APIServerHealthCheck, &out.APIServerHealthCheck
		*out = new(string)
		**out = **in
	}
	if in.APIServerInstanceGroups != nil {
		in, out := &in.APIServerInstanceGroups, &out.APIServerInstanceGroups
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.APIServerBackendService != nil {
		in, out := &in.APIServerBackendService, &out.APIServerBackendService
		*out = new(string)
		**out = **in
	}
	if in.APIServerTargetProxy != nil {
		in, out := &in.APIServerTargetProxy, &out.APIServerTargetProxy
		*out = new(string)
		**out = **in
	}
	if in.APIServerForwardingRule != nil {
		in, out := &in.APIServerForwardingRule, &out.APIServerForwardingRule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(Subnets, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SubnetSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccount) DeepCopyInto(out *ServiceAccount) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccount.
func (in *ServiceAccount) DeepCopy() *ServiceAccount {
	if in == nil {
		return nil
	}
	out := new(ServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.SecondaryCidrBlocks != nil {
		in, out := &in.SecondaryCidrBlocks, &out.SecondaryCidrBlocks
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PrivateGoogleAccess != nil {
		in, out := &in.PrivateGoogleAccess, &out.PrivateGoogleAccess
		*out = new(bool)
		**out = **in
	}
	if in.FlowLogs != nil {
		in, out := &in.FlowLogs, &out.FlowLogs
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Subnets) DeepCopyInto(out *Subnets) {
	{
		in := &in
		*out = make(Subnets, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SubnetSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subnets.
func (in Subnets) DeepCopy() Subnets {
	if in == nil {
		return nil
	}
	out := new(Subnets)
	in.DeepCopyInto(out)
	return *out
}