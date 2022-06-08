package main

import (
	"fmt"

	"github.com/krls08/examples-go/ex_websockets_01/internal/platform/wsserver"
)

func main() {
	server := wsserver.StartServer(messageHandler)

	for {
		server.WriteMessage([]byte("Hello"))
	}
}

func messageHandler(message []byte) {
	fmt.Println(string(message))
}
