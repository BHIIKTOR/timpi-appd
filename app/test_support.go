package app

import (
	capabilitykeeper "github.com/cosmos/ibc-go/modules/capability/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"

	"github.com/cosmos/cosmos-sdk/baseapp"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

func (app *TimpiApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

func (app *TimpiApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

func (app *TimpiApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

func (app *TimpiApp) GetBankKeeper() bankkeeper.Keeper {
	return app.BankKeeper
}

func (app *TimpiApp) GetStakingKeeper() *stakingkeeper.Keeper {
	return app.StakingKeeper
}

func (app *TimpiApp) GetAccountKeeper() authkeeper.AccountKeeper {
	return app.AccountKeeper
}

func (app *TimpiApp) GetWasmKeeper() wasmkeeper.Keeper {
	return app.WasmKeeper
}
