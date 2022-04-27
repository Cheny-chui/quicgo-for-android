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

var listener *quic.Listener

func Listen(addr string) string {
	config := quic.Config{
		EnableDatagrams: true,
	}
	tempListener, err := quic.ListenAddr(addr, generateTLSConfig(), &config)
	if err != nil {
		return myMarshal(errorReturn{Error: err.Error()})
	} else {
		listener = &tempListener
		return ""
	}
}

func Close() string {
	if listener == nil {
		return myMarshal(errorReturn{Error: "Server MUST listen first before close."})
	}
	err := (*listener).Close()
	if err != nil {
		return myMarshal(errorReturn{Error: err.Error()})
	}
	return myMarshal(errorReturn{Error: ""})
}

func Accept() string {
	if listener == nil {
		return myMarshal(connectReturn{Error: "Server MUST listen first before accept.", ConnectID: 0})
	}
	conn, err := (*listener).Accept(context.Background())
	if err != nil {
		return myMarshal(connectReturn{Error: err.Error(), ConnectID: 0})
	} else {
		connections[len(connections)+1] = &conn
		return myMarshal(connectReturn{Error: "", ConnectID: len(connections)})
	}
}

func AcceptStream(connectID int) string {
	conn := connections[connectID]
	if conn == nil {
		return myMarshal(streamReturn{Error: "Can't find the target connection.", StreamID: 0})
	}
	stream, err := (*conn).AcceptStream(context.Background())
	if err != nil {
		return myMarshal(streamReturn{Error: err.Error(), StreamID: 0})
	} else {
		streams[len(streams)+1] = &stream
		return myMarshal(streamReturn{Error: "", StreamID: len(streams)})
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
