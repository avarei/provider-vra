package cloudaccount

import "github.com/crossplane/upjet/v2/pkg/config"

const version string = "v1alpha1"
const shortGroup string = "cloudAccount"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_cloud_account_aws", func(r *config.Resource) {
		r.Version = version
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_azure", func(r *config.Resource) {
		r.Version = version
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_gcp", func(r *config.Resource) {
		r.Version = version
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_nsxt", func(r *config.Resource) {
		r.Version = version
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_nsxv", func(r *config.Resource) {
		r.Version = version
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_vmc", func(r *config.Resource) {
		r.Version = version
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("vra_cloud_account_vsphere", func(r *config.Resource) {
		r.Version = version
		r.ShortGroup = shortGroup
	})
}
