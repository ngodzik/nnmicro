package main

// Websocket client test file

import (
	"log"

	"nnmicro/pb"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

func main() {

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/api/v1/nn/learning", nil)
	if err != nil {
		log.Fatal("Failed to dial:", err)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Failed to read:", err)
				return
			}

			learning := &pb.Learning{}

			err = proto.Unmarshal(message, learning)

			if err != nil {
				log.Fatalln("Failed to decode learning message:", err)
			}

			log.Printf("received neural network ID : %d\n\n", learning.IdNN)
			log.Printf("received output : %+v\n\n", learning.Outputs)
			log.Printf("received weights : %+v\n\n", learning.Weights)
			log.Printf("#######################\n\n")
		}
	}()

	<-done
	return
}
