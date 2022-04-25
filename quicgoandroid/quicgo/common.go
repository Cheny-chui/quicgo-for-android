package quicgo

import "github.com/lucas-clemente/quic-go"

var Streams = make(map[int]quic.Stream)
var Connections = make(map[int]quic.Connection)

func ReadStream(streamID int) string {
	stream := Streams[streamID]
	if stream == nil {

	}
}

func WriteStream(streamID int) string {
	stream := Streams[streamID]
	if stream == nil {

	}
}

func SendMessage(connectID int) string {
	conn := Connections[connectID]
	if conn == nil {

	}
}

func ReceiveMessage(connectID int) string {
	conn := Connections[connectID]
	if conn == nil {

	}
}
