package wallet_test

import (
	"testing"

	wallet "github.com/Ankr-network/dccn-common/wallet"
)

func TestGenerateKeys(t *testing.T) {
	t.Log("Testing GenerateKeys")
	privateKey, publicKey, address := wallet.GenerateKeys()

	t.Logf("\nPrivate Key: %s\nPublic Key: %s\nAddress: %s", privateKey, publicKey, address)
}

func TestSign(t *testing.T) {
	t.Log("Testing TestSign")
	privateKey := "vfGLyfQawsZArHk45G45TXCDCobaybILJICYYimTUugVblxJkpsSHl4EnM5YgD41VEoJ+7m82fzTS0EvSwwLDw=="
	strToSign := "test string"
	verified := "owARjBRSY1YNGc36mmWvhxj8IA+icQ0Nd2PNqETUq/NefY8CeqpCvulHEXxvJ76GEFRiVFOkRhUbb881yQhyBg=="
	sig, err := wallet.Sign(strToSign, privateKey)

	if err != nil {
		t.Error(err)
	}

	if verified != sig {
		t.Error("Signature doesn't match!")
	}
}
