#!/bin/sh
#set -o errexit -o nounset -o pipefail

PASSWORD=${PASSWORD:-1234567890}
STAKE=${STAKE_TOKEN:-ustake}
FEE=${FEE_TOKEN:-ucosm}
CHAIN_ID=${CHAIN_ID:-testing}
MONIKER=${MONIKER:-node001}

timpid init --chain-id "$CHAIN_ID" "$MONIKER"
# staking/governance token is hardcoded in config, change this
sed -i "s/\"stake\"/\"$STAKE\"/" "$HOME"/.timpid/config/genesis.json
# this is essential for sub-1s block times (or header times go crazy)
sed -i 's/"time_iota_ms": "1000"/"time_iota_ms": "10"/' "$HOME"/.timpid/config/genesis.json

if ! timpid keys show validator; then
  (echo "$PASSWORD"; echo "$PASSWORD") | timpid keys add validator
fi
# hardcode the validator account for this instance
echo "$PASSWORD" | timpid genesis add-genesis-account validator "1000000000$STAKE,1000000000$FEE"

# (optionally) add a few more genesis accounts
for addr in "$@"; do
  echo $addr
  timpid genesis add-genesis-account "$addr" "1000000000$STAKE,1000000000$FEE"
done

# submit a genesis validator tx
## Workraround for https://github.com/cosmos/cosmos-sdk/issues/8251
(echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | timpid genesis gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID" --amount="250000000$STAKE"
## should be:
# (echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | timpid gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID"
timpid genesis collect-gentxs
