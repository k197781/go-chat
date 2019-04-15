package main

type room struct {
	// message sent to other clients is added to forward field
	forward chan []byte
	// channel for user will join is added to join filed
	join chan *client
	// channel for user will join is added to join filed
	leave chan *client
	// all active clients is added to clients field
	clients map[*client]bool
}

// goroutine is perfomed in the background
func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// send massage for all clients
			for client := range r.clients {
				select {
				case client.send <- msg:
				default:
					// if a failure occurs
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}
