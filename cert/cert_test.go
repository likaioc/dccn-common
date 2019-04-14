package cert_test

import (
	"testing"
	"fmt"

	certmanager "github.com/Ankr-network/dccn-common/cert"
)

func TestGenerateKeys(t *testing.T) {
	t.Log("Testing GenerateCA")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")

	if err == nil {
		t.Logf("\nCA Cert: %s\n CA Public Key: %s\n", caCert, caKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateServerCert(t *testing.T) {
	t.Log("Testing GenerateServerCert")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}

	cert, privateKey, err := certmanager.GenerateServerCert("myhub.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", cert, privateKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateClientCert(t *testing.T) {
	t.Log("Testing GenerateClientCert")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}

	cert, privateKey, err := certmanager.GenerateClientCert("mydatacenter.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", cert, privateKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateServerAndClientCert(t *testing.T) {
	t.Log("Testing GenerateClientCert")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("CA CERT:")
	fmt.Println(caCert)
	fmt.Println("CA KEY:")
	fmt.Println(caKey)

	scert, sprivateKey, err := certmanager.GenerateServerCert("myhub.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", scert, sprivateKey)
	} else {
		t.Error(err)
	}
	fmt.Println("SERVER CERT:")
	fmt.Println(scert)
	fmt.Println("SERVER KEY:")
	fmt.Println(sprivateKey)

	ccert, cprivateKey, err := certmanager.GenerateClientCert("mydatacenter.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", ccert, cprivateKey)
	} else {
		t.Error(err)
	}
	fmt.Println("CLIENT CERT:")
	fmt.Println(ccert)
	fmt.Println("CLIENT KEY:")
	fmt.Println(cprivateKey)
}
