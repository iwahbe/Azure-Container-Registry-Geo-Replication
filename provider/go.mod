module github.com/pulumi/pulumi-registrygeoreplication

go 1.15

// TODO: This can go away when we cut v3.13.j
replace github.com/pulumi/pulumi/pkg/v3 => github.com/pulumi/pulumi/pkg/v3 v3.12.1-0.20210920171144-c338876b9f5b

replace github.com/pulumi/pulumi/sdk/v3 => github.com/pulumi/pulumi/sdk/v3 v3.12.1-0.20210920171144-c338876b9f5b

require (
	github.com/pkg/errors v0.9.1
	github.com/pulumi/pulumi-azure-native/sdk v1.28.0
	github.com/pulumi/pulumi/pkg/v3 v3.12.0
	github.com/pulumi/pulumi/sdk/v3 v3.12.0
)
