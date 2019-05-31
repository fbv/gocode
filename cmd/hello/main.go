package main

import (
	"fmt"
	"log"

	"github.com/fbv/gocode/service"
)

func main() {
	reply := service.HelloReply{}
	s := service.HelloService{}
	err := s.Say(nil, &service.HelloArgs{Name: "Bob"}, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.Message)
}
