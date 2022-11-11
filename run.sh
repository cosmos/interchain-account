#!/bin/bash

# Store the following account addresses within the current shell env
export WALLET_1=$(icad keys show wallet1 -a --keyring-backend test --home ./data/test-1) && echo $WALLET_1
export WALLET_2=$(icad keys show wallet2 -a --keyring-backend test --home ./data/test-1) && echo $WALLET_2
export WALLET_3=$(icad keys show wallet3 -a --keyring-backend test --home ./data/test-2) && echo $WALLET_3
export WALLET_4=$(icad keys show wallet4 -a --keyring-backend test --home ./data/test-2) && echo $WALLET_4

# Register an interchain account on behalf of WALLET_1 where chain test-2 is the interchain accounts host
printf "\n== Registering interchain account.. ==\n\n"
icad tx intertx register --from $WALLET_1 --connection-id connection-0 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y 1>/dev/null 2>&1

#adjust sleep time to wait for relayer as needed
sleep 10 
printf "\n== Interchain account registered ==\n\n"

# Query the address of the interchain account
icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/test-1 --node tcp://localhost:16657
sleep 3
# Store the interchain account address by parsing the query result: cosmos1hd0f4u7zgptymmrn55h3hy20jv2u0ctdpq23cpe8m9pas8kzd87smtf8al
export ICA_ADDR=$(icad query intertx interchainaccounts connection-0 $WALLET_1 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') && echo $ICA_ADDR

# Query the interchain account balance on the host chain. It should be empty.
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657

# Send funds to the interchain account.
icad tx bank send $WALLET_3 $ICA_ADDR 10000stake --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 --keyring-backend test -y 1>/dev/null 2>&1
sleep 3
printf "\n== Interchain account funded ==\n\n"

# Query the balance once again and observe the changes
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657

# Submit a bank send tx using the interchain account via ibc

printf "\n== Submitting interchain transaction.. ==\n\n"

icad tx intertx submit \
"{
    \"@type\":\"/cosmos.bank.v1beta1.MsgSend\",
    \"from_address\": \"$ICA_ADDR\",
    \"to_address\":\"cosmos10h9stc5v6ntgeygf5xf945njqq5h32r53uquvw\",
    \"amount\": [
        {
            \"denom\": \"stake\",
            \"amount\": \"1000\"
        }
    ]
}" --connection-id connection-0 --from $WALLET_1 --chain-id test-1 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y 1>/dev/null 2>&1

# Wait until the relayer has relayed the packet, adjust as needed
sleep 10
printf "\n== Interchain transaction submitted ==\n\n"

# Query the interchain account balance on the host chain
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657
