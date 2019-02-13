# Wallet API

## DESCRIPTION
    * wallet APIs to for clients to talk to blockchain
    * account query and token send APIs will fail if there is no blockchain or wrong ip/port
    * sample code is at 
[ANKR blockchain API samples](https://github.com/Ankr-network/dccn-tendermint/tree/release/v0.28.1/abci/ankrchain/api_sample)

## Usage
    * import "github.com/Ankr-network/dccn-common/wallet"
    * follow the sample code to call,

- wallet.wallet.GenerateKeys(), 
- wallet.QueryBalanceByAddress(), 
- wallet.SendCoins(), 
- wallet.Sha256Sign()
