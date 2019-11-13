// Copyright © 2019 The Things Industries B.V.

package tenantbillingserver

import (
	"context"

	"github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/pkg/log"
	"go.thethings.network/lorawan-stack/pkg/tenant"
	"go.thethings.network/lorawan-stack/pkg/ttipb"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

// Report implements ttipb.TbsServer.
func (tbs *TenantBillingServer) Report(ctx context.Context, data *ttipb.MeteringData) (*types.Empty, error) {
	tenantFetcher, ok := tenant.FetcherFromContext(ctx)
	if !ok {
		panic("tenant fetcher not available")
	}
	logger := log.FromContext(ctx)
outer:
	for _, tenantData := range data.Tenants {
		tenant, err := tenantFetcher.FetchTenant(ctx, &tenantData.TenantIdentifiers, "attributes", "state")
		if err != nil {
			logger.WithError(err).Error("Failed to retrieve tenant")
			break outer
		}
		for _, backend := range tbs.backends {
			err := backend.Report(ctx, tenant, tenantData.Totals)
			if err != nil {
				logger.WithError(err).Error("Failed to report metrics to backend")
				break outer
			}
		}
	}
	return ttnpb.Empty, nil
}
