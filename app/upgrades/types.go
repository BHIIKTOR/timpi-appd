package upgrades

import (
	"context"

	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	storetypes "cosmossdk.io/store/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	// ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/keeper"
	icahostkeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/keeper"
	ibcfeekeeper "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/keeper"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	// ibchookskeeper "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/keeper"
)

type AppKeepers struct {
	// keys    map[string]*storetypes.KVStoreKey
	// tkeys   map[string]*storetypes.TransientStoreKey
	// memKeys map[string]*storetypes.MemoryStoreKey

	GetStoreKey func(storeKey string) *storetypes.KVStoreKey

	AccountKeeper         *authkeeper.AccountKeeper
	ParamsKeeper          *paramskeeper.Keeper
	ConsensusParamsKeeper *consensusparamkeeper.Keeper
	Codec                 codec.Codec
	CapabilityKeeper      *capabilitykeeper.Keeper

	IBCKeeper           *ibckeeper.Keeper
	IBCFeeKeeper        ibcfeekeeper.Keeper
	ICAControllerKeeper icacontrollerkeeper.Keeper
	ICAHostKeeper       icahostkeeper.Keeper
	TransferKeeper      ibctransferkeeper.Keeper
	// IBCHooksKeeper      *ibchookskeeper.Keeper
}

type ModuleManager interface {
	RunMigrations(ctx context.Context, cfg module.Configurator, fromVM module.VersionMap) (module.VersionMap, error)
	GetVersionMap() module.VersionMap
}

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler
	CreateUpgradeHandler func(ModuleManager, module.Configurator, *AppKeepers) upgradetypes.UpgradeHandler
	StoreUpgrades        storetypes.StoreUpgrades
}

// type Upgrade struct {
// 	// Upgrade version name, for the upgrade handler, e.g. `v7`
// 	UpgradeName string

// 	// CreateUpgradeHandler defines the function that creates an upgrade handler
// 	CreateUpgradeHandler func(
// 		*module.Manager,
// 		module.Configurator,
// 		*keepers.AppKeepers,
// 	) upgradetypes.UpgradeHandler

// 	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
// 	StoreUpgrades store.StoreUpgrades
// }
