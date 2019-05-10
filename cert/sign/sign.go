package sign

import (
	"crypto"
	"crypto/x509"
	"crypto/rsa"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"errors"
	"crypto/sha256"
	"crypto/rand"
	"encoding/pem"
)


/*
use pem format string RSA private_key to sign the input
*/
func RsaSign(privpem, input string) (string, error) {
	sum := sha256.Sum256([]byte(input))
	priv, err := parseRsaPrivateKeyFromPemStr(privpem)
	if (err != nil) {
		fmt.Println(err)
		return "", nil
	}

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	signature, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, sum[:32], &opts)

	if err != nil {
		fmt.Println(err)
		return  "", err
	}

	return string(signature), nil

}

/*
use pem format string ECDSA private_key to sign the input
*/
func EcdsaSign(privpem, input string) (rSig, sSig string, err error) {
	sum := sha256.Sum256([]byte(input))
	priv, err := parseEcdsaPrivateKeyFromPemStr(privpem)
	if (err != nil) {
		fmt.Println(err)
		return "", "", err
	}

	r, s, err := ecdsa.Sign(rand.Reader, priv, sum[:32])

	if err != nil {
		fmt.Println(err)
		return  "", "", err
	}

	return r.String(), s.String(), nil

}

/*
use pem format string RSA public_key to verify the input's signature.
*/
func RsaVerify(pubpem, input, signature string) bool{
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	rsaPublicKey, err := parseRsaPublicKeyFromPemStr(pubpem)
	if (err != nil) {
		fmt.Println(err)
		return false
	}

	sum := sha256.Sum256([]byte(input))
	err = rsa.VerifyPSS(rsaPublicKey, crypto.SHA256, sum[:32], []byte(signature), &opts)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

/*
use pem format string ECDSA public_key to verify the input's signature.
*/
func EcdsaVerify(pubpem, input, signature1, signature2 string) bool{

	rSig := new(big.Int)
	rSig, ok := rSig.SetString(signature1, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}

	sSig := new(big.Int)
	sSig, ok = sSig.SetString(signature2, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false
	}

	ecPublicKey, err := parseEcdsaPublicKeyFromPemStr(pubpem)
	if (err != nil) {
		fmt.Println(err)
		return false
	}

	sum := sha256.Sum256([]byte(input))
	valid := ecdsa.Verify(ecPublicKey, sum[:32], rSig, sSig)

	return valid
}


//=================================================================================
/* helper functions */

func parseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the privkey")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
}

func parseEcdsaPrivateKeyFromPemStr(privPEM string) (*ecdsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the privkey")
    }

    priv, err := x509.ParseECPrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
}

func parseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
        block, _ := pem.Decode([]byte(pubPEM))
        if block == nil {
                fmt.Println("failed to parse certificate PEM")
		return nil, errors.New("failed to parse PEM block containing the cert")
        }
        cert, err := x509.ParseCertificate(block.Bytes)
        if err != nil {
                fmt.Println("failed to parse certificate: ", err.Error())
		return nil, err
        }

	pub := cert.PublicKey.(*rsa.PublicKey)

	return pub, nil
}

func parseEcdsaPublicKeyFromPemStr(pubPEM string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		fmt.Println("failed to parse certificate PEM")
		return nil, errors.New("failed to parse PEM block containing the cert")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("failed to parse certificate: ", err.Error())
		return nil, err
	}

	pub := cert.PublicKey.(*ecdsa.PublicKey)

	return pub, nil
}
