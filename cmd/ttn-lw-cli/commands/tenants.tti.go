// Copyright © 2019 The Things Industries B.V.

package commands

import (
	"os"

	"github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/io"
	"go.thethings.network/lorawan-stack/cmd/ttn-lw-cli/internal/util"
	"go.thethings.network/lorawan-stack/pkg/errors"
	"go.thethings.network/lorawan-stack/pkg/ttipb"
	"google.golang.org/grpc"
)

var (
	selectTenantFlags = util.FieldMaskFlags(&ttipb.Tenant{})
	setTenantFlags    = util.FieldFlags(&ttipb.Tenant{})
)

func tenantIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("tenant-id", "", "")
	return flagSet
}

var errNoTenantID = errors.DefineInvalidArgument("no_tenant_id", "no tenant ID set")

func getTenantID(flagSet *pflag.FlagSet, args []string) *ttipb.TenantIdentifiers {
	var tenantID string
	if len(args) > 0 {
		if len(args) > 1 {
			logger.Warn("multiple IDs found in arguments, considering only the first")
		}
		tenantID = args[0]
	} else {
		tenantID, _ = flagSet.GetString("tenant-id")
	}
	if tenantID == "" {
		return nil
	}
	return &ttipb.TenantIdentifiers{TenantID: tenantID}
}

func getTenantAdminCreds(cmd *cobra.Command) grpc.CallOption {
	tenantAdminKey, _ := cmd.Flags().GetString("tenant-admin-key")
	return api.GetCredentials("TenantAdminKey", tenantAdminKey)
}

var (
	tenantsCommand = &cobra.Command{
		Use:               "tenants",
		Hidden:            true,
		Aliases:           []string{"tenant", "tnt", "t"},
		Short:             "Tenant commands",
		PersistentPreRunE: preRun(),
	}
	tenantsListCommand = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List tenants",
		RunE: func(cmd *cobra.Command, args []string) error {
			paths := util.SelectFieldMask(cmd.Flags(), selectTenantFlags)

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			limit, page, opt, getTotal := withPagination(cmd.Flags())
			res, err := ttipb.NewTenantRegistryClient(is).List(ctx, &ttipb.ListTenantsRequest{
				FieldMask: types.FieldMask{Paths: paths},
				Limit:     limit,
				Page:      page,
			}, opt, getTenantAdminCreds(cmd))
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res.Tenants)
		},
	}
	tenantsGetCommand = &cobra.Command{
		Use:     "get [tenant-id]",
		Aliases: []string{"info"},
		Short:   "Get a tenant",
		RunE: func(cmd *cobra.Command, args []string) error {
			refreshToken() // NOTE: ignore errors.
			optionalAuth()

			cliID := getTenantID(cmd.Flags(), args)
			if cliID == nil {
				return errNoTenantID
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectTenantFlags)

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttipb.NewTenantRegistryClient(is).Get(ctx, &ttipb.GetTenantRequest{
				TenantIdentifiers: *cliID,
				FieldMask:         types.FieldMask{Paths: paths},
			}, getTenantAdminCreds(cmd))
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	tenantsCreateCommand = &cobra.Command{
		Use:     "create [tenant-id]",
		Aliases: []string{"add", "register"},
		Short:   "Create a tenant",
		RunE: asBulk(func(cmd *cobra.Command, args []string) (err error) {
			cliID := getTenantID(cmd.Flags(), args)
			var tenant ttipb.Tenant
			if inputDecoder != nil {
				_, err := inputDecoder.Decode(&tenant)
				if err != nil {
					return err
				}
			}
			if err := util.SetFields(&tenant, setTenantFlags); err != nil {
				return err
			}
			tenant.Attributes = mergeAttributes(tenant.Attributes, cmd.Flags())
			if cliID != nil && cliID.TenantID != "" {
				tenant.TenantID = cliID.TenantID
			}
			if tenant.TenantID == "" {
				return errNoTenantID
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttipb.NewTenantRegistryClient(is).Create(ctx, &ttipb.CreateTenantRequest{
				Tenant: tenant,
			}, getTenantAdminCreds(cmd))
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		}),
	}
	tenantsUpdateCommand = &cobra.Command{
		Use:     "update [tenant-id]",
		Aliases: []string{"set"},
		Short:   "Update a tenant",
		RunE: func(cmd *cobra.Command, args []string) error {
			refreshToken() // NOTE: ignore errors.
			optionalAuth()

			cliID := getTenantID(cmd.Flags(), args)
			if cliID == nil {
				return errNoTenantID
			}
			paths := util.UpdateFieldMask(cmd.Flags(), setTenantFlags, attributesFlags())
			if len(paths) == 0 {
				logger.Warn("No fields selected, won't update anything")
				return nil
			}
			var tenant ttipb.Tenant
			if err := util.SetFields(&tenant, setTenantFlags); err != nil {
				return err
			}
			tenant.Attributes = mergeAttributes(tenant.Attributes, cmd.Flags())
			tenant.TenantIdentifiers = *cliID

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttipb.NewTenantRegistryClient(is).Update(ctx, &ttipb.UpdateTenantRequest{
				Tenant:    tenant,
				FieldMask: types.FieldMask{Paths: paths},
			}, getTenantAdminCreds(cmd))
			if err != nil {
				return err
			}

			res.SetFields(&tenant, "ids")
			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	tenantsDeleteCommand = &cobra.Command{
		Use:   "delete [tenant-id]",
		Short: "Delete a tenant",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliID := getTenantID(cmd.Flags(), args)
			if cliID == nil {
				return errNoTenantID
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttipb.NewTenantRegistryClient(is).Delete(ctx, cliID, getTenantAdminCreds(cmd))
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	tenantsListCommand.Flags().AddFlagSet(selectTenantFlags)
	tenantsListCommand.Flags().AddFlagSet(paginationFlags())
	tenantsCommand.AddCommand(tenantsListCommand)
	tenantsGetCommand.Flags().AddFlagSet(tenantIDFlags())
	tenantsGetCommand.Flags().AddFlagSet(selectTenantFlags)
	tenantsCommand.AddCommand(tenantsGetCommand)
	tenantsCreateCommand.Flags().AddFlagSet(tenantIDFlags())
	tenantsCreateCommand.Flags().AddFlagSet(setTenantFlags)
	tenantsCreateCommand.Flags().AddFlagSet(attributesFlags())
	tenantsCommand.AddCommand(tenantsCreateCommand)
	tenantsUpdateCommand.Flags().AddFlagSet(tenantIDFlags())
	tenantsUpdateCommand.Flags().AddFlagSet(setTenantFlags)
	tenantsUpdateCommand.Flags().AddFlagSet(attributesFlags())
	tenantsCommand.AddCommand(tenantsUpdateCommand)
	tenantsDeleteCommand.Flags().AddFlagSet(tenantIDFlags())
	tenantsCommand.AddCommand(tenantsDeleteCommand)
	tenantsCommand.PersistentFlags().String("tenant-admin-key", "", "Tenant Admin Key")
	Root.AddCommand(tenantsCommand)
}
