package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/sebach1/coding-challenge/microservices/backend/films/proxy"
	"github.com/sebach1/coding-challenge/microservices/backend/films/server"
	"github.com/sebach1/coding-challenge/microservices/pb/pbconf"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		svc := pbconf.Films
		port := svc.GetPort()
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		log.Println("listening grpc on port " + port)
		pbfilms.RegisterFilmsServer(s, &server.Server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		port := ":10000"
		httpSrv := proxy.NewProxy(pbconf.Films.GetPort(), port)

		log.Printf("proxy server running at %s ...\n", port)
		if err := httpSrv.Serve(context.Background()); err != nil {
			log.Fatalf("failed to proxy: %v", err)
		}
	}()
	wg.Wait()
}
