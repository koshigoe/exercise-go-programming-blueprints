package main

type room struct {
	// forward は他のクライアントに転送するたえのメッセージを保持するチャネルです。
	forward chan []byte
	// join はチャットルームに参加しようとしているクライアントのためのチャネルです。
	join chan *client
	// leave はチャットルームから退室しようとしているクライアントのためのチャネルです。
	leave chan *client
	// clients には在室しているすべてのクライアントが保持されます。
	clients map[*client]bool
}
