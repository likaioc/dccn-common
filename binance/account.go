package binance

import (
    "github.com/binance-chain/go-sdk/keys"
    log "github.com/sirupsen/logrus"
)

type Account struct {
    Private string
    Address string
}

func CreatAccount(num int) (accounts []Account, addresses []string, err error) {
    var ch = make(chan int)
    for i := 0; i < num; i++ {
        go func() {
            KeyManager, err := keys.NewKeyManager()
            if err != nil {
                log.Errorf("NewKeyManager error", err)
            }
            priv, _ := KeyManager.ExportAsPrivateKey()
            addr := KeyManager.GetAddr().String()
            accounts = append(accounts, Account{priv, addr})
            addresses = append(addresses, addr)
            <-ch
        }()
        ch <- i
    }
    close(ch)
    return accounts, addresses, err
}
