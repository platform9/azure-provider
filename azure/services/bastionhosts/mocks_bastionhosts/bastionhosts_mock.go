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

// Code generated by MockGen. DO NOT EDIT.
// Source: ../bastionhosts.go

// Package mock_bastionhosts is a generated GoMock package.
package mock_bastionhosts

import (
	reflect "reflect"

	autorest "github.com/Azure/go-autorest/autorest"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	v1alpha4 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha4"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
)

// MockBastionScope is a mock of BastionScope interface.
type MockBastionScope struct {
	ctrl     *gomock.Controller
	recorder *MockBastionScopeMockRecorder
}

// MockBastionScopeMockRecorder is the mock recorder for MockBastionScope.
type MockBastionScopeMockRecorder struct {
	mock *MockBastionScope
}

// NewMockBastionScope creates a new mock instance.
func NewMockBastionScope(ctrl *gomock.Controller) *MockBastionScope {
	mock := &MockBastionScope{ctrl: ctrl}
	mock.recorder = &MockBastionScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBastionScope) EXPECT() *MockBastionScopeMockRecorder {
	return m.recorder
}

// APIServerLBName mocks base method.
func (m *MockBastionScope) APIServerLBName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServerLBName")
	ret0, _ := ret[0].(string)
	return ret0
}

// APIServerLBName indicates an expected call of APIServerLBName.
func (mr *MockBastionScopeMockRecorder) APIServerLBName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServerLBName", reflect.TypeOf((*MockBastionScope)(nil).APIServerLBName))
}

// APIServerLBPoolName mocks base method.
func (m *MockBastionScope) APIServerLBPoolName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIServerLBPoolName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// APIServerLBPoolName indicates an expected call of APIServerLBPoolName.
func (mr *MockBastionScopeMockRecorder) APIServerLBPoolName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIServerLBPoolName", reflect.TypeOf((*MockBastionScope)(nil).APIServerLBPoolName), arg0)
}

// AdditionalTags mocks base method.
func (m *MockBastionScope) AdditionalTags() v1alpha4.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha4.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockBastionScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockBastionScope)(nil).AdditionalTags))
}

// Authorizer mocks base method.
func (m *MockBastionScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockBastionScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockBastionScope)(nil).Authorizer))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockBastionScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockBastionScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockBastionScope)(nil).AvailabilitySetEnabled))
}

// BaseURI mocks base method.
func (m *MockBastionScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockBastionScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockBastionScope)(nil).BaseURI))
}

// BastionSpec mocks base method.
func (m *MockBastionScope) BastionSpec() azure.BastionSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BastionSpec")
	ret0, _ := ret[0].(azure.BastionSpec)
	return ret0
}

// BastionSpec indicates an expected call of BastionSpec.
func (mr *MockBastionScopeMockRecorder) BastionSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BastionSpec", reflect.TypeOf((*MockBastionScope)(nil).BastionSpec))
}

// ClientID mocks base method.
func (m *MockBastionScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockBastionScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockBastionScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockBastionScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockBastionScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockBastionScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockBastionScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockBastionScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockBastionScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockBastionScope) CloudProviderConfigOverrides() *v1alpha4.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1alpha4.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockBastionScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockBastionScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockBastionScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockBastionScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockBastionScope)(nil).ClusterName))
}

// ControlPlaneRouteTable mocks base method.
func (m *MockBastionScope) ControlPlaneRouteTable() v1alpha4.RouteTable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneRouteTable")
	ret0, _ := ret[0].(v1alpha4.RouteTable)
	return ret0
}

// ControlPlaneRouteTable indicates an expected call of ControlPlaneRouteTable.
func (mr *MockBastionScopeMockRecorder) ControlPlaneRouteTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneRouteTable", reflect.TypeOf((*MockBastionScope)(nil).ControlPlaneRouteTable))
}

// ControlPlaneSubnet mocks base method.
func (m *MockBastionScope) ControlPlaneSubnet() v1alpha4.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ControlPlaneSubnet")
	ret0, _ := ret[0].(v1alpha4.SubnetSpec)
	return ret0
}

// ControlPlaneSubnet indicates an expected call of ControlPlaneSubnet.
func (mr *MockBastionScopeMockRecorder) ControlPlaneSubnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ControlPlaneSubnet", reflect.TypeOf((*MockBastionScope)(nil).ControlPlaneSubnet))
}

// Enabled mocks base method.
func (m *MockBastionScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockBastionScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockBastionScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockBastionScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockBastionScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockBastionScope)(nil).Error), varargs...)
}

// GetPrivateDNSZoneName mocks base method.
func (m *MockBastionScope) GetPrivateDNSZoneName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDNSZoneName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPrivateDNSZoneName indicates an expected call of GetPrivateDNSZoneName.
func (mr *MockBastionScopeMockRecorder) GetPrivateDNSZoneName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDNSZoneName", reflect.TypeOf((*MockBastionScope)(nil).GetPrivateDNSZoneName))
}

