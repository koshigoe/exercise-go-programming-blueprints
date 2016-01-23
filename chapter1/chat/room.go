package main

type room struct {
	// forward は他のクライアントに転送するたえのメッセージを保持するチャネルです。
	forward chan []byte
}
