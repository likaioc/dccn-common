# Wallet API

## DESCRIPTION
    * wallet APIs to for clients to talk to blockchain
    * account query and token send APIs will fail if there is no blockchain or wrong ip/port
    * sample code is at 
[ANKR blockchain API samples](https://github.com/Ankr-network/dccn-tendermint/tree/release/v0.28.1/abci/ankrchain/api_sample)

## Usage
    * import "github.com/Ankr-network/dccn-common/wallet"
    * follow the sample code to call,

- wallet.GenerateKeys() 
- wallet.GetBalance() 
- wallet.SendCoins() 
- wallet.Sign()
- wallet.GetPublicKeyByPrivateKey()
- wallet.GetAddressByPublicKey()
- wallet.GetStake()
- wallet.GetHistorySend()
- wallet.GetHistoryReceive()


Admin:

- wallet.SetStake()
- wallet.SetBalance()
