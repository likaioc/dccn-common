package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"net/rpc"
)

const (
        address     = "127.0.0.1:50051"

	SERVER_CERT = `
-----BEGIN CERTIFICATE-----
MIICGTCCAcCgAwIBAgIUDdbnNF2SmmBmitcmd5W3jBqvlPgwCgYIKoZIzj0EAwIw
dDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEUMBIGA1UE
CRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQKEwVIVUJDQTEV
MBMGA1UEAxMMbXlodWItY2EuY29tMB4XDTE5MDUxMjAxNDY1NVoXDTI5MDUxMjAx
NDY1NVowbzELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEU
MBIGA1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQwwCgYDVQQKEwNI
VUIxEjAQBgNVBAMTCW15aHViLmNvbTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IA
BDceYDhhJ+RCjko3ukllR8Iu+YmHZ+KlJR126pzHrS5vpePTJTXJJOomsrHRpoRN
wPecG89K2vr9G97Jpgmpp16jNTAzMA4GA1UdDwEB/wQEAwIHgDATBgNVHSUEDDAK
BggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMAoGCCqGSM49BAMCA0cAMEQCICr2Ug9e
JXt6KQTdY/83r8jeqiCtaT0GRcxld4IvFJS/AiBrUFLsKvv5pebAZN11wgG5MH0K
oqNA8mabaOO/iihdWQ==
-----END CERTIFICATE-----`

	SERVER_KEY = `
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIGvPS4I+OY0pT2a3Ln+SIw1RDMW3IoMzV8t2xmKe14yQoAoGCCqGSM49
AwEHoUQDQgAENx5gOGEn5EKOSje6SWVHwi75iYdn4qUlHXbqnMetLm+l49MlNckk
6iaysdGmhE3A95wbz0ra+v0b3smmCamnXg==
-----END EC PRIVATE KEY-----`

	CA_CERT = `
-----BEGIN CERTIFICATE-----
MIICKzCCAdKgAwIBAgIUW56lhwrEBMk7QKWYY/BZDAl3mLIwCgYIKoZIzj0EAwIw
dDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEUMBIGA1UE
CRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQKEwVIVUJDQTEV
MBMGA1UEAxMMbXlodWItY2EuY29tMB4XDTE5MDUxMjAxNDY1NVoXDTI5MDUxMjAx
NDY1NVowdDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEU
MBIGA1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQKEwVI
VUJDQTEVMBMGA1UEAxMMbXlodWItY2EuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0D
AQcDQgAEhIHiabGkTozh88/+TZrcAH0R5x5CrDN+Y4Czvt5AqqEfFXwU5Ihtt8a1
Pj87hqc6rQVYkwwc8Dgj3u60JgnZQaNCMEAwDgYDVR0PAQH/BAQDAgKEMB0GA1Ud
JQQWMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MAoGCCqG
SM49BAMCA0cAMEQCIAYmfH43DBO8956XxOh+T3Dr/ijf0QmsIE1hb7jlh1MQAiAd
Lu3tEHLLmzQlal9FYxReurigU2kk0fngzbN7HAf8zQ==
-----END CERTIFICATE-----`
)

func main() {
	server(SERVER_CERT, SERVER_KEY, CA_CERT)
}

func server(scert, skey, cacert string) {
	if err := rpc.Register(new(Foo)); err != nil {
		log.Fatal("Failed to register RPC method")
	}
	cert, err := tls.X509KeyPair([]byte(scert), []byte(skey))
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(cacert))
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	config.Rand = rand.Reader
	listener, err := tls.Listen("tcp", address, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	rpc.ServeConn(conn)
	log.Println("server: conn: closed")
}

type Foo bool

type Result struct {
	Data int
}

func (f *Foo) Bar(args *string, res *Result) error {
	res.Data = len(*args)
	log.Printf("Received %q, its length is %d", *args, res.Data)
	return nil
}
