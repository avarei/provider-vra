/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document

	"context"
	_ "embed"

	blockdeviceCluster "github.com/avarei/provider-vra/v2/config/cluster/block_device"
	blueprintCluster "github.com/avarei/provider-vra/v2/config/cluster/blueprint"
	catalogitementitlementCluster "github.com/avarei/provider-vra/v2/config/cluster/catalog_item"
	catalogsourceCluster "github.com/avarei/provider-vra/v2/config/cluster/catalog_source"
	cloudaccountCluster "github.com/avarei/provider-vra/v2/config/cluster/cloud_account"
	contentsharingCluster "github.com/avarei/provider-vra/v2/config/cluster/content_sharing"
	contentsourceCluster "github.com/avarei/provider-vra/v2/config/cluster/content_source"
	deploymentCluster "github.com/avarei/provider-vra/v2/config/cluster/deployment"
	fabricCluster "github.com/avarei/provider-vra/v2/config/cluster/fabric"
	flavorprofileCluster "github.com/avarei/provider-vra/v2/config/cluster/flavor_profile"
	imageprofileCluster "github.com/avarei/provider-vra/v2/config/cluster/image_profile"
	integrationCluster "github.com/avarei/provider-vra/v2/config/cluster/integration"
	loadbalancerCluster "github.com/avarei/provider-vra/v2/config/cluster/load_balancer"
	machineCluster "github.com/avarei/provider-vra/v2/config/cluster/machine"
	networkCluster "github.com/avarei/provider-vra/v2/config/cluster/network"
	projectCluster "github.com/avarei/provider-vra/v2/config/cluster/project"
	storageCluster "github.com/avarei/provider-vra/v2/config/cluster/storage"
	zoneCluster "github.com/avarei/provider-vra/v2/config/cluster/zone"
	blockdeviceNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/block_device"
	blueprintNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/blueprint"
	catalogitementitlementNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/catalog_item"
	catalogsourceNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/catalog_source"
	cloudaccountNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/cloud_account"
	contentsharingNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/content_sharing"
	contentsourceNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/content_source"
	deploymentNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/deployment"
	fabricNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/fabric"
	flavorprofileNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/flavor_profile"
	imageprofileNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/image_profile"
	integrationNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/integration"
	loadbalancerNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/load_balancer"
	machineNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/machine"
	networkNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/network"
	projectNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/project"
	storageNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/storage"
	zoneNamespaced "github.com/avarei/provider-vra/v2/config/namespaced/zone"

	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/v2/pkg/types/conversion/tfjson"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"

	tfjson "github.com/hashicorp/terraform-json"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfvra "github.com/vmware/terraform-provider-vra/vra"
)

const (
	resourcePrefix = "vra"
	modulePath     = "github.com/avarei/provider-vra/v2"
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
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		projectCluster.Configure,
		blueprintCluster.Configure,
		deploymentCluster.Configure,
		fabricCluster.Configure,
		blockdeviceCluster.Configure,
		flavorprofileCluster.Configure,
		imageprofileCluster.Configure,
		storageCluster.Configure,
		catalogsourceCluster.Configure,
		catalogitementitlementCluster.Configure,
		cloudaccountCluster.Configure,
		contentsharingCluster.Configure,
		contentsourceCluster.Configure,
		integrationCluster.Configure,
		loadbalancerCluster.Configure,
		machineCluster.Configure,
		networkCluster.Configure,
		zoneCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// GetProvider returns provider configuration
func GetProviderNamespaced(_ context.Context, generationProvider bool) (*ujconfig.Provider, error) {
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
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithRootGroup("vra.m.crossplane.io"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		projectNamespaced.Configure,
		blueprintNamespaced.Configure,
		deploymentNamespaced.Configure,
		fabricNamespaced.Configure,
		blockdeviceNamespaced.Configure,
		flavorprofileNamespaced.Configure,
		imageprofileNamespaced.Configure,
		storageNamespaced.Configure,
		catalogsourceNamespaced.Configure,
		catalogitementitlementNamespaced.Configure,
		cloudaccountNamespaced.Configure,
		contentsharingNamespaced.Configure,
		contentsourceNamespaced.Configure,
		integrationNamespaced.Configure,
		loadbalancerNamespaced.Configure,
		machineNamespaced.Configure,
		networkNamespaced.Configure,
		zoneNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// ResourcesWithExternalNameConfig returns the list of resources that have external
// name configured in ExternalNameConfigs table.
func ResourcesWithExternalNameConfig() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}
