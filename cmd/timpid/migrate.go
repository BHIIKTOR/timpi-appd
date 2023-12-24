package main

import (
	"fmt"
	"path/filepath"
	"timpid/app"

	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cobra"
)

func MigrateCmd(appCreator servertypes.AppCreator) *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate [version]",
		Short: "migrates version",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := server.GetServerContextFromCmd(cmd)

			home := ctx.Config.RootDir
			db, err := openDB(home, server.GetAppDBBackend(ctx.Viper))
			if err != nil {
				return err
			}
			logger := log.NewLogger(cmd.OutOrStdout())
			timpiApp := appCreator(logger, db, nil, ctx.Viper)
			timpiApp.(*app.TimpiApp).Migrate(0)
			fmt.Println("200")
			return nil
		},
	}

	return migrateCmd
}

func openDB(rootDir string, backendType dbm.BackendType) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	return dbm.NewDB("application", backendType, dataDir)
}
