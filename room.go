package main

type room struct {
	// message sent to other clients is added to forward field.
	forward chan []byte
}
