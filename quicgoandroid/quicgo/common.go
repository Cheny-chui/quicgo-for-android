package quicgo

import "github.com/lucas-clemente/quic-go"

var streams = make(map[int]*quic.Stream)
var connections = make(map[int]*quic.Connection)

func ReadStream(streamID int) string {
	stream := streams[streamID]
	if stream == nil {
		return myMarshal(DataReturn{Error: "Can't find the target stream.", Data: ""})
	}
	data := make([]byte, 1024)
	n, err := (*stream).Read(data)
	if err != nil {
		return myMarshal(DataReturn{Error: err.Error(), Data: ""})
	} else {
		return myMarshal(DataReturn{Error: "", Data: string(data[0:n])})
	}

}

func WriteStream(streamID int, message string) string {
	stream := streams[streamID]
	if stream == nil {
		return myMarshal(ErrorReturn{Error: "Can't find the target stream."})
	}
	_, err := (*stream).Write([]byte(message))
	if err != nil {
		return myMarshal(ErrorReturn{Error: err.Error()})
	} else {
		return myMarshal(ErrorReturn{Error: ""})
	}
}

func ReceiveMessage(connectID int) string {
	conn := connections[connectID]
	if conn == nil {
		return myMarshal(DataReturn{Error: "Can't find the target connection.", Data: ""})
	}
	data, err := (*conn).ReceiveMessage()
	if err != nil {
		return myMarshal(DataReturn{Error: err.Error(), Data: ""})
	}
	return myMarshal(DataReturn{Error: "", Data: string(data)})
}

func SendMessage(connectID int, message string) string {
	conn := connections[connectID]
	if conn == nil {
		return myMarshal(ErrorReturn{Error: "Can't find the target connection."})
	}
	err := (*conn).SendMessage([]byte(message))
	if err != nil {
		return myMarshal(ErrorReturn{Error: err.Error()})
	}
	return myMarshal(ErrorReturn{Error: ""})
}
