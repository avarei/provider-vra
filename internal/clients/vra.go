/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/avarei/provider-vra/apis/v1beta1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal vRA credentials as JSON"
)

const (
	keyURL                = "url"
	keyAccessToken        = "access_token"
	keyRefreshToken       = "refresh_token"
	keyInsecure           = "insecure"
	keyReauthorizeTimeout = "reauthorize_timeout"
	keyApiTimeout         = "api_timeout"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
// nolint:gocyclo
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		// Configuration is a map of provider configuration values.
		vraCreds := map[string]string{}
		if err := json.Unmarshal(data, &vraCreds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration = map[string]interface{}{}
		if v, ok := vraCreds[keyURL]; ok {
			ps.Configuration[keyURL] = v
		}
		if v, ok := vraCreds[keyAccessToken]; ok {
			ps.Configuration[keyAccessToken] = v
		}
		if v, ok := vraCreds[keyRefreshToken]; ok {
			ps.Configuration[keyRefreshToken] = v
		}
		if v, ok := vraCreds[keyInsecure]; ok {
			ps.Configuration[keyInsecure] = v
		}
		if v, ok := vraCreds[keyReauthorizeTimeout]; ok {
			ps.Configuration[keyReauthorizeTimeout] = v
		}
		if v, ok := vraCreds[keyApiTimeout]; ok {
			ps.Configuration[keyApiTimeout] = v
		}

		// Set credentials in Terraform provider environment.
		/*ps.Env = []string{
			fmt.Sprintf("%s=%s", envURL, vraCreds[keyURL]),
			fmt.Sprintf("%s=%s", envRefreshToken, vraCreds[keyRefreshToken]),
		}*/

		// Set credentials in Terraform provider configuration.
		/*ps.Configuration = map[string]any{
			"username": creds["username"],
			"password": creds["password"],
		}*/
		return ps, errors.Wrap(
			configureNoForkVaultClient(ctx, &ps, *tfProvider),
			"failed to configure the no-fork GCP client",
		)
	}
}
func configureNoForkVaultClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}
