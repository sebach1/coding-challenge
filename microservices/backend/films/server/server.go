package server

import "github.com/sebach1/coding-challenge/microservices/pb/pbfilms"

type Server struct {
	*pbfilms.UnimplementedFilmsServer
}
