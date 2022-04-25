package quicgosdk

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"github.com/lucas-clemente/quic-go"
	"math/big"
)

func BuildServer(addr string) string {
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		return myMarshal(QReturn{err.Error(), ""})
	}
	conn, err := listener.Accept(context.Background())
	if err != nil {
		return myMarshal(QReturn{err.Error(), ""})
	}
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	data := make([]byte, 1024)
	n, err := stream.Read(data)
	if err != nil {
		return myMarshal(QReturn{err.Error(), ""})
	} else {
		return myMarshal(QReturn{"", string(data[0:n])})
	}
}

func BuildClient(addr string, message string) string {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-example"},
	}
	conn, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		return myMarshal(QReturn{err.Error(), ""})
	}

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		return myMarshal(QReturn{err.Error(), ""})
	}

	_, err = stream.Write([]byte(message))
	if err != nil {
		return myMarshal(QReturn{err.Error(), ""})
	}
	return myMarshal(QReturn{"", string(message)})
}

type QReturn struct {
	Err  string `json:"error"`
	Data string `json:"data"`
}

func myMarshal(qReturn QReturn) string {
	b, err := json.Marshal(qReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-example"},
	}
}
