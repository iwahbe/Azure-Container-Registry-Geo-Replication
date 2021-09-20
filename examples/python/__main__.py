"""An Azure RM Python Pulumi program"""

import pulumi
from pulumi_azure_native import resources
from pulumi_registrygeoreplication import RegistryGeoReplication

# Create an Azure Resource Group
resource_group = resources.ResourceGroup('resource_group')

# Create an Azure resource (Storage Account)
registry = RegistryGeoReplication('registry',
    name="registry",
    resource_group_name=resource_group.name,
    replication_location="westus")

pulumi.export("login", registry.login_server_out)
pulumi.export("underlying_registry_id", registry.registry.Id)
