// Copyright © 2019 The Things Industries B.V.

package tenant

import (
	"context"

	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/ttipb"
)

var errNoTenantID = errors.DefineNotFound("no_tenant_id", "no tenant ID")
var errNoFetcher = errors.DefineNotFound("no_fetcher", "no fetcher")
var errNoConfig = errors.DefineNotFound("no_config", "no configuration")

func ConfigFromContext(ctx context.Context) (*ttipb.Configuration, error) {
	id := FromContext(ctx)
	if id.TenantID == "" {
		return nil, errNoTenantID
	}
	fetcher, ok := FetcherFromContext(ctx)
	if !ok {
		return nil, errNoFetcher
	}
	tenant, err := fetcher.FetchTenant(ctx, &id, "configuration")
	if err != nil {
		return nil, err
	}
	conf := tenant.GetConfiguration()
	if conf == nil {
		return nil, errNoConfig
	}
	return conf, nil
}
