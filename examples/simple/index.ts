import * as registrygeoreplication from "@pulumi/registrygeoreplication";
import * as resources from "@pulumi/azure-native";

const group = new resources.ResourceGroup("group")

const registry = new registrygeoreplication.RegistryGeoReplication("registry", {
    resourceGroup: group,
});

export const login = registry.acrLoginServerOut;
