import * as registrygeoreplication from "@pulumi/registrygeoreplication";
import * as resources from "@pulumi/azure-native/resources";

const resourceGroup = new resources.ResourceGroup("resourceGroup");

const registry = new registrygeoreplication.RegistryGeoReplication("registry", {
    resourceGroup: resourceGroup,
});

export const login = registry.acrLoginServerOut;
