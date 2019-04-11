package cert_test

import (
	"testing"
	_ "fmt"

	wallet "github.com/Ankr-network/dccn-common/cert"
)

func TestGenerateKeys(t *testing.T) {
	t.Log("Testing GenerateCA")

	privateKey, publicKey, err := wallet.GenerateCA("test-ca.com")

	if err == nil {
		t.Logf("\nPrivate Key: %s\nPublic Key: %s\n", privateKey, publicKey)
	} else {
		t.Error(err)
	}
}
