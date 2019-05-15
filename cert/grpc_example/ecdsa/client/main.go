package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/rpc"
)

const (
	address     = "127.0.0.1:50051"

	CLIENT_CERT = `
-----BEGIN CERTIFICATE-----
MIICJzCCAc6gAwIBAgIUQNK8zuB47TrjMK/9apa4+ODmGP8wCgYIKoZIzj0EAwIw
dDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEUMBIGA1UE
CRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQKEwVIVUJDQTEV
MBMGA1UEAxMMbXlodWItY2EuY29tMB4XDTE5MDUxMjAxNDY1NVoXDTI5MDUxMjAx
NDY1NVowfTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEU
MBIGA1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MRMwEQYDVQQKEwpE
YXRhQ2VudGVyMRkwFwYDVQQDExBteWRhdGFjZW50ZXIuY29tMFkwEwYHKoZIzj0C
AQYIKoZIzj0DAQcDQgAEM49mdr428vS5+uHc0wjJBqyQ5n8d0QLra97C40uaEw94
l6RWjMOGbQfHGg6YbZzQ6Zc0qIxf7xu+RX//sTmqCaM1MDMwDgYDVR0PAQH/BAQD
AgeAMBMGA1UdJQQMMAoGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwCgYIKoZIzj0E
AwIDRwAwRAIgUxRoNWAjjyvTmnzU8c8s02g0wZURKGo76kh9LNVXcp4CIBAvaZ5u
Y88YwWeiSVJNBDC6MIcgPLAM4YuLvNjP6M6W
-----END CERTIFICATE-----`

	CLIENT_KEY = `
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAHFNZ8+2UnV72fsnUciUAoHYiBKY+FO7IZoT2TPMUUaoAoGCCqGSM49
AwEHoUQDQgAEM49mdr428vS5+uHc0wjJBqyQ5n8d0QLra97C40uaEw94l6RWjMOG
bQfHGg6YbZzQ6Zc0qIxf7xu+RX//sTmqCQ==
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


func client(ccert, ckey, cacert string) {
	cert, err := tls.X509KeyPair([]byte(ccert), []byte(ckey))
	if err != nil {
		log.Fatalf("client: loadkeys: %s", err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(cacert))
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", address, &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}

	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())
	rpcClient := rpc.NewClient(conn)
	res := new(Result)
	if err := rpcClient.Call("Foo.Bar", "twenty-three", &res); err != nil {
		log.Fatal("Failed to call RPC", err)
	}
	log.Printf("Returned result is %d", res.Data)
}

type Result struct {
	Data int
}

func main() {
	client(CLIENT_CERT, CLIENT_KEY, CA_CERT)
}
