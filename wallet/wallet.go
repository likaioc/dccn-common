package wallet

import (
    "fmt"
    "encoding/hex"
    "encoding/base64"
    "github.com/tendermint/tendermint/crypto/ed25519"
    _ "github.com/tendermint/tendermint/crypto"
    _ "encoding/json"
)

const PubKeyEd25519Size = 32
const PrivKeyEd25519Size = 64
const PubkeyStart = "PubKeyEd25519{"

/*
  Generate new key-pair. return priv_key, pub_key, address in string format.
  Caller does not need to care about the encoding, and API will handle internally.

  example output for priv_key, pub_key, address:
  mB3bPsj9kUsAPmFw11Gqb6AYi7nQ8PFrNI+G62IRyYnstkObfl1KIeQi8pOfpNovC2iikxivhW9baCLStM2hyA==
  7LZDm35dSiHkIvKTn6TaLwtoopMYr4VvW2gi0rTNocg=
  0D9FE6A785C830D2BE66FE40E0E7FE3D9838456C
*/
func GenerateKeys() (priv_key_base64, pub_key_base64, address string) {
        mykey := ed25519.GenPrivKey()
        privArray := [PrivKeyEd25519Size]byte(mykey)
        privBytes := privArray[:]
        privB64 := base64.StdEncoding.EncodeToString(privBytes)
        priv_key_base64 = privB64

        mystr := fmt.Sprintf("%s", mykey.PubKey())
        mystr = mystr[len(PubkeyStart):len(PubkeyStart) + PrivKeyEd25519Size]

        src := []byte(mystr)
        dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)

        pubB64 := base64.StdEncoding.EncodeToString([]byte(dst))
        pub_key_base64 = pubB64

        address = fmt.Sprintf("%s", mykey.PubKey().Address())
        return
}

/*
 do sha256 and sign
 metering feature needs this.
*/
func doSha256Sign(input, priv_key string) (result string) {
	//sum := sha256.Sum256([]byte(input))
	//encrypted_result, _ := priv.Sign([]byte(string(sum[:32])))
	//encrypted_result_b64 := base64.StdEncoding.EncodeToString([]byte(encrypted_result))
	//result = encrypted_result_b64

	result = ""
	return
}

func QueryBalanceByAddress(address string) (balance int) {
	return 0
}

func SendCoins(priv_key, from_address, to_address, amount, public_key string) (result int) {
	return 0
}


