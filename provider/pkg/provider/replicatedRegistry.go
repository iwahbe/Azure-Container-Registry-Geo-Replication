// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"

	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/containerregistry"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a ReplicatedRegistryArgs component resource
type ReplicatedRegistryArgs struct {
	// Globally unique name of your azure container registry
	Name string `pulumi:"name"`
	// Enable admin user that has push / pull permissions to the registry
	AdminUserEnabled bool `pulumi:"adminUserEnabled"`
	// Tier of your Azure Container Registry. Geo-replication requires the Premium SKU
	Sku string `pulumi:"sku"`
	// The location of the registry replica location
	ReplicationLocation string `pulumi:"replicationLocation"`
	// The name of the enclosing resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The ReplicatedRegistry component resource.
type ReplicatedRegistry struct {
	pulumi.ResourceState

	Registry    *containerregistry.Registry    `pulumi:"registry"`
	Replication *containerregistry.Replication `pulumi:"replication"`
	LoginServer pulumi.StringOutput            `pulumi:"loginServer"`
}

// NewReplicatedRegistry creates a new ACR Replicated Registry component resource.
func NewReplicatedRegistry(ctx *pulumi.Context,
	name string, args *ReplicatedRegistryArgs, opts ...pulumi.ResourceOption) (*ReplicatedRegistry, error) {
	if args == nil {
		args = &ReplicatedRegistryArgs{}
	}

	component := &ReplicatedRegistry{}

	err := ctx.RegisterComponentResource("azure-quickstart-acr-geo-replication:index:ReplicatedRegistry", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Required parameters
	acrAdminUserEnabledParam := args.AdminUserEnabled
	acrNameParam := args.Name
	acrReplicaLocationParam := args.ReplicationLocation
	resourceGroupNameParam := args.ResourceGroupName

	// Optional parameters
	acrSkuParam := args.Sku
	if acrSkuParam == "" {
		acrSkuParam = string(containerregistry.SkuNamePremium)
	}

	resourceGroupVar, err := resources.LookupResourceGroup(ctx, &resources.LookupResourceGroupArgs{
		ResourceGroupName: resourceGroupNameParam,
	}, nil)
	if err != nil {
		return nil, err
	}

	registryResource, err := containerregistry.NewRegistry(ctx, "registryResource", &containerregistry.RegistryArgs{
		AdminUserEnabled:  pulumi.Bool(acrAdminUserEnabledParam),
		Location:          pulumi.String(resourceGroupVar.Location),
		RegistryName:      pulumi.String(acrNameParam),
		ResourceGroupName: pulumi.String(resourceGroupNameParam),
		Sku: &containerregistry.SkuArgs{
			Name: pulumi.String(acrSkuParam),
		},
		Tags: pulumi.StringMap{
			"container.registry": pulumi.String(acrNameParam),
			"displayName":        pulumi.String("Container Registry"),
		},
	})
	if err != nil {
		return nil, err
	}

	replication, err := containerregistry.NewReplication(ctx, "replicationResource", &containerregistry.ReplicationArgs{
		Location:          pulumi.String(acrReplicaLocationParam),
		RegistryName:      registryResource.Name,
		ReplicationName:   pulumi.String(fmt.Sprintf("%v/%v", acrNameParam, acrReplicaLocationParam)),
		ResourceGroupName: pulumi.String(resourceGroupNameParam),
	})
	if err != nil {
		return nil, err
	}

	component.Replication = replication
	component.Registry = registryResource
	component.LoginServer = registryResource.LoginServer

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"registry":    registryResource,
		"loginServer": registryResource.LoginServer,
		"replication": replication,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
