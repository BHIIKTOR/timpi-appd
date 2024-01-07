package v020

import (
	"context"
	"fmt"

	// "context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	// capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"cosmossdk.io/x/nft"
	"github.com/cosmos/cosmos-sdk/x/group"

	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"

	ibccapabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	v6 "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/migrations/v6"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	ibctmmigrations "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint/migrations"

	storetypes "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/cosmos/cosmos-sdk/types/module"

	"timpid/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name
const UpgradeName = "v0.2.0"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added: []string{
			authtypes.ModuleName,
			genutiltypes.ModuleName,
			banktypes.ModuleName,
			stakingtypes.ModuleName,
			distrtypes.ModuleName,
			govtypes.ModuleName,
			paramstypes.ModuleName,

			ibccapabilitytypes.ModuleName,

			// SDK 46
			group.ModuleName,
			nft.ModuleName,

			ibcfeetypes.ModuleName,

			// SDK 47
			crisistypes.ModuleName,
			consensusparamtypes.ModuleName,

			// SDK 50
			circuittypes.ModuleName,
		},
		Deleted: []string{},
	},
}

func CreateUpgradeHandler(
	mm upgrades.ModuleManager,
	configurator module.Configurator,
	ak *upgrades.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		sdkCtx := sdk.UnwrapSDKContext(ctx)

		logger := sdkCtx.Logger().With("upgrade", UpgradeName)

		logger.Info(` _    _ _____   _____ _____            _____  ______ `)
		logger.Info(`| |  | |  __ \ / ____|  __ \     /\   |  __ \|  ____|`)
		logger.Info(`| |  | | |__) | |  __| |__) |   /  \  | |  | | |__   `)
		logger.Info(`| |  | |  ___/| | |_ |  _  /   / /\ \ | |  | |  __|  `)
		logger.Info(`| |__| | |    | |__| | | \ \  / ____ \| |__| | |____ `)
		logger.Info(` \____/|_|     \_____|_|  \_\/_/    \_\_____/|______|`)

		// fromVM[ibcfeetypes.ModuleName] = mm.Modules[ibcfeetypes.ModuleName].ConsensusVersion()
		// sdkCtx.Logger().Info(fmt.Sprintf("ibcfee module version %s set", fmt.Sprint(fromVM[ibcfeetypes.ModuleName])))

		// ibc v6
		// NOTE: The moduleName arg of v6.CreateUpgradeHandler refers to the auth module ScopedKeeper name to which the channel capability should be migrated from.
		// This should be the same string value provided upon instantiation of the ScopedKeeper with app.CapabilityKeeper.ScopeToModule()
		const moduleName = icacontrollertypes.SubModuleName
		if err := v6.MigrateICS27ChannelCapability(sdkCtx, ak.Codec, ak.GetStoreKey(ibccapabilitytypes.ModuleName),
			ak.CapabilityKeeper, moduleName); err != nil {
			return nil, err
		}

		// ibc v7
		if _, err := ibctmmigrations.PruneExpiredConsensusStates(sdkCtx, ak.Codec, ak.IBCKeeper.ClientKeeper); err != nil {
			return nil, err
		}

		// run migrations
		versionMap, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}
		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))

		return versionMap, err
	}
}
