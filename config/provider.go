/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document

	"context"
	_ "embed"

	blockdevice "github.com/avarei/provider-vra/config/block_device"
	blueprint "github.com/avarei/provider-vra/config/blueprint"
	catalogitementitlement "github.com/avarei/provider-vra/config/catalog_item"
	catalogsource "github.com/avarei/provider-vra/config/catalog_source"
	cloudaccount "github.com/avarei/provider-vra/config/cloud_account"
	contentsharing "github.com/avarei/provider-vra/config/content_sharing"
	contentsource "github.com/avarei/provider-vra/config/content_source"
	deployment "github.com/avarei/provider-vra/config/deployment"
	fabric "github.com/avarei/provider-vra/config/fabric"
	flavorprofile "github.com/avarei/provider-vra/config/flavor_profile"
	imageprofile "github.com/avarei/provider-vra/config/image_profile"
	integration "github.com/avarei/provider-vra/config/integration"
	loadbalancer "github.com/avarei/provider-vra/config/load_balancer"
	machine "github.com/avarei/provider-vra/config/machine"
	network "github.com/avarei/provider-vra/config/network"
	project "github.com/avarei/provider-vra/config/project"
	storage "github.com/avarei/provider-vra/config/storage"
	zone "github.com/avarei/provider-vra/config/zone"

	"github.com/pkg/errors"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"

	tfjson "github.com/hashicorp/terraform-json"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfvra "github.com/vmware/terraform-provider-vra/vra"
)

const (
	resourcePrefix = "vra"
	modulePath     = "github.com/avarei/provider-vra"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// workaround for the no-fork release: We would like to
// keep the types in the generated CRDs intact
// (prevent number->int type replacements).
func getProviderSchema(s string) (*tfschema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &tfschema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*ujconfig.Provider, error) {
	sdkProvider := tfvra.Provider()

	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		blueprint.Configure,
		deployment.Configure,
		fabric.Configure,
		blockdevice.Configure,
		flavorprofile.Configure,
		imageprofile.Configure,
		storage.Configure,
		catalogsource.Configure,
		catalogitementitlement.Configure,
		cloudaccount.Configure,
		contentsharing.Configure,
		contentsource.Configure,
		integration.Configure,
		loadbalancer.Configure,
		machine.Configure,
		network.Configure,
		zone.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}
