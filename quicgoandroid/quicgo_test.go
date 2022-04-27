package main

import (
	"encoding/json"
	"quicgosdk/quicgo"
	"testing"
)

const addr = "127.0.0.1:41420"

func echoStreamServerInit(t *testing.T) {
	t.Logf("%s\n", quicgo.Listen(addr))

	var acceptReturn quicgo.connectReturn
	err := json.Unmarshal([]byte(quicgo.Accept()), &acceptReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(acceptReturn)
	connectID := acceptReturn.connectID

	var streamReturn quicgo.streamReturn
	err = json.Unmarshal([]byte(quicgo.AcceptStream(connectID)), &streamReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(streamReturn)
	streamID := streamReturn.StreamID

	var dataReturn quicgo.dataReturn
	err = json.Unmarshal([]byte(quicgo.ReadStream(streamID)), &dataReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(dataReturn)

	quicgo.WriteStream(streamID, dataReturn.Data)
}

func TestStreamCommunicate(t *testing.T) {
	go func() { echoStreamServerInit(t) }()
	var dialReturn quicgo.connectReturn
	err := json.Unmarshal([]byte(quicgo.Dial(addr)), &dialReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(dialReturn)

	var streamReturn quicgo.streamReturn
	err = json.Unmarshal([]byte(quicgo.OpenStreamSync(dialReturn.connectID)), &streamReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(streamReturn)

	message := "1234567"
	quicgo.WriteStream(streamReturn.StreamID, message)

	var dataReturn quicgo.dataReturn
	err = json.Unmarshal([]byte(quicgo.ReadStream(streamReturn.StreamID)), &dataReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(dataReturn)

	if dataReturn.Data != message {
		t.Errorf("expect %s, receive %s", message, dataReturn.Data)
	}

	quicgo.Close()
}

func echoPacketServerInit(t *testing.T) {
	t.Logf("%s\n", quicgo.Listen(addr))

	var acceptReturn quicgo.connectReturn
	err := json.Unmarshal([]byte(quicgo.Accept()), &acceptReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(acceptReturn)
	connectID := acceptReturn.connectID

	var dataReturn quicgo.dataReturn
	err = json.Unmarshal([]byte(quicgo.ReceiveMessage(connectID)), &dataReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(dataReturn)

	quicgo.SendMessage(connectID, dataReturn.Data)
}

func TestPacketCommunicate(t *testing.T) {
	go func() { echoPacketServerInit(t) }()
	var dialReturn quicgo.connectReturn
	err := json.Unmarshal([]byte(quicgo.Dial(addr)), &dialReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(dialReturn)

	message := "1234567"
	t.Log(quicgo.SendMessage(dialReturn.connectID, message))

	var dataReturn quicgo.dataReturn
	err = json.Unmarshal([]byte(quicgo.ReceiveMessage(dialReturn.connectID)), &dataReturn)
	if err != nil {
		t.Error(err)
	}
	t.Log(dataReturn)

	if dataReturn.Data != message {
		t.Errorf("expect %s, receive %s", message, dataReturn.Data)
	}

	quicgo.Close()
}
