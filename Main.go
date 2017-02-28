/*
	Created on: 30 Jan 2017
	Author: Nilav

	Starts server on default port and registers different handlers
	The mosquitto client is created before starting the server
*/
package main

import (
	Mb "FutsalApi/programs/messageBroker"
	"net/http"
	"FutsalApi/programs/handlers"
)

func main() {

	Mb.CreateClient()

	//ListenAndServe listens on the TCP network address addr
	// and then calls Serve with handler to handle requests
	// on incoming connections.
	http.ListenAndServe(":8080", handlers.Handler())
}
