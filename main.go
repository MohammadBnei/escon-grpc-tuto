package main

import (
	"app/proto"
	"app/services"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.Println("Server starting now...")

	port := "3000"

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()

	s := &services.HelloServer{
		AddedMsg: "+ grpc",
	}

	proto.RegisterAppServiceServer(srv, s)
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
