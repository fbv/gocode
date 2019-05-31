package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fbv/gocode/service"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

func setupRpc() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	err := s.RegisterService(new(service.HelloService), "")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/rpc", s)
}

func main() {
	fmt.Println("Starting server...")
	setupRpc()
	fmt.Println("Listen on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
