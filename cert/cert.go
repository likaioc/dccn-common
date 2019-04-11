package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	_ "math/big"
	_ "os"
	"time"
	_ "fmt"
	"bytes"
	"github.com/hashicorp/vault/helper/certutil"
)

// create a CA
// input: name
// output: cert, priv_key
func GenerateCA(name string) (cert string, priv_key string, err error) {
	sn, err := certutil.GenerateSerialNumber()
	if err != nil {
		return "", "", err
	}

	ca := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"DEV"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"mission st."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Println("create ca failed", err)
		return "", "", err
	}

	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	//fmt.Println(certOut.String())

	keyOut := new(bytes.Buffer)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}
