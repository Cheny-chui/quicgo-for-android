package quicgo

import "encoding/json"

func myMarshal(r qReturn) string {
	return r.marshal()
}

type qReturn interface {
	marshal() string
}

type ErrorReturn struct {
	Error string `json:"error"`
}

func (listenReturn ErrorReturn) marshal() string {
	b, err := json.Marshal(listenReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type ConnectReturn struct {
	ConnectID int    `json:"connect_id"`
	Error     string `json:"error"`
}

func (acceptReturn ConnectReturn) marshal() string {
	b, err := json.Marshal(acceptReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type StreamReturn struct {
	StreamID int    `json:"stream_id"`
	Error    string `json:"error"`
}

func (acceptStreamReturn StreamReturn) marshal() string {
	b, err := json.Marshal(acceptStreamReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type DataReturn struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

func (dataReturn DataReturn) marshal() string {
	b, err := json.Marshal(dataReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
