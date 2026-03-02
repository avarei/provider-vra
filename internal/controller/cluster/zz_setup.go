/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	blockdevice "github.com/avarei/provider-vra/v2/internal/controller/cluster/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/avarei/provider-vra/v2/internal/controller/cluster/blockdevice/blockdevicesnapshot"
	blueprint "github.com/avarei/provider-vra/v2/internal/controller/cluster/blueprint/blueprint"
	version "github.com/avarei/provider-vra/v2/internal/controller/cluster/blueprint/version"
	catalogitementitlement "github.com/avarei/provider-vra/v2/internal/controller/cluster/catalogitementitlement/catalogitementitlement"
	catalogsourceblueprint "github.com/avarei/provider-vra/v2/internal/controller/cluster/catalogsourceblueprint/catalogsourceblueprint"
	catalogsourceentitlement "github.com/avarei/provider-vra/v2/internal/controller/cluster/catalogsourceentitlement/catalogsourceentitlement"
	accountaws "github.com/avarei/provider-vra/v2/internal/controller/cluster/cloudaccount/accountaws"
	accountazure "github.com/avarei/provider-vra/v2/internal/controller/cluster/cloudaccount/accountazure"
	accountgcp "github.com/avarei/provider-vra/v2/internal/controller/cluster/cloudaccount/accountgcp"
	accountnsxt "github.com/avarei/provider-vra/v2/internal/controller/cluster/cloudaccount/accountnsxt"
	accountvmc "github.com/avarei/provider-vra/v2/internal/controller/cluster/cloudaccount/accountvmc"
	accountvsphere "github.com/avarei/provider-vra/v2/internal/controller/cluster/cloudaccount/accountvsphere"
	contentsharingpolicy "github.com/avarei/provider-vra/v2/internal/controller/cluster/contentsharing/contentsharingpolicy"
	contentsource "github.com/avarei/provider-vra/v2/internal/controller/cluster/contentsource/contentsource"
	deployment "github.com/avarei/provider-vra/v2/internal/controller/cluster/deployment/deployment"
	compute "github.com/avarei/provider-vra/v2/internal/controller/cluster/fabric/compute"
	datastorevsphere "github.com/avarei/provider-vra/v2/internal/controller/cluster/fabric/datastorevsphere"
	networkvsphere "github.com/avarei/provider-vra/v2/internal/controller/cluster/fabric/networkvsphere"
	profile "github.com/avarei/provider-vra/v2/internal/controller/cluster/flavorprofile/profile"
	profileimageprofile "github.com/avarei/provider-vra/v2/internal/controller/cluster/imageprofile/profile"
	integration "github.com/avarei/provider-vra/v2/internal/controller/cluster/integration/integration"
	loadbalancer "github.com/avarei/provider-vra/v2/internal/controller/cluster/loadbalancer/loadbalancer"
	machine "github.com/avarei/provider-vra/v2/internal/controller/cluster/machine/machine"
	network "github.com/avarei/provider-vra/v2/internal/controller/cluster/network/network"
	networkiprange "github.com/avarei/provider-vra/v2/internal/controller/cluster/network/networkiprange"
	networkprofile "github.com/avarei/provider-vra/v2/internal/controller/cluster/network/networkprofile"
	project "github.com/avarei/provider-vra/v2/internal/controller/cluster/project/project"
	providerconfig "github.com/avarei/provider-vra/v2/internal/controller/cluster/providerconfig"
	profilestorage "github.com/avarei/provider-vra/v2/internal/controller/cluster/storage/profile"
	profileaws "github.com/avarei/provider-vra/v2/internal/controller/cluster/storage/profileaws"
	profileazure "github.com/avarei/provider-vra/v2/internal/controller/cluster/storage/profileazure"
	profilevsphere "github.com/avarei/provider-vra/v2/internal/controller/cluster/storage/profilevsphere"
	zone "github.com/avarei/provider-vra/v2/internal/controller/cluster/zone/zone"
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
