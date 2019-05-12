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

func TestGetPublicKeyByPrivKey(t *testing.T) {
	t.Log("Testing GetPublicKeyByPrivateKey")
	verified := "7LZDm35dSiHkIvKTn6TaLwtoopMYr4VvW2gi0rTNocg="
	address, err := wallet.GetPublicKeyByPrivateKey("mB3bPsj9kUsAPmFw11Gqb6AYi7nQ8PFrNI+G62IRyYnstkObfl1KIeQi8pOfpNovC2iikxivhW9baCLStM2hyA==")
	if err != nil {
		t.Error(err)
	}

	if verified != address {
		t.Error("pubkey result doesn't match!")
	}
}

func TestGetAddressByPublicKey(t *testing.T) {
	t.Log("Testing GetAddressByPublicKey")
	verified := "5003BDAF691E5FCEB88FFD14EC6DF6ACC58DEE8C"
	address, err := wallet.GetAddressByPublicKey("wRRGKfjZ/MfMQ1iImhXIkStokco8UckOfee4/rXaLjQ=")
	if err != nil {
		t.Error(err)
	}

	if verified != address {
		t.Error("Address doesn't match!")
	}
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
