package main

import (
	"github.com/gorilla/websocket"
)

// client はチャットを行っている1人のユーザーを表します。
type client struct {
	// socket はこのクライアントのための WebSocket です。
	socket *websocket.Conn
	// send はメッセージが送られるチャネルです。
	send chan []byte
	// room はこのクライアントが参加しているチャットルームです。
	room *room
}
