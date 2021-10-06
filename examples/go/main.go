package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	acr "github.com/pulumi/pulumi-azure-quickstart-acr-geo-replication/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		resourceGroup, err := resources.NewResourceGroup(ctx, "resourceGroup", nil)
		if err != nil {
			return err
		}

		registry, err := acr.NewReplicatedRegistry(ctx, "registry", &acr.ReplicatedRegistryArgs{
			Name:                "registry",
			ReplicationLocation: "westus",
			ResourceGroupName:   resourceGroup.Name,
		})

		if err != nil {
			return err
		}

		ctx.Export("login", registry.LoginServer)

		return nil
	})
}
