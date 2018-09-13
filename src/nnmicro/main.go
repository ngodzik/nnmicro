package main

import (
	"fmt"
	"log"
	"net/http"

	"nnmicro/handler"

	"github.com/dimfeld/httptreemux"
	"github.com/justinas/alice"
	//"github.com/ngodzik/ann"
)

func main() {

	// TODO middleware: authentication, ...
	middleware := alice.New()

	router := httptreemux.NewContextMux()

	group := router.NewGroup("/api")

	// Rest api for creating, updating networks
	// TODO listing, run, backprop, ...

	// Learning streaming
	// TODO in memory parameter (redis) is not used yet
	group.GET("/v1/nn/learning", middleware.Then(handler.Learning(nil)).(http.HandlerFunc))

	// Create new neural network
	//TODO
	group.POST("/v1/nn", func(w http.ResponseWriter, r *http.Request) {
		params := httptreemux.ContextParams(r.Context())
		id := params["id"]
		fmt.Fprintf(w, "POST /api/v1/nn/%s", id)
	})

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
