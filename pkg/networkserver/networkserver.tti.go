// Copyright © 2020 The Things Industries B.V.

package networkserver

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/pkg/tenant"
	"go.thethings.network/lorawan-stack/pkg/ttipb"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
)

func init() {
	DefaultOptions = append(DefaultOptions,
		WithTenantOverrides(),
	)
}

func tenantConfigFromContext(ctx context.Context) (*ttipb.Configuration_Cluster_NetworkServer, bool) {
	conf, err := tenant.ConfigFromContext(ctx)
	if err != nil {
		return nil, false
	}
	nsConf := conf.GetDefaultCluster().GetNS()
	return nsConf, nsConf != nil
}

func WithTenantOverrides() Option {
	return func(ns *NetworkServer) {
		origNewDevAddr := ns.newDevAddr
		ns.newDevAddr = func(ctx context.Context, dev *ttnpb.EndDevice) types.DevAddr {
			conf, ok := tenantConfigFromContext(ctx)
			if !ok || len(conf.DevAddrPrefixes) == 0 {
				return origNewDevAddr(ctx, dev)
			}
			return makeNewDevAddrFunc(conf.DevAddrPrefixes...)(ctx, dev)
		}

		origDeduplicationDone := ns.deduplicationDone
		ns.deduplicationDone = func(ctx context.Context, up *ttnpb.UplinkMessage) <-chan time.Time {
			conf, ok := tenantConfigFromContext(ctx)
			if !ok || conf.DeduplicationWindow == nil {
				return origDeduplicationDone(ctx, up)
			}
			return makeWindowEndAfterFunc(*conf.DeduplicationWindow)(ctx, up)
		}

		origCollectionDone := ns.collectionDone
		ns.collectionDone = func(ctx context.Context, up *ttnpb.UplinkMessage) <-chan time.Time {
			conf, ok := tenantConfigFromContext(ctx)
			if !ok || conf.DeduplicationWindow == nil || conf.CooldownWindow == nil {
				return origCollectionDone(ctx, up)
			}
			return makeWindowEndAfterFunc(*conf.DeduplicationWindow+*conf.CooldownWindow)(ctx, up)
		}
	}
}
