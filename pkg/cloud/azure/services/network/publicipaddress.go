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

package network

import (
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-10-01/network"
)

func (s *Service) GetPublicIpAddress(resourceGroup string, IPName string) (network.PublicIPAddress, error) {
	return s.scope.AzureClients.PublicIPAddresses.Get(s.scope.Context, resourceGroup, IPName, "")
}

func (s *Service) DeletePublicIpAddress(resourceGroup string, IPName string) (network.PublicIPAddressesDeleteFuture, error) {
	return s.scope.AzureClients.PublicIPAddresses.Delete(s.scope.Context, resourceGroup, IPName)
}

// TODO: Dead code
/*
func (s *Service) WaitForPublicIpAddressDeleteFuture(future network.PublicIPAddressesDeleteFuture) error {
	return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AzureClients.PublicIPAddresses.Client)
}
*/