// HashKey mocks base method.
func (m *MockBastionScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockBastionScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockBastionScope)(nil).HashKey))
}

// Info mocks base method.
func (m *MockBastionScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockBastionScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBastionScope)(nil).Info), varargs...)
}

// IsAPIServerPrivate mocks base method.
func (m *MockBastionScope) IsAPIServerPrivate() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAPIServerPrivate")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAPIServerPrivate indicates an expected call of IsAPIServerPrivate.
func (mr *MockBastionScopeMockRecorder) IsAPIServerPrivate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAPIServerPrivate", reflect.TypeOf((*MockBastionScope)(nil).IsAPIServerPrivate))
}

// IsIPv6Enabled mocks base method.
func (m *MockBastionScope) IsIPv6Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsIPv6Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsIPv6Enabled indicates an expected call of IsIPv6Enabled.
func (mr *MockBastionScopeMockRecorder) IsIPv6Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsIPv6Enabled", reflect.TypeOf((*MockBastionScope)(nil).IsIPv6Enabled))
}

// IsVnetManaged mocks base method.
func (m *MockBastionScope) IsVnetManaged() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsVnetManaged")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsVnetManaged indicates an expected call of IsVnetManaged.
func (mr *MockBastionScopeMockRecorder) IsVnetManaged() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsVnetManaged", reflect.TypeOf((*MockBastionScope)(nil).IsVnetManaged))
}

// Location mocks base method.
func (m *MockBastionScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockBastionScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockBastionScope)(nil).Location))
}

// NodeSubnets mocks base method.
func (m *MockBastionScope) NodeSubnets() []v1alpha4.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeSubnets")
	ret0, _ := ret[0].([]v1alpha4.SubnetSpec)
	return ret0
}

// NodeSubnets indicates an expected call of NodeSubnets.
func (mr *MockBastionScopeMockRecorder) NodeSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeSubnets", reflect.TypeOf((*MockBastionScope)(nil).NodeSubnets))
}

// OutboundLBName mocks base method.
func (m *MockBastionScope) OutboundLBName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutboundLBName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// OutboundLBName indicates an expected call of OutboundLBName.
func (mr *MockBastionScopeMockRecorder) OutboundLBName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundLBName", reflect.TypeOf((*MockBastionScope)(nil).OutboundLBName), arg0)
}

// OutboundPoolName mocks base method.
func (m *MockBastionScope) OutboundPoolName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OutboundPoolName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// OutboundPoolName indicates an expected call of OutboundPoolName.
func (mr *MockBastionScopeMockRecorder) OutboundPoolName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutboundPoolName", reflect.TypeOf((*MockBastionScope)(nil).OutboundPoolName), arg0)
}

// ResourceGroup mocks base method.
func (m *MockBastionScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockBastionScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockBastionScope)(nil).ResourceGroup))
}

// SetSubnet mocks base method.
func (m *MockBastionScope) SetSubnet(arg0 v1alpha4.SubnetSpec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubnet", arg0)
}

// SetSubnet indicates an expected call of SetSubnet.
func (mr *MockBastionScopeMockRecorder) SetSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnet", reflect.TypeOf((*MockBastionScope)(nil).SetSubnet), arg0)
}

// Subnet mocks base method.
func (m *MockBastionScope) Subnet(arg0 string) v1alpha4.SubnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnet", arg0)
	ret0, _ := ret[0].(v1alpha4.SubnetSpec)
	return ret0
}

// Subnet indicates an expected call of Subnet.
func (mr *MockBastionScopeMockRecorder) Subnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnet", reflect.TypeOf((*MockBastionScope)(nil).Subnet), arg0)
}

// Subnets mocks base method.
func (m *MockBastionScope) Subnets() v1alpha4.Subnets {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].(v1alpha4.Subnets)
	return ret0
}

// Subnets indicates an expected call of Subnets.
func (mr *MockBastionScopeMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockBastionScope)(nil).Subnets))
}

// SubscriptionID mocks base method.
func (m *MockBastionScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockBastionScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockBastionScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockBastionScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockBastionScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockBastionScope)(nil).TenantID))
}

// V mocks base method.
func (m *MockBastionScope) V(level int) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockBastionScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockBastionScope)(nil).V), level)
}

// Vnet mocks base method.
func (m *MockBastionScope) Vnet() *v1alpha4.VnetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vnet")
	ret0, _ := ret[0].(*v1alpha4.VnetSpec)
	return ret0
}

// Vnet indicates an expected call of Vnet.
func (mr *MockBastionScopeMockRecorder) Vnet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vnet", reflect.TypeOf((*MockBastionScope)(nil).Vnet))
}

// WithName mocks base method.
func (m *MockBastionScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockBastionScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockBastionScope)(nil).WithName), name)
}

// WithValues mocks base method.
func (m *MockBastionScope) WithValues(keysAndValues ...interface{}) logr.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithValues", varargs...)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithValues indicates an expected call of WithValues.
func (mr *MockBastionScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockBastionScope)(nil).WithValues), keysAndValues...)
}
