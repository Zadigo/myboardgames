package main

type WebsocketMessage struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}
