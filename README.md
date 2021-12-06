# Interchain Accounts
### Warning 
> Beware of dragons!</br></br>
> The interchain accounts module is currently under development and has been moved to the `ibc-go` repo [here](https://github.com/cosmos/ibc-go/pull/380). Interchain Accounts is aiming to release in January 2022.</br></br>
> This repo aims to demonstrate demo modules that utilize interchain accounts and serve as a developer guide for teams aiming to use interchain accounts functionality.</br></br>
> The existing demo outlined below is outdated and will be updated to coincide with the interchain accounts relase.</br>

## Local Demo

### Setup

```bash
# Clone this repository and build
git clone https://github.com/cosmos/interchain-accounts.git
cd interchain-accounts
make install 

# Hermes Relayer
# [Hermes](https://hermes.informal.systems/) is a Rust implementation of a relayer for the [Inter-Blockchain Communication (IBC)](https://ibcprotocol.org/) protocol.
#
# In order to use the hermes relayer you will need to check out a specific branch that can be used with interchain-accounts. 
# 
# In the variables.sh file inside /network/hermes/ replace the $HERMES_BINARY variable with a path to the hermes binary generated from the build step below. 
# You can find this in the /target/debug/ directory inside ibc-rs. 

git clone https://github.com/informalsystems/ibc-rs
git checkout adi/ibcgo_v2
cd relayer-cli
cargo build


# Bootstrap two local chains & create a connection using the hermes relayer
make init

# Wait for the ibc connection & channel handshake to complete and the relayer to start
```

### Demo

```bash
# Open a seperate terminal

# Store the following account addresses within the current shell env
export DEMOWALLET_1=$(icad keys show demowallet1 -a --keyring-backend test --home ./data/test-1);
echo $DEMOWALLET_1;

export DEMOWALLET_2=$(icad keys show demowallet2 -a --keyring-backend test --home ./data/test-2);
echo $DEMOWALLET_2;

# Register an interchain account on behalf of DEMOWALLET_1 where chain test-2 is the interchain accounts host
icad tx intertx register --from $DEMOWALLET_1 --connection-id connection-0 --counterparty-connection-id connection-0 --chain-id test-1 --gas 150000 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y

# Start the hermes relayer in the first terminal
# This will also finish the channel creation handshake signalled during the register step
make start-rly

# Query the address of the interchain account
icad query intertx interchainaccounts $DEMOWALLET_1 connection-0 connection-0 --home ./data/test-1 --node tcp://localhost:16657

# Store the interchain account address by parsing the query result
export ICA_ADDR=$(icad query intertx interchainaccounts $DEMOWALLET_1 connection-0 connection-0 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address') && echo $ICA_ADDR

# Check the interchain account's balance on test-2 chain. It should be empty.
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657

# Send some assets to $ICA_ADDR.
icad tx bank send $DEMOWALLET_2 $ICA_ADDR 10000stake --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 --keyring-backend test -y

# Check that the balance has been updated
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657

# Test sending assets from interchain account via ibc.
icad tx intertx send $ICA_ADDR $DEMOWALLET_2 5000stake --connection-id connection-0 --counterparty-connection-id connection-0 --chain-id test-1 --gas 90000 --home ./data/test-1 --node tcp://localhost:16657 --from $DEMOWALLET_1 --keyring-backend test -y

# Wait until the relayer has relayed the packet

# Query the interchain account balance and observe the changes in funds
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657
```

## Collaboration

Please use conventional commits  https://www.conventionalcommits.org/en/v1.0.0/

```
chore(bump): bumping version to 2.0
fix(bug): fixing issue with...
feat(featurex): adding feature...
```
