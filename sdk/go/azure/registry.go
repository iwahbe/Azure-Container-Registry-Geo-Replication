// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	containerregistry "github.com/pulumi/pulumi-azure-native/sdk/go/azure/containerregistry"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Registry struct {
	pulumi.ResourceState

	// The login server url
	LoginServer pulumi.StringOutput `pulumi:"loginServer"`
	// The Registry
	Registry containerregistry.RegistryOutput `pulumi:"registry"`
	// The replication policy
	Replication containerregistry.ReplicationOutput `pulumi:"replication"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ReplicationLocation == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource Registry
	err := ctx.RegisterRemoteComponentResource("azure-quickstart-acr-geo-replication:index:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type registryArgs struct {
	// Enable admin user that has push / pull permissions to the registry
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// Globally unique name of your azure container registry
	Name string `pulumi:"name"`
	// The location of the registry replica location
	ReplicationLocation string `pulumi:"replicationLocation"`
	// The name of the enclosing resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tier of your Azure Container Registry. Geo-replication requires the Premium SKU
	Sku *string `pulumi:"sku"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// Enable admin user that has push / pull permissions to the registry
	AdminUserEnabled pulumi.BoolPtrInput
	// Globally unique name of your azure container registry
	Name pulumi.StringInput
	// The location of the registry replica location
	ReplicationLocation pulumi.StringInput
	// The name of the enclosing resource group
	ResourceGroupName pulumi.StringInput
	// Tier of your Azure Container Registry. Geo-replication requires the Premium SKU
	Sku pulumi.StringPtrInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil))
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

func (i *Registry) ToRegistryPtrOutput() RegistryPtrOutput {
	return i.ToRegistryPtrOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPtrOutput)
}

type RegistryPtrInput interface {
	pulumi.Input

	ToRegistryPtrOutput() RegistryPtrOutput
	ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput
}

type registryPtrType RegistryArgs

func (*registryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil))
}

func (i *registryPtrType) ToRegistryPtrOutput() RegistryPtrOutput {
	return i.ToRegistryPtrOutputWithContext(context.Background())
}

func (i *registryPtrType) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPtrOutput)
}

// RegistryArrayInput is an input type that accepts RegistryArray and RegistryArrayOutput values.
// You can construct a concrete instance of `RegistryArrayInput` via:
//
//          RegistryArray{ RegistryArgs{...} }
type RegistryArrayInput interface {
	pulumi.Input

	ToRegistryArrayOutput() RegistryArrayOutput
	ToRegistryArrayOutputWithContext(context.Context) RegistryArrayOutput
}

type RegistryArray []RegistryInput

func (RegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Registry)(nil)).Elem()
}

func (i RegistryArray) ToRegistryArrayOutput() RegistryArrayOutput {
	return i.ToRegistryArrayOutputWithContext(context.Background())
}

func (i RegistryArray) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryArrayOutput)
}

// RegistryMapInput is an input type that accepts RegistryMap and RegistryMapOutput values.
// You can construct a concrete instance of `RegistryMapInput` via:
//
//          RegistryMap{ "key": RegistryArgs{...} }
type RegistryMapInput interface {
	pulumi.Input

	ToRegistryMapOutput() RegistryMapOutput
	ToRegistryMapOutputWithContext(context.Context) RegistryMapOutput
}

type RegistryMap map[string]RegistryInput

func (RegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Registry)(nil)).Elem()
}

func (i RegistryMap) ToRegistryMapOutput() RegistryMapOutput {
	return i.ToRegistryMapOutputWithContext(context.Background())
}

func (i RegistryMap) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryMapOutput)
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Registry)(nil))
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryPtrOutput() RegistryPtrOutput {
	return o.ToRegistryPtrOutputWithContext(context.Background())
}

func (o RegistryOutput) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Registry) *Registry {
		return &v
	}).(RegistryPtrOutput)
}

type RegistryPtrOutput struct{ *pulumi.OutputState }

func (RegistryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil))
}

func (o RegistryPtrOutput) ToRegistryPtrOutput() RegistryPtrOutput {
	return o
}

func (o RegistryPtrOutput) ToRegistryPtrOutputWithContext(ctx context.Context) RegistryPtrOutput {
	return o
}

func (o RegistryPtrOutput) Elem() RegistryOutput {
	return o.ApplyT(func(v *Registry) Registry {
		if v != nil {
			return *v
		}
		var ret Registry
		return ret
	}).(RegistryOutput)
}

type RegistryArrayOutput struct{ *pulumi.OutputState }

func (RegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Registry)(nil))
}

func (o RegistryArrayOutput) ToRegistryArrayOutput() RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) Index(i pulumi.IntInput) RegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Registry {
		return vs[0].([]Registry)[vs[1].(int)]
	}).(RegistryOutput)
}

type RegistryMapOutput struct{ *pulumi.OutputState }

func (RegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Registry)(nil))
}

func (o RegistryMapOutput) ToRegistryMapOutput() RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) MapIndex(k pulumi.StringInput) RegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Registry {
		return vs[0].(map[string]Registry)[vs[1].(string)]
	}).(RegistryOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryOutput{})
	pulumi.RegisterOutputType(RegistryPtrOutput{})
	pulumi.RegisterOutputType(RegistryArrayOutput{})
	pulumi.RegisterOutputType(RegistryMapOutput{})
}
