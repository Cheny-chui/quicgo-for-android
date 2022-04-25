package quicgo

import (
	"context"
	"crypto/tls"
	"github.com/lucas-clemente/quic-go"
)

func Dial(addr string) string {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-example"},
	}
	conn, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {

	} else {

	}
}

func OpenStreamSync(connectionID int) string {
	conn := Connections[connectionID]
	if conn == nil {

	} else {
		stream, err := conn.OpenStreamSync(context.Background())
		if err != nil {

		} else {

		}
	}
}
