package server

import (
	"context"
	"database/sql"
	"io"
	"strconv"

	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/go-pg/pg"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sebach1/coding-challenge/microservices/backend/films/ent"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	defer rows.Close()

	var films []*pbfilms.FilmData
	for rows.Next() {
		film := &ent.Film{}
		err = rows.StructScan(film)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}
		films = append(films, film.ToPb())
	}
	return &pbfilms.RetrieveFilmsResponse{Films: films}, nil
}

func (s *Server) RetrieveFilmsWithPeople(ctx context.Context, in *pbfilms.RetrieveFilmsWithPeopleRequest) (*pbfilms.RetrieveFilmsWithPeopleResponse, error) {
	db, err := ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	limit := in.GetLimit()
	if limit == 0 {
		limit = 250
	}
	filmsWithPeople, err := ent.GetFilmsWithPeople(ctx, db, limit)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "GetFilmsWithPeople: %v", err)
	}
	var result []*pbfilms.FilmDataWithPeople
	for film, peopleSl := range filmsWithPeople {
		peopleMap := make(map[string]*pbfilms.PeopleData)
		for _, ppl := range peopleSl {
			peopleMap[strconv.Itoa(int(ppl.Id))] = ppl.ToPb()
		}
		result = append(result, &pbfilms.FilmDataWithPeople{Film: film.ToPb(), People: peopleMap})
	}
	return &pbfilms.RetrieveFilmsWithPeopleResponse{Films: result}, nil
}

func (s *Server) CreateFilms(srv pbfilms.Films_CreateFilmsServer) error {
	ctx := srv.Context()
	db, err := ConnDB()
	if err != nil {
		return err
	}
	var resp *emptypb.Empty
	for {
		select {
		case <-ctx.Done():
			err = srv.SendAndClose(resp)
			if err != nil {
				return err
			}
			return ctx.Err()
		default: // no-op
		}

		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		resp, err = s.createFilms(ctx, db, in.GetData())
		if err != nil {
			return err
		}
	}

}

func (s *Server) createFilms(ctx context.Context, db *sqlx.DB, in *pbfilms.FilmData) (*emptypb.Empty, error) {
	film := ent.FilmFromPb(in)
	err := storeql.InsertIntoDB(ctx, db, film)
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if ok && !pgErr.IntegrityViolation() { // Ignore integrity violations (to follow up same behavior as w django)
			return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) CreatePeople(srv pbfilms.Films_CreatePeopleServer) error {
	ctx := srv.Context()
	db, err := ConnDB()
	if err != nil {
		return err
	}
	var resp *emptypb.Empty
	for {
		select {
		case <-ctx.Done():
			err = srv.SendAndClose(resp)
			if err != nil {
				return err
			}
			return ctx.Err()
		default: // no-op
		}

		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		resp, err = s.createPeople(ctx, db, in.GetData())
		if err != nil {
			return err
		}
	}

}

func (s *Server) createPeople(ctx context.Context, db *sqlx.DB, in *pbfilms.PeopleData) (*emptypb.Empty, error) {
	people := ent.PeopleFromPb(in)
	err := storeql.InsertIntoDB(ctx, db, people)
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if ok && !pgErr.IntegrityViolation() { // Ignore integrity violations (to follow up same behavior as w django)
			return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
		}
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) CreateJoinPeopleFilm(srv pbfilms.Films_CreateJoinPeopleFilmServer) error {
	ctx := srv.Context()
	db, err := ConnDB()
	if err != nil {
		return err
	}
	var resp *emptypb.Empty
	for {
		select {
		case <-ctx.Done():
			err = srv.SendAndClose(resp)
			if err != nil {
				return err
			}
			return ctx.Err()
		default: // no-op
		}

		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		resp, err = s.createJoinPeopleFilm(ctx, db, in)
		if err != nil {
			return err
		}
	}
}

func (s *Server) createJoinPeopleFilm(ctx context.Context, db *sqlx.DB, in *pbfilms.CreateJoinPeopleFilmRequest) (*emptypb.Empty, error) {
	filmId, err := ent.GetFilmIdByExternalReference(ctx, db, in.GetFilmExternalReference())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "GetFilmIdByExternalReference: %v", err)
	}
	pplId, err := ent.GetPeopleIdByExternalReference(ctx, db, in.GetPeopleExternalReference())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "GetPeopleIdByExternalReference: %v", err)
	}
	join := &ent.JoinPeopleFilm{FilmId: filmId, PeopleId: pplId}

	err = join.InsertIntoDB(ctx, db)
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if ok && !pgErr.IntegrityViolation() { // Ignore integrity violations (to follow up same behavior as w django)
			return nil, status.Errorf(xerrors.Internal, "InsertIntoDB: %v", err)
		}
	}
	return &emptypb.Empty{}, nil
}
