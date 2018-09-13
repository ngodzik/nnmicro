package handler

import (
	"nnmicro/model"
	"nnmicro/pb"

	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// TODO
func Create(mem model.Memory) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		nn := &pb.NeuralNetwork{}

		bytes, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Fatalln("Failed to retreive request body:", err)
		}

		if err = proto.Unmarshal(bytes, nn); err != nil {
			log.Fatalln("Failed to parse neural network:", err)
		}

		mem.CreatePerceptron(nn.Layer)

	})
}
