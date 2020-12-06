package server

import (
	"github.com/athomecomar/storeql/name"
	"github.com/athomecomar/xerrors"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sebach1/coding-challenge/microservices/backend/people/filmsconf"
	"google.golang.org/grpc/status"
)

func ConnDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", filmsconf.GetDATABASE_SRC())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "sqlx.Open: %v", err)
	}
	db.MapperFunc(name.ToSnakeCase)
	return db, nil
}
