package server

import (
	"context"
	"database/sql"

	"github.com/athomecomar/xerrors"
	"github.com/pkg/errors"
	"github.com/sebach1/coding-challenge/microservices/backend/films/ent"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveFilm(ctx context.Context, in *pbfilms.RetrieveFilmRequest) (*pbfilms.RetrieveFilmResponse, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	film, err := ent.FindFilm(ctx, db, in.GetId())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "cant find film by id: %v", film)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindFilm: %v", err)
	}
	return &pbfilms.RetrieveFilmResponse{Data: film.ToPb()}, nil
}

func (s *Server) RetrieveFilms(ctx context.Context, in *pbfilms.RetrieveFilmsRequest) (*pbfilms.RetrieveFilmsResponse, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryxContext(ctx, `SELECT * FROM films LIMIT $1`, in.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()

	films := make(map[uint64]*pbfilms.FilmData)
	for rows.Next() {
		film := &ent.Film{}
		err = rows.StructScan(film)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		films[film.Id] = film.ToPb()
	}
	return &pbfilms.RetrieveFilmsResponse{Films: films}, nil
}

func (s *Server) RetrieveFilmsWithPeople(ctx context.Context, in *pbfilms.RetrieveFilmsWithPeopleRequest) (*pbfilms.RetrieveFilmsWithPeopleResponse, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	filmsWithPeople, err := ent.GetFilmsWithPeople(ctx, db, in.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "GetFilmsWithPeople")
	}
	result := make(map[uint64]*pbfilms.FilmDataWithPeople)
	for film, peopleSl := range filmsWithPeople {
		peopleMap := make(map[uint64]*pbfilms.PeopleData)
		for _, ppl := range peopleSl {
			peopleMap[ppl.Id] = ppl.ToPb()
		}
		result[film.Id] = &pbfilms.FilmDataWithPeople{Film: film.ToPb(), People: peopleMap}
	}
	return &pbfilms.RetrieveFilmsWithPeopleResponse{Films: result}, nil
}
