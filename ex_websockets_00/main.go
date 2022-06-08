package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IncMssg struct {
	K1 string `json:"k1"`
	K2 string `json:"k2"`
}

var upgrader = websocket.Upgrader{
	//check origin will check the cross region source (note : please not using in production)
	CheckOrigin: func(r *http.Request) bool {
		//Here we just allow the chrome extension client accessable (you should check this verify accourding your client source)
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/api/ws", func(c *gin.Context) {
		//upgrade get request to websocket protocol
		//on connection:
		fmt.Println("connected from", c.Params)
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		for {
			//Read Message from client
			fmt.Println("___>> p0")
			mt, message0, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("message type =>", mt)
			var message IncMssg
			err = json.Unmarshal([]byte(message0), &message)

			//err = c.BindJSON(imsg)
			//err = json.Unmarshal(message, imsg)
			//err = conn.ReadJSON(&message)

			//fmt.Println("----------- message type:", mt)
			fmt.Println("----------- message content bin:", message)
			//fmt.Println("----------- message content str:", string(message))
			//If client message is ping will return pong
			if false {
				//		if string(message) == "ping" {
				//		message = []byte("pong")
			} else {
				//	message = []byte("Your message is: " + string(message))
				if true {
					//imsg := incMssg{}
					fmt.Println("___>> p1")
					//	err := json.Unmarshal(message, &jr)
					if err != nil {
						fmt.Println("err found =>", err.Error())
						//				log.Fatal(err)
					}
					fmt.Println("imsg =>", message)
					fmt.Println("incMssg K1:", message.K1)
					fmt.Println("incMssg K2:", message.K2)
				}
			}
			//Response message to client
			msg, err := json.Marshal(message)
			err = conn.WriteMessage(1, msg)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	})
	r.Run(":8000") // listen and serve on 0.0.0.0:8080
}
