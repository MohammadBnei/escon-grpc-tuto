package services

import (
	"app/proto"
	"context"
	"log"
)

type HelloServer struct {
	AddedMsg string
}

func (s *HelloServer) SayHello(_ context.Context, hello *proto.Hello) (*proto.Response, error) {
	log.Println(hello.Msg, s.AddedMsg)

	// if err != nil {
	// 	return nil, errors.New("Somethine went wrong")
	// }

	response := &proto.Response{
		Ok: "ok",
	}

	return response, nil
}
