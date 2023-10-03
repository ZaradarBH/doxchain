#!/bin/bash

BINARY=${1:-build/doxchaind}
DENOM=${2:-udox}
CHAIN_ID="localdoxchain"
KEYRING="test"
KEY="test0"
KEY1="test1"
KEY2="test2"
HOME_DIR=mytestnet
SLEEP=5

list_did () {
    echo "QUERYING FOR DID"
    $BINARY q did list-did
}

creator=$($BINARY keys show $KEY --keyring-backend $KEYRING --home $HOME_DIR -a)
w3c_identifer=""
method_name="example1"
# create did
create_did () {
    MOCK_DID=$(jq -n '
    {
        "creator":"'$creator'",
        "url":"did:example:123457/path",
        "methodName":"'$method_name'"
    }')

    echo "CREATING DID: $MOCK_DID"

    res=$($BINARY tx did create-did "$MOCK_DID" --from $KEY --keyring-backend $KEYRING --home $HOME_DIR --chain-id $CHAIN_ID --gas auto --yes -o json)
    txhash=$(echo $res | jq -r '.txhash')
    echo $res

    sleep $SLEEP

    w3c_identifer=$($BINARY query tx $txhash -o json | jq -r ".raw_log" | jq -r ".[0].events[0].attributes[0].value")
    echo "result: $w3c_identifer"

    list_did
}

# update did
update_did () {
    UPDATE_DID=$(jq -n '
    {
        "creator":"'$creator'",
        "url":"did:example:123457/path2",
        "methodName":"'$method_name'"
    }')

    echo "UPDATING DID TO: $UPDATE_DID"

    res=$($BINARY tx did update-did "$UPDATE_DID" --from $KEY --keyring-backend $KEYRING --home $HOME_DIR --chain-id $CHAIN_ID --gas auto --yes -o json)
    txhash=$(echo $res | jq -r '.txhash')
    echo $res

    sleep $SLEEP

    query=$($BINARY query tx $txhash -o json | jq ".raw_log")
    echo "result: $query"

    list_did
}

# delete did
delete_did () {
    echo "DELETING DID: $w3c_identifer"

    res=$($BINARY tx did delete-did $w3c_identifer --from $KEY --keyring-backend $KEYRING --home $HOME_DIR --chain-id $CHAIN_ID --gas auto --yes -o json)
    txhash=$(echo $res | jq -r '.txhash')
    echo $res

    sleep $SLEEP

    query=$($BINARY query tx $txhash -o json | jq ".raw_log")
    echo "result: $query"

    list_did
}

create_did
update_did
delete_did