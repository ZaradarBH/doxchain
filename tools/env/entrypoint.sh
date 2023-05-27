#!/bin/bash

rm -rf mytestnet
pkill terrad

# enable api and swagger
export DOXCHAIND_API_ENABLE=true
export DOXCHAIND_API_SWAGGER=true

BINARY=${BINARY:-doxchaind}
# check DENOM is set. If not, set to uluna
DENOM=${DENOM:-udox}

CHAIN_ID=${CHAIN_ID:-.test}
HOME_DIR=${HOME_DIR:-.doxchaind}
KEYRING="test"
KEY="test"
KEY1="test1"
KEY2="test2"

# Function updates the config based on a jq argument as a string
update_test_genesis () {
    # update_test_genesis '.consensus_params["block"]["max_gas"]="100000000"'
    cat $HOME_DIR/config/genesis.json | jq "$1" > $HOME_DIR/config/tmp_genesis.json && mv $HOME_DIR/config/tmp_genesis.json $HOME_DIR/config/genesis.json
}

$BINARY init --chain-id $CHAIN_ID moniker --home $HOME_DIR

$BINARY keys add $KEY --keyring-backend $KEYRING --home $HOME_DIR
$BINARY keys add $KEY1 --keyring-backend $KEYRING --home $HOME_DIR
$BINARY keys add $KEY2 --keyring-backend $KEYRING --home $HOME_DIR

# Allocate genesis accounts (cosmos formatted addresses)
$BINARY add-genesis-account $KEY "1000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME_DIR
$BINARY add-genesis-account $KEY1 "1000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME_DIR
$BINARY add-genesis-account $KEY2 "1000000000000${DENOM}" --keyring-backend $KEYRING --home $HOME_DIR

update_test_genesis '.app_state["gov"]["voting_params"]["voting_period"]="50s"'
update_test_genesis '.app_state["mint"]["params"]["mint_denom"]="'$DENOM'"'
update_test_genesis '.app_state["gov"]["deposit_params"]["min_deposit"]=[{"denom":"'$DENOM'","amount": "1000000"}]'
update_test_genesis '.app_state["crisis"]["constant_fee"]={"denom":"'$DENOM'","amount":"1000"}'
update_test_genesis '.app_state["staking"]["params"]["bond_denom"]="'$DENOM'"'

# Sign genesis transaction
$BINARY gentx $KEY "1000000${DENOM}" --keyring-backend $KEYRING --chain-id $CHAIN_ID --home $HOME_DIR

# Collect genesis tx
$BINARY collect-gentxs --home $HOME_DIR

# Run this to ensure everything worked and that the genesis file is setup correctly
$BINARY validate-genesis --home $HOME_DIR

$BINARY start --home $HOME_DIR