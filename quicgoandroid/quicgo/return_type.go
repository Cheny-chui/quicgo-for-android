package quicgo

import "encoding/json"

func myMarshal(r qReturn) string {
	return r.marshal()
}

type qReturn interface {
	marshal() string
}

type ListenReturn struct {
	Error string `json:"error"`
}

func (listenReturn ListenReturn) marshal() string {
	b, err := json.Marshal(listenReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type AcceptReturn struct {
	ConnectID int    `json:"connect_id"`
	Error     string `json:"error"`
}

func (acceptReturn AcceptReturn) marshal() string {
	b, err := json.Marshal(acceptReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type AcceptStreamReturn struct {
	StreamID int    `json:"stream_id"`
	Error    string `json:"error"`
}

func (acceptStreamReturn AcceptStreamReturn) marshal() string {
	//TODO implement me
	b, err := json.Marshal(acceptStreamReturn)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
