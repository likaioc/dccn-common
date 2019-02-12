package wallet

import (
    "fmt"
    "errors"
    "strings"
    "encoding/hex"
    "encoding/base64"
    "crypto/sha256"
    "github.com/tendermint/tendermint/crypto/ed25519"
    "github.com/tendermint/tendermint/rpc/client"
    cmn "github.com/tendermint/tendermint/libs/common"
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
 do sha256 and sign.
 metering feature needs this.
*/
func Sha256Sign(input, priv_key string) (result string, err_ret error) {
        privKeyObject, err := _deserilizePrivKey(priv_key)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256([]byte(input))
	encrypted_result, err := privKeyObject.Sign([]byte(string(sum[:32])))
	if err != nil {
		return "", err
	}

	encrypted_result_b64 := base64.StdEncoding.EncodeToString([]byte(encrypted_result))

	result = encrypted_result_b64
	err_ret = nil

	return
}

/*
 query balance based on (ip, port, address).
 ip: blockchain node ip or domain
 port: blockchain node port, usually 26657
 address: wallet address
 return: balance or error
*/
func QueryBalanceByAddress(ip, port, address string) (balance string, err_ret error) {

	cl := _getHTTPClient(ip, port)

        _, err := cl.Status()
        if err != nil {
		return "", err
	} else { 
                //fmt.Println("LatestBlockHeight:", status.SyncInfo.LatestBlockHeight)
        }

        // curl  'localhost:26657/abci_query?data="bal:1234567890123456789012345678901234567890"'
        res, err := cl.ABCIQuery("/websocket", cmn.HexBytes(fmt.Sprintf("%s:%s", "bal", address)))
        qres := res.Response
	if !qres.IsOK() {
		return "", errors.New("Query balance failure, connect error.")
        } else {
		balanceNonceSlices := strings.Split(string(qres.Value), ":")
		if len(balanceNonceSlices) == 2 {
			balance = balanceNonceSlices[0]
	        } else {
			return "", errors.New("Query balance failure, balance format incorrect.")
		}
	}

	return balance, nil
}

func SendCoins(priv_key, from_address, to_address, amount, public_key string) (result int) {
	return 0
}

/* helper functions */
/*-------------------------------------------------------------------------------------------------*/
func _getHTTPClient(ip, port string) *client.HTTP {
        return client.NewHTTP(ip + ":" + port, "/websocket")
}

func _deserilizePrivKey(priv_key_b64 string) (ed25519.PrivKeyEd25519, error){
        kDec, err := base64.StdEncoding.DecodeString(priv_key_b64)
	if err != nil {
		return ed25519.PrivKeyEd25519{}, err
	}

        pp := []byte(kDec)
        var keyObject ed25519.PrivKeyEd25519 = ed25519.PrivKeyEd25519{pp[0], pp[1], pp[2], pp[3], pp[4], pp[5], pp[6], pp[7], pp[8], pp[9], pp[10], pp[11], pp[12], pp[13], pp[14], pp[15], pp[16], pp[17], pp[18], pp[19], pp[20], pp[21], pp[22], pp[23], pp[24], pp[25], pp[26], pp[27], pp[28], pp[29], pp[30], pp[31], pp[32], pp[33], pp[34], pp[35], pp[36], pp[37], pp[38], pp[39], pp[40], pp[41], pp[42], pp[43], pp[44], pp[45], pp[46], pp[47], pp[48], pp[49], pp[50], pp[51], pp[52], pp[53], pp[54], pp[55], pp[56], pp[57], pp[58], pp[59], pp[60], pp[61], pp[62], pp[PrivKeyEd25519Size - 1]}

        return keyObject, nil
}
 
