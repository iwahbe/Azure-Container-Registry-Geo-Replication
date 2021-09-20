import * as registrygeoreplication from "@pulumi/registrygeoreplication";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("resourceGroup");

const registry = new registrygeoreplication.RegistryGeoReplication("registry", {
    name: "registry",
    location: "northcentralus",
    replicationLocation: "westus",
    resourceGroupName: resourceGroup.name,
});

export const login = registry.loginServerOut;
export const underlying_registry_id = registry.registry.id;
