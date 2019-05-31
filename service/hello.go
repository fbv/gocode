package service

import (
	"fmt"
	"net/http"
)

type HelloArgs struct {
	name string
}

type HelloReply struct {
	message string
}

type HelloService struct {
}

func (s *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	fmt.Println("HelloService.Say executed with args:", args.name)
	reply.message = "Hello " + args.name
	return nil
}
