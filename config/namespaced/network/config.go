package network

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_network", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "Network"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/avarei/provider-vra/v2/apis/cluster/project/v1alpha1.Project",
		}
	})
	p.AddResourceConfigurator("vra_network_ip_range", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "NetworkIPRange"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/avarei/provider-vra/v2/apis/cluster/project/v1alpha1.Project",
		}
	})
	p.AddResourceConfigurator("vra_network_profile", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.Kind = "NetworkProfile"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/avarei/provider-vra/v2/apis/cluster/project/v1alpha1.Project",
		}
	})
}
