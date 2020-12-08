package proxy

import (
	"context"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
	"google.golang.org/grpc"
)

type Proxy struct {
	listen string
	port   string
}

func (s *Proxy) Serve(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pbfilms.RegisterFilmsHandlerFromEndpoint(ctx, mux, s.listen, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(
		s.port,
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(mux),
	)
}

func NewProxy(listen, httpAddr string) *Proxy {
	return &Proxy{listen: listen, port: httpAddr}
}
