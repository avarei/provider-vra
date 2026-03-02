/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	blockdevice "github.com/avarei/provider-vra/v2/internal/controller/namespaced/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/avarei/provider-vra/v2/internal/controller/namespaced/blockdevice/blockdevicesnapshot"
	blueprint "github.com/avarei/provider-vra/v2/internal/controller/namespaced/blueprint/blueprint"
	version "github.com/avarei/provider-vra/v2/internal/controller/namespaced/blueprint/version"
	catalogitementitlement "github.com/avarei/provider-vra/v2/internal/controller/namespaced/catalogitementitlement/catalogitementitlement"
	catalogsourceblueprint "github.com/avarei/provider-vra/v2/internal/controller/namespaced/catalogsourceblueprint/catalogsourceblueprint"
	catalogsourceentitlement "github.com/avarei/provider-vra/v2/internal/controller/namespaced/catalogsourceentitlement/catalogsourceentitlement"
	accountaws "github.com/avarei/provider-vra/v2/internal/controller/namespaced/cloudaccount/accountaws"
	accountazure "github.com/avarei/provider-vra/v2/internal/controller/namespaced/cloudaccount/accountazure"
	accountgcp "github.com/avarei/provider-vra/v2/internal/controller/namespaced/cloudaccount/accountgcp"
	accountnsxt "github.com/avarei/provider-vra/v2/internal/controller/namespaced/cloudaccount/accountnsxt"
	accountvmc "github.com/avarei/provider-vra/v2/internal/controller/namespaced/cloudaccount/accountvmc"
	accountvsphere "github.com/avarei/provider-vra/v2/internal/controller/namespaced/cloudaccount/accountvsphere"
	contentsharingpolicy "github.com/avarei/provider-vra/v2/internal/controller/namespaced/contentsharing/contentsharingpolicy"
	contentsource "github.com/avarei/provider-vra/v2/internal/controller/namespaced/contentsource/contentsource"
	deployment "github.com/avarei/provider-vra/v2/internal/controller/namespaced/deployment/deployment"
	compute "github.com/avarei/provider-vra/v2/internal/controller/namespaced/fabric/compute"
	datastorevsphere "github.com/avarei/provider-vra/v2/internal/controller/namespaced/fabric/datastorevsphere"
	networkvsphere "github.com/avarei/provider-vra/v2/internal/controller/namespaced/fabric/networkvsphere"
	profile "github.com/avarei/provider-vra/v2/internal/controller/namespaced/flavorprofile/profile"
	profileimageprofile "github.com/avarei/provider-vra/v2/internal/controller/namespaced/imageprofile/profile"
	integration "github.com/avarei/provider-vra/v2/internal/controller/namespaced/integration/integration"
	loadbalancer "github.com/avarei/provider-vra/v2/internal/controller/namespaced/loadbalancer/loadbalancer"
	machine "github.com/avarei/provider-vra/v2/internal/controller/namespaced/machine/machine"
	network "github.com/avarei/provider-vra/v2/internal/controller/namespaced/network/network"
	networkiprange "github.com/avarei/provider-vra/v2/internal/controller/namespaced/network/networkiprange"
	networkprofile "github.com/avarei/provider-vra/v2/internal/controller/namespaced/network/networkprofile"
	project "github.com/avarei/provider-vra/v2/internal/controller/namespaced/project/project"
	providerconfig "github.com/avarei/provider-vra/v2/internal/controller/namespaced/providerconfig"
	profilestorage "github.com/avarei/provider-vra/v2/internal/controller/namespaced/storage/profile"
	profileaws "github.com/avarei/provider-vra/v2/internal/controller/namespaced/storage/profileaws"
	profileazure "github.com/avarei/provider-vra/v2/internal/controller/namespaced/storage/profileazure"
	profilevsphere "github.com/avarei/provider-vra/v2/internal/controller/namespaced/storage/profilevsphere"
	zone "github.com/avarei/provider-vra/v2/internal/controller/namespaced/zone/zone"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockdevice.Setup,
		blockdevicesnapshot.Setup,
		blueprint.Setup,
		version.Setup,
		catalogitementitlement.Setup,
		catalogsourceblueprint.Setup,
		catalogsourceentitlement.Setup,
		accountaws.Setup,
		accountazure.Setup,
		accountgcp.Setup,
		accountnsxt.Setup,
		accountvmc.Setup,
		accountvsphere.Setup,
		contentsharingpolicy.Setup,
		contentsource.Setup,
		deployment.Setup,
		compute.Setup,
		datastorevsphere.Setup,
		networkvsphere.Setup,
		profile.Setup,
		profileimageprofile.Setup,
		integration.Setup,
		loadbalancer.Setup,
		machine.Setup,
		network.Setup,
		networkiprange.Setup,
		networkprofile.Setup,
		project.Setup,
		providerconfig.Setup,
		profilestorage.Setup,
		profileaws.Setup,
		profileazure.Setup,
		profilevsphere.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockdevice.SetupGated,
		blockdevicesnapshot.SetupGated,
		blueprint.SetupGated,
		version.SetupGated,
		catalogitementitlement.SetupGated,
		catalogsourceblueprint.SetupGated,
		catalogsourceentitlement.SetupGated,
		accountaws.SetupGated,
		accountazure.SetupGated,
		accountgcp.SetupGated,
		accountnsxt.SetupGated,
		accountvmc.SetupGated,
		accountvsphere.SetupGated,
		contentsharingpolicy.SetupGated,
		contentsource.SetupGated,
		deployment.SetupGated,
		compute.SetupGated,
		datastorevsphere.SetupGated,
		networkvsphere.SetupGated,
		profile.SetupGated,
		profileimageprofile.SetupGated,
		integration.SetupGated,
		loadbalancer.SetupGated,
		machine.SetupGated,
		network.SetupGated,
		networkiprange.SetupGated,
		networkprofile.SetupGated,
		project.SetupGated,
		providerconfig.SetupGated,
		profilestorage.SetupGated,
		profileaws.SetupGated,
		profileazure.SetupGated,
		profilevsphere.SetupGated,
		zone.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
