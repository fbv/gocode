package service

import (
	"fmt"
	"net/http"
)

type HelloArgs struct {
	Name string
}

type HelloReply struct {
	Message string
}

type HelloService struct {
}

func (s *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
	fmt.Println("HelloService.Say executed with args:", args.Name)
	reply.Message = "Hello " + args.Name
	return nil
}
