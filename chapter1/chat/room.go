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

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true
		case client := <-r.leave:
			// 退室
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// すべてのクライアントにメッセージを転送
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
