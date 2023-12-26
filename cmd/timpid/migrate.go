package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"timpid/app"

	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cobra"
)

const (
	flagVersion = "version"
)

func MigrateCmd(appCreator servertypes.AppCreator) *cobra.Command {
	cmd := &cobra.Command{
		Use:                "migrate [version]",
		Aliases:            []string{"m"},
		Short:              "Migrates version",
		DisableFlagParsing: false,
		Args:               cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			ctx := server.GetServerContextFromCmd(cmd)

			home := ctx.Config.RootDir
			db, err := openDB(home, server.GetAppDBBackend(ctx.Viper))
			if err != nil {
				return err
			}

			logger := log.NewLogger(cmd.OutOrStdout())

			timpiApp := appCreator(logger, db, nil, ctx.Viper)

			timpiApp.(*app.TimpiApp).Migrate(version)

			fmt.Println("MigrateCmd realized")

			return nil
		},
	}

	return cmd
}

func openDB(rootDir string, backendType dbm.BackendType) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	return dbm.NewDB("application", backendType, dataDir)
}
