// Copyright © 2019 The Things Industries B.V.

package identityserver

import (
	"fmt"
	"testing"

	pbtypes "github.com/gogo/protobuf/types"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/rpcmetadata"
	"go.thethings.network/lorawan-stack/pkg/tenant"
	"go.thethings.network/lorawan-stack/pkg/ttipb"
	"go.thethings.network/lorawan-stack/pkg/types"
	"go.thethings.network/lorawan-stack/pkg/util/test"
	"google.golang.org/grpc"
)

func TestTenantsUnauthenticated(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttipb.NewTenantRegistryClient(cc)

		eui1 := types.EUI64{1, 2, 3, 4, 5, 6, 7, 8}
		eui2 := types.EUI64{8, 7, 6, 5, 4, 3, 2, 1}

		_, err := reg.Create(ctx, &ttipb.CreateTenantRequest{
			Tenant: ttipb.Tenant{
				TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: "foo-tnt"},
			},
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.Get(ctx, &ttipb.GetTenantRequest{
			TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: "foo-tnt"},
			FieldMask:         pbtypes.FieldMask{Paths: []string{"name"}},
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.GetIdentifiersForGatewayEUI(ctx, &ttipb.GetTenantIdentifiersForGatewayEUIRequest{
			EUI: eui1,
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.GetIdentifiersForEndDeviceEUIs(ctx, &ttipb.GetTenantIdentifiersForEndDeviceEUIsRequest{
			JoinEUI: eui1,
			DevEUI:  eui2,
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.List(ctx, &ttipb.ListTenantsRequest{
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.Update(ctx, &ttipb.UpdateTenantRequest{
			Tenant: ttipb.Tenant{
				TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: "foo-tnt"},
				Name:              "Updated Name",
			},
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
		})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}

		_, err = reg.Delete(ctx, &ttipb.TenantIdentifiers{TenantID: "foo-tnt"})

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsUnauthenticated(err), should.BeTrue)
		}
	})
}

func TestTenantsPermissionDenied(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	creds := userCreds(defaultUserIdx, "key without rights")

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttipb.NewTenantRegistryClient(cc)

		_, err := reg.Create(ctx, &ttipb.CreateTenantRequest{
			Tenant: ttipb.Tenant{
				TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: "foo-tnt"},
			},
		}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Get(ctx, &ttipb.GetTenantRequest{
			TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: "foo-tnt"},
			FieldMask:         pbtypes.FieldMask{Paths: []string{"name"}},
		}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.List(ctx, &ttipb.ListTenantsRequest{
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
		}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Update(ctx, &ttipb.UpdateTenantRequest{
			Tenant: ttipb.Tenant{
				TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: "foo-tnt"},
				Name:              "Updated Name",
			},
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
		}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Delete(ctx, &ttipb.TenantIdentifiers{TenantID: "foo-tnt"}, creds)

		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}
	})
}

func TestTenantsCRUD(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		is.config.Tenancy.AdminKeys = []string{"BEEFBEEFBEEFBEEFBEEFBEEFBEEFBEEF"}
		a.So(is.config.Tenancy.decodeAdminKeys(ctx), should.BeNil)

		reg := ttipb.NewTenantRegistryClient(cc)

		creds := grpc.PerRPCCredentials(rpcmetadata.MD{
			AuthType:  TenantAdminAuthType,
			AuthValue: is.config.Tenancy.AdminKeys[0],
		})

		clusterCreds := is.WithClusterAuth()

		created, err := reg.Create(ctx, &ttipb.CreateTenantRequest{
			Tenant: ttipb.Tenant{
				TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: "foo"},
				Name:              "Foo Tenant",
			},
		}, creds)

		a.So(err, should.BeNil)
		a.So(created.Name, should.Equal, "Foo Tenant")

		got, err := reg.Get(ctx, &ttipb.GetTenantRequest{
			TenantIdentifiers: created.TenantIdentifiers,
			FieldMask:         pbtypes.FieldMask{Paths: []string{"name"}},
		}, creds)

		a.So(err, should.BeNil)
		a.So(got.Name, should.Equal, created.Name)

		got, err = reg.Get(tenant.NewContext(ctx, created.TenantIdentifiers), &ttipb.GetTenantRequest{
			TenantIdentifiers: created.TenantIdentifiers,
			FieldMask:         pbtypes.FieldMask{Paths: []string{"ids"}},
		}, clusterCreds)

		a.So(err, should.BeNil)

		updated, err := reg.Update(ctx, &ttipb.UpdateTenantRequest{
			Tenant: ttipb.Tenant{
				TenantIdentifiers: created.TenantIdentifiers,
				Name:              "Updated Name",
			},
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
		}, creds)

		a.So(err, should.BeNil)
		a.So(updated.Name, should.Equal, "Updated Name")

		list, err := reg.List(ctx, &ttipb.ListTenantsRequest{
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
		}, creds)
		a.So(err, should.BeNil)
		if a.So(list.Tenants, should.NotBeEmpty) {
			var found bool
			for _, item := range list.Tenants {
				if item.TenantIdentifiers == created.TenantIdentifiers {
					found = true
					a.So(item.Name, should.Equal, updated.Name)
				}
			}
			a.So(found, should.BeTrue)
		}

		_, err = reg.Delete(ctx, &created.TenantIdentifiers, creds)
		a.So(err, should.BeNil)
	})
}

func TestTenantsPagination(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		is.config.Tenancy.AdminKeys = []string{"BEEFBEEFBEEFBEEFBEEFBEEFBEEFBEEF"}
		a.So(is.config.Tenancy.decodeAdminKeys(ctx), should.BeNil)

		reg := ttipb.NewTenantRegistryClient(cc)

		creds := grpc.PerRPCCredentials(rpcmetadata.MD{
			AuthType:  TenantAdminAuthType,
			AuthValue: is.config.Tenancy.AdminKeys[0],
		})

		for i := 0; i < 3; i++ {
			reg.Create(ctx, &ttipb.CreateTenantRequest{
				Tenant: ttipb.Tenant{
					TenantIdentifiers: ttipb.TenantIdentifiers{TenantID: fmt.Sprintf("foo-%d", i)},
					Name:              fmt.Sprintf("Foo Tenant %d", i),
				},
			}, creds)
		}

		list, err := reg.List(test.Context(), &ttipb.ListTenantsRequest{
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
			Limit:     2,
			Page:      1,
		}, creds)

		a.So(list, should.NotBeNil)
		a.So(list.Tenants, should.HaveLength, 2)
		a.So(err, should.BeNil)

		list, err = reg.List(test.Context(), &ttipb.ListTenantsRequest{
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
			Limit:     2,
			Page:      2,
		}, creds)

		a.So(list, should.NotBeNil)
		a.So(list.Tenants, should.HaveLength, 1)
		a.So(err, should.BeNil)

		list, err = reg.List(test.Context(), &ttipb.ListTenantsRequest{
			FieldMask: pbtypes.FieldMask{Paths: []string{"name"}},
			Limit:     2,
			Page:      3,
		}, creds)

		a.So(list, should.NotBeNil)
		a.So(list.Tenants, should.BeEmpty)
		a.So(err, should.BeNil)
	})
}
