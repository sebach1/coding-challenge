package wrk

import (
	"context"

	"github.com/athomecomar/xerrors"
	"github.com/sebach1/coding-challenge/microservices/pb/pbconf"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ConnFilms(ctx context.Context) (pbfilms.FilmsClient, func() error, error) {
	conn, err := conn(ctx, pbconf.Films.GetHost(), pbconf.Films.GetPort())
	if err != nil {
		return nil, nil, err
	}
	c := pbfilms.NewFilmsClient(conn)
	return c, conn.Close, nil
}

func conn(ctx context.Context, host, port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(host+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, host)
	}
	return conn, nil
}
