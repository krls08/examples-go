package ws

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type jm struct {
	K  string `json:"k"`
	K2 string `json:"k2"`
}

func ServeWs(w http.ResponseWriter, req *http.Request) {
	fmt.Println("ws server start")
	//fmt.Fprintf(w, "hello\n")
	c, err := websocket.Accept(w, req, nil)
	if err != nil {
		// ...
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	ctx, cancel := context.WithTimeout(req.Context(), time.Second*1000)
	defer cancel()

	var v interface{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		// ...
	}

	log.Printf("received: %v", v)

	c.Close(websocket.StatusNormalClosure, "")
}
