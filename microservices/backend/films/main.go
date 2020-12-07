package main

import (
	"log"
	"net"

	"github.com/sebach1/coding-challenge/microservices/backend/films/server"
	"github.com/sebach1/coding-challenge/microservices/pb/pbconf"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Films
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)
	pbfilms.RegisterFilmsServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
