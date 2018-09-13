package handler

import (
	"nnmicro/model"
	"nnmicro/pb"

	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// Learning handler create a ws stream to send messages about learning evolution
func Learning(mem model.Memory) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO work in progress
		// Simulation, the neural network package is not finished yet
		// We will simulate sending of random weights and outputs
		nWeigths := 100
		nOutputs := 10

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("Failed to upgrade:", err)
			return
		}
		defer conn.Close()

		message := &pb.Learning{}

		message.IdNN = 1
		message.Outputs = make([]float64, nOutputs)
		message.Weights = make([]float64, nWeigths)

		ticker := time.NewTicker(1000 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {

			case <-ticker.C:

				for i := range message.Outputs {
					message.Outputs[i] = rand.Float64()
				}

				for i := range message.Weights {
					message.Weights[i] = rand.Float64()
				}

				out, err := proto.Marshal(message)
				if err != nil {
					log.Fatalln("Failed to encode learning message:", err)
				}

				err = conn.WriteMessage(websocket.BinaryMessage, out)
				if err != nil {
					log.Println("Failed to write learning message:", err)
					return
				}
			}
		}
	})
}
