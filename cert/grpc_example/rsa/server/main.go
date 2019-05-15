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
MIIDpjCCAo6gAwIBAgIUGww4NvBx7czb44X7CKv+pHsSGV8wDQYJKoZIhvcNAQEL
BQAwdDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEUMBIG
A1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQKEwVIVUJD
QTEVMBMGA1UEAxMMbXlodWItY2EuY29tMB4XDTE5MDQxMzIzNDEzM1oXDTI5MDQx
MzIzNDEzM1owbzELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJT
RjEUMBIGA1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQwwCgYDVQQK
EwNIVUIxEjAQBgNVBAMTCW15aHViLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEP
ADCCAQoCggEBALPBKcDT9GvYNAmhxI02fp8MAQJGKqqCHZAGcTmUEoePmq6d1frD
8nVHA+iJPe9teVSJc2JZ1FXXuWdCPzrJlXBldTCLD8PGvYu4tyXAUJL9gb+xsGC0
H4w4z/gPICwD5K9I1oiVf4pe9w5q2oytYQE5zATz6VXcbuiGhbGZTPl7ZMfnrcbj
VJf/kdiShV2uN5/U2EFC3FMi+/z4tdPWVtF9nfqCuiWeuJ4got9CmztARpjqDtkX
knIvCyBD9MSc+00PAhEPEjhx9/lChxXFIwyvsqwCagqLwFJT7YdEAZARR+YnPIxc
rQgLA17Yus1KSCLxnXJ++X5lV9unHq+F7bsCAwEAAaM1MDMwDgYDVR0PAQH/BAQD
AgeAMBMGA1UdJQQMMAoGCCsGAQUFBwMBMAwGA1UdEwEB/wQCMAAwDQYJKoZIhvcN
AQELBQADggEBAI+fHeHFfM/aQjiz1VZYA/QUXYl8dDHsUdf4KLK1OThUSnxz6vlz
uqH4r+Qb8wXpoT5TVQRpNWtBr+B+iGuHnw5otH0phIdcneM5XGMC5LX1jVr5IWMP
Al8qHfFPKjn6YVUFAUcx4YaDroNOEwobCnCFZNYTtMin3nVG1X31VFKkJgCMaY/8
41BJwh3GEsMTBK4lOAr64ZOkVVcuuOLRLl6sKYfK3MRYGDZOdjZwslmon/ioSoNq
MvHxSvidvUMnjJE9sTRM+R4QdaNV8H2jfhU5mrXnelbstJMTFzebskhW0ai7jZow
6v+y6m9UEdPS9YIS8w2wI0oLvPNtdd9COFM=
-----END CERTIFICATE-----`

	SERVER_KEY = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAs8EpwNP0a9g0CaHEjTZ+nwwBAkYqqoIdkAZxOZQSh4+arp3V
+sPydUcD6Ik97215VIlzYlnUVde5Z0I/OsmVcGV1MIsPw8a9i7i3JcBQkv2Bv7Gw
YLQfjDjP+A8gLAPkr0jWiJV/il73DmrajK1hATnMBPPpVdxu6IaFsZlM+Xtkx+et
xuNUl/+R2JKFXa43n9TYQULcUyL7/Pi109ZW0X2d+oK6JZ64niCi30KbO0BGmOoO
2ReSci8LIEP0xJz7TQ8CEQ8SOHH3+UKHFcUjDK+yrAJqCovAUlPth0QBkBFH5ic8
jFytCAsDXti6zUpIIvGdcn75fmVX26cer4XtuwIDAQABAoIBAC6+1e/abt64y9eR
ZP7gJA+SXJTN0Hyk97Ejq8AwfJ0aQDyWuTXMTToobKDFSAECtCHC2OnDuI9WWVW8
CZttHtqq18326/OatGIoeCBuIcYH/Pzv5ZFQZj/d6l9094XawVbAF2JPlmpShecO
7KTW1bou4JqNH6n/eCwqB4yG9FiQSYQvARNthTpuayzblbAaXSbQVBd0vZwAUyeF
BXp8+CLcW/v+Q13lFpQH9q/ZQknvLNmy34ziIMjUpnrCWlTXfY7uSAHyl0kp4QPN
h6iwwwjlttTs8L7oM2nSVWgqx59+uEBOC3TYymQYjG8ZzvOyeQpshWw3Au8UMLIU
RABwjykCgYEAyoCHLDYP+IkAmJJ+Bs2AQeBDzQOf6N58QvUNQAYWb9zFdrKMmUyv
A8YGniwuqKTGLT0LH/6QwWzPO3qpIxR6IqjeJnuziuih+sxqHPbK3LTDMB7k3Ru4
+jSUJKfY1SciUPJwDu+b1020yrCD5oILSl3+1+iLMD9Ds/JC4OOezCUCgYEA4z4x
AaB35yb1ov1ti3DUjVnSwu8I+LOqJSsQ3LBRLJwt2+a6VwpJuBoUetsvMX38bEpv
kJoAgVYLa2BvtTxUBLifIGfKVW5oTCjTna8oAJeCxHP5Vxm9GBHL33yQUrlR3Ziz
K9OTWSWWIIhilmpgu2TSVDXvLChMQfyiQtzpvF8CgYAQPMQZ+G/JZvrkQQLX+sxo
+9mT69F7s7NVDx8z3gkYiSRQQzrg4/Q4oj1Pm41WOnglknp840WqLNZlF43OYUr1
0oKVpV+dKNAsMw3jKqrTja8FAcUqMZngfizOz6KkXNH0mawAlpfxaKqxgfvhlNXg
7a6Mk3ntsXT1Tdys+bN2/QKBgQC1kR9++URUgSCLForMe16ZgkeQUYWNkaBFCl/y
tT5msnNB3NzOZv34B36Nm0vd85gYvnDdzO7SaHm/VdJrsMRA8XLTZZqZIBpdn9FV
kHoZ3vUxGqkDGyGndbC75L88Ga57rGKchfSZqhqF6/M5081ubWRQvLdZROmJX94W
JoTraQKBgQCQ356HyhghNw1m4K/AmMfSdpAXT9iYleKhJclX29C0A2od56OWNjW1
piQh4Ye4GYRpYr8sdZlX2jFD6V/i2FTXUniES0RFj8d+g6eBvAmvfyZiTWDyQ4TW
BsQTGj3gqZgVJt7lx25bXijM6PzZ6kKF4/0FkE51YAV1vFL4TMkP2Q==
-----END RSA PRIVATE KEY-----`

	CA_CERT = `
-----BEGIN CERTIFICATE-----
MIIDuDCCAqCgAwIBAgIUTAytotqCfHmnAl3Nfer0tDQk76cwDQYJKoZIhvcNAQEL
BQAwdDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEUMBIG
A1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQKEwVIVUJD
QTEVMBMGA1UEAxMMbXlodWItY2EuY29tMB4XDTE5MDQxMzIzNDEzM1oXDTI5MDQx
MzIzNDEzM1owdDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJT
RjEUMBIGA1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQK
EwVIVUJDQTEVMBMGA1UEAxMMbXlodWItY2EuY29tMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAqUzfp+xrZLZKFEFqSGtii/9+mBs+DsGp9Cv56//Tmcez
mCh4g06bfE8lr3+A0nHIuqTQZ3zQFgDRReapJ9hEx6hrGxq06fnVlsoTeL+OmJ4O
mPUmIEYX5GorvnUYLwzOGOmjKjdntHajLmfDUlYEKqH5juV6JLhMJi6/6/P4E+jn
hwHGlkiJtaCxC4qdhEVpwxf9qiY/UGlq+xonR6OCPcXCuLeMilP2kV3CBoL28W3s
5bDG1mIFykZL3F2brfUjbQIZ79QhaqbcyLMwhMH8RCPUK3VpSlRkBP5P+QQ4jdxl
Cfahv3nXCKX79vbPWMEDPlpdK3f2e62Umuw66M4ycQIDAQABo0IwQDAOBgNVHQ8B
Af8EBAMCAoQwHQYDVR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMBMA8GA1UdEwEB
/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAIELj4M3We9FVfJgeTDcXv2KspW2
mlcFPOWJrIu0nE5CIqyIq6I0oxwkWWITm+W0n1MhsSBiY0UqcJ7/M6YZIsrftaAp
cs8MAhvLioQe+wOhXuEn9Q3/B4lT7WooKQzLA4/C31qSqaVR1/+mon4Tjj9iBZOt
z1KPh64OOMCVdzdesK/mU7OIk7h889wtaU69KH29SPPMVoqZgD7ltd0SxubnVaQk
HvGkoT8cdI6RHj3MwXuQVDZiNPNbPjxZ083X5yAVo8msiRP9LGv8ZCXTHiakp6Qw
djfzH42ERwtGA6UOtS2z9NyaY90x8j0IwuJGTxEiNMG4vHMXbmZp8Brd5dY=
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
