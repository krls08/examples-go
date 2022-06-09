package main

import (
	"fmt"
	"net/http"

	"github.com/krls08/examples-go/ex_ws_nhoory_2/ws"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/ws", ws.ServeWs)

	http.ListenAndServe(":8090", nil)
}
