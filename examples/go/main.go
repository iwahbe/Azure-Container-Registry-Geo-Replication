package main

import (
	"github.com/pulumi/pulumi-azure-quickstart-acr-geo-replication/sdk/go/azure"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an Azure Resource Group
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}
		azure-quickstart-acr-geo-replication
		registry, err := acr.NewRegistry(ctx, "registry", &azure.RegistryGeoReplicationArgs{
			Name:                "registry",
			ReplicationLocation: "westus",
			ResourceGroupName:   resourceGroup.Name,
		})

		if err != nil {
			return err
		}

		// Export the primary key of the Storage Account
		ctx.Export("login", registry.LoginServer)

		return nil
	})
}
