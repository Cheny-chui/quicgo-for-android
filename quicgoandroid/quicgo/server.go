package quicgo

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/lucas-clemente/quic-go"
	"math/big"
)

var listener quic.Listener

func Listen(addr string) string {
	tempListener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		return myMarshal(ListenReturn{Error: err.Error()})
	} else {
		listener = tempListener
		return ""
	}
}

func Accept() string {
	if listener == nil {
		return myMarshal(AcceptReturn{Error: "Server MUST listen first before accept", ConnectID: 0})
	}
	conn, err := listener.Accept(context.Background())
	if err != nil {
		return myMarshal(AcceptReturn{Error: err.Error(), ConnectID: 0})
	} else {
		Connections[len(Connections)+1] = conn
		return myMarshal(AcceptReturn{Error: "", ConnectID: len(Connections)})
	}
}

func AcceptStream(connectID int) string {
	conn := Connections[connectID]
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		return myMarshal(AcceptStreamReturn{Error: err.Error(), StreamID: 0})
	} else {
		Streams[len(Streams)+1] = stream
		return myMarshal(AcceptStreamReturn{Error: "", StreamID: len(Streams)})
	}

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
