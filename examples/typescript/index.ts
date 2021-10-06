import * as acr from "@pulumi/azure-quickstart-acr-geo-replication";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("resourceGroup");

const registry = new acr.ReplicatedRegistry("registry", {
    name: "registry",
    replicationLocation: "westus",
    resourceGroupName: resourceGroup.name,
});

export const login = registry.loginServer;
export const underlying_registry_id = registry.registry.id;
