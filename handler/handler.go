package handler

import (
	"cbsignal/client"
	"encoding/json"
)

type Handler interface {
	Handle()
}

type SignalMsg struct {
	Action     string          `json:"action"`
	To_peer_id string          `json:"to_peer_id"`
	Data  interface{}          `json:"data"`
}

type SignalResp struct {
	Action string              `json:"action"`
	FromPeerId string          `json:"from_peer_id"`
	Data interface{}           `json:"data,omitempty"`
}

func NewHandler(message []byte, cli *client.Client) (Handler, error) {
	signalMsg := SignalMsg{}
	if err := json.Unmarshal(message, &signalMsg); err != nil {
		//log.Println(err)
		return nil, err
	}
	switch signalMsg.Action {
	case "signal":
		return &SignalHandler{Msg: &signalMsg, Cli: cli}, nil
	default:
		return &ExceptionHandler{Msg: &signalMsg, Cli: cli}, nil
	}
}
