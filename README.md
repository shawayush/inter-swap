
An Implimentation of 
1. ICQ for querying swap in OSMOSIS DEX pool
2. ICA for swapping ICA in OSMOSIS DEX

## Setup
1. Install golang 1.19 +
2. Install hermes
3. git clone this repository
4. git clone osmosis (preferable v14)

We need to create a local osmosis chain for dex and IBC some token through these repository, so that we can create a succesful DEX

### Starting osmosis
```
1. make build
2. cd build
3. ./osmosisd keys add my_validator --keyring-backend test    //save the address and mnemonic
4. MY_VALIDATOR_ADDRESS=$(./osmosisd keys show my_validator -a --keyring-backend test)
5. ./osmosisd init shaw1 --chain-id my-test-chain1
6. ./osmosisd add-genesis-account $MY_VALIDATOR_ADDRESS 100000000000stake
7. ./osmosisd gentx my_validator 100000000stake --chain-id my-test-chain1 --keyring-backend test
8. ./osmosisd collect-gentxs
9. change all the usomo to stake in $HOME/.osmosisd/config/genesis.json also low down stakes for modules such as GAMM
10. change turn on swagger and API in app.toml
11. set chain-id = "my-test-chain1" and keyring-backend = "test" in client.toml at .osmosisd/config
12. ./osmosisd start
```

### Starting knstld (in different terminal)
```
1. make build
2. cd build
3. ./knstld keys add my_validator --keyring-backend test
4. MY_VALIDATOR_ADDRESS=$(./knstld keys show my_validator -a --keyring-backend test)
5. ./knstld init shaw2 --chain-id my-test-chain2
6. ./knstld add-genesis-account $MY_VALIDATOR_ADDRESS 100000000000udarc
7. ./knstld gentx my_validator 100000000udarc --chain-id my-test-chain2 --keyring-backend test
8. ./knstld collect-gentxs
9. change all the stake to uosmo in $HOME/.knstld/config/genesis.json also low down stakes for modules
10. change turn on swagger and API in app.toml and change ports to 1317-->1318(API), 9090-->9090(gRPC), 9091-->9081(gRPC web)
11. change ports in config.toml as 26658-->26659(proxy.app), 26657-->26668(RPC server,) 6060-->6061(pprof_laddr), 26656-->2666(laddr)
12. set chain-id = "my-test-chain2" and keyring-backend = "test" in client.toml at .knstld/config
13. ./knstld start
```

Create new wallets in both the chains and send tokens in the respective wallets - 
```
./knstld keys add <wallet name>
knstld tx bank send [from_key_or_address] [to_address] [amount]
```

### Setup hermes for relaying
1. setup kets
```
echo word1 ... word12or24 > mnemonic_file_hub_osmo
hermes keys add --key-name keyosmo --chain my-test-chain2--mnemonic-file mnemonic_file_hub_osmo.json
echo word1 ... word12or24 > mnemonic_file_hub_darc
hermes keys add --key-name keydarc --chain my-test-chain1--mnemonic-file mnemonic_file_hub_darc.json
```
2. configure hermers in $HOME/.hermes/config.toml (you can copy and paste if you ahve followed the above configuration)
```
[global]
log_level = 'info'

[mode]

[mode.clients]
enabled = true
refresh = true
misbehaviour = true

[mode.connections]
enabled = true

[mode.channels]
enabled = true

[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = true

[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001

[[chains]]
id = 'my-test-chain2'
rpc_addr = 'http://localhost:26668'
grpc_addr = 'http://localhost:9080'
websocket_addr = 'ws://localhost:26668/websocket'
rpc_timeout = '15s'
account_prefix = 'darc'
key_name = 'keydarc'
store_prefix = 'ibc'
gas_price = { price = 0.00, denom = 'udarc' }
max_gas = 10000000
clock_drift = '5s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }

[[chains]]
id = 'my-test-chain1'
rpc_addr = 'http://localhost:26657'
grpc_addr = 'http://localhost:9090'
websocket_addr = 'ws://localhost:26657/websocket'
rpc_timeout = '15s'
account_prefix = 'osmo'
key_name = 'keyosmosis'
store_prefix = 'ibc'
gas_price = { price = 0.00, denom = 'stake' }
max_gas = 10000000
clock_drift = '5s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
```
3. run 'hermes health-check' to check to check things are configure well
4. 'hermes start' to start running hermes
5. need to set up port and channel-  (You can check hermes documentation of you have an doubts)
```
 hermes create connection --a-chain my-test-chain2 --b-chain my-test-chain1
 hermes create channel --order unordered --a-chain my-test-chain2 --a-connection connection-0 --a-port  transfer --b-port transfer
```
6. now chaines are ready to trasnfer IBC
7. transfer tokens to osmosis running chain
```
./knstld tx ibc-transfer transfer transfer channel-0 osmo1s3rnvyqaxm8u5e33z5vj3v7ehc6qhgkkr6ktnd 1000darc --from shaw2 -fees 2udarc
```
### Setup GAMM
./osmosisd tx gamm create-pool --pool-file data.json --from shaw1 --chain-id my-test-chain1
data.json as -
```
{
 "weights": "5ibc/99BA144DEE4A6BA7DF7497ABA9AC7A69F8B1451E98199870B1E738751A59C798,5stake",
 "initial-deposit": "4ibc/99BA144DEE4A6BA7DF7497ABA9AC7A69F8B1451E98199870B1E738751A59C798,5stake",
 "swap-fee": "0.002",
 "exit-fee": "0",
 "future-governor": ""
}

```