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

creator=$($BINARY keys show $KEY --keyring-backend $KEYRING --home $HOME_DIR -a)
w3c_identifer=""
method_name="example1"

# create did document
create_did_document () {
    MOCK_DID_DOCUMENT=$(jq -n '
    {
        "id" : {
            "creator":"'$creator'",
            "url":"did:example:123457/path",
            "methodName":"'$method_name'"
        }
    }')

    echo $MOCK_DID_DOCUMENT > scripts/chain-interaction/did-document.json

    echo "CREATING DID DOCUMENT: $MOCK_DID_DOCUMENT"

    res=$($BINARY tx did create-did-document scripts/chain-interaction/did-document.json --from $KEY --keyring-backend $KEYRING --home $HOME_DIR --chain-id $CHAIN_ID --gas auto --yes -o json)
    txhash=$(echo $res | jq -r '.txhash')
    echo $res

    sleep $SLEEP

    w3c_identifer=$($BINARY query tx $txhash -o json | jq -r ".raw_log" | jq -r ".[0].events[0].attributes[0].value")
    echo "result: $w3c_identifer"
}

delete_did_document () {
    res=$($BINARY tx did delete-did-document $w3c_identifer --from $KEY --keyring-backend $KEYRING --home $HOME_DIR --chain-id $CHAIN_ID --gas auto --yes -o json)
    txhash=$(echo $res | jq -r '.txhash')
    echo $res

    sleep $SLEEP

    query=$($BINARY query tx $txhash -o json | jq ".raw_log")
    echo "result: $query"
}

create_did_document
delete_did_document