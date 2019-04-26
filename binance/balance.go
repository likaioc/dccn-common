package binance

import (
    "github.com/binance-chain/go-sdk/client"
    "github.com/binance-chain/go-sdk/client/transaction"
    "github.com/binance-chain/go-sdk/types/msg"
    log "github.com/sirupsen/logrus"
)

func SendBalance(cli client.DexClient, transfer [] msg.Transfer, sync bool) (result *transaction.SendTokenResult, err error) {
    result, err = cli.SendToken(transfer, sync)
    if err != nil {
        log.Errorf("send Error", err)
        return result, err
    }
    return result, err
}
