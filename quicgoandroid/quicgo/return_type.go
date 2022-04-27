package quicgo

import "encoding/json"

func myMarshal(r qReturn) string {
	return r.marshal()
}

type qReturn interface {
	marshal() string
}

type errorReturn struct {
	Error string `json:"error"`
}

func (listenReturn errorReturn) marshal() string {
	b, err := json.Marshal(listenReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type connectReturn struct {
	ConnectID int    `json:"connect_id"`
	Error     string `json:"error"`
}

func (acceptReturn connectReturn) marshal() string {
	b, err := json.Marshal(acceptReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type streamReturn struct {
	StreamID int    `json:"stream_id"`
	Error    string `json:"error"`
}

func (acceptStreamReturn streamReturn) marshal() string {
	b, err := json.Marshal(acceptStreamReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type dataReturn struct {
	Data  string `json:"data"`
	Error string `json:"error"`
}

func (dataReturn dataReturn) marshal() string {
	b, err := json.Marshal(dataReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
