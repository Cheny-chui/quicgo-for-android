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
	config := quic.Config{
		EnableDatagrams: true,
	}
	conn, err := quic.DialAddr(addr, tlsConf, &config)
	if err != nil {
		return myMarshal(ConnectReturn{Error: err.Error(), ConnectID: 0})
	} else {
		connections[len(connections)+1] = &conn
		return myMarshal(ConnectReturn{Error: "", ConnectID: len(connections)})
	}
}

func OpenStreamSync(connectionID int) string {
	conn := connections[connectionID]
	if conn == nil {
		return myMarshal(StreamReturn{Error: "Server MUST listen first before accept.", StreamID: 0})
	}
	stream, err := (*conn).OpenStreamSync(context.Background())
	if err != nil {
		return myMarshal(StreamReturn{Error: err.Error(), StreamID: 0})
	} else {
		streams[len(streams)+1] = &stream
		return myMarshal(StreamReturn{Error: "", StreamID: len(streams)})
	}
}
