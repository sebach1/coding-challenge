package wrk

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/sebach1/coding-challenge/microservices/backend/ghibli/sdk"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
)

type Job struct {
	Name   string
	RampUp time.Duration
	Task   Task
}

type Task func(ctx context.Context) error

func SyncGhibliFilms(ctx context.Context) error {
	filmsCli, filmsCloser, err := ConnFilms(ctx)
	if err != nil {
		return errors.Wrap(err, "ConnFilms")
	}
	defer filmsCloser()
	filmsCreateCli, err := filmsCli.CreateFilms(ctx)
	if err != nil {
		return errors.Wrap(err, "CreateFilms")
	}
	return syncGhibliFilms(ctx, filmsCreateCli)
}

func syncGhibliFilms(ctx context.Context, filmsCli pbfilms.Films_CreateFilmsClient) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	ghibliCli := &sdk.Ghibli{Client: http.Client{Transport: tr}}
	films, err := ghibliCli.GetFilms(ctx, 250)
	if err != nil {
		return errors.Wrap(err, "GetFilms")
	}
	for _, film := range films {
		pbFilm, err := film.ToPb()
		if err != nil {
			return errors.Wrap(err, "ToPb")
		}
		err = filmsCli.Send(&pbfilms.CreateFilmRequest{Data: pbFilm})
		if err != nil {
			return errors.Wrap(err, "Send")
		}
	}
	return SyncGhibliPeople(ctx)
}

func SyncGhibliPeople(ctx context.Context) error {
	filmsCli, filmsCloser, err := ConnFilms(ctx)
	if err != nil {
		return errors.Wrap(err, "ConnFilms")
	}
	defer filmsCloser()
	return syncGhibliPeople(ctx, filmsCli)
}

func syncGhibliPeople(ctx context.Context, filmsCli pbfilms.FilmsClient) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	ghibliCli := &sdk.Ghibli{Client: http.Client{Transport: tr}}
	ppl, err := ghibliCli.GetPeople(ctx, 250)
	if err != nil {
		return errors.Wrap(err, "GetPeople")
	}

	peopleCreateClient, err := filmsCli.CreatePeople(ctx)
	if err != nil {
		return errors.Wrap(err, "CreatePeople")
	}
	err = syncGhibliPeopleCreation(ctx, peopleCreateClient, ppl)
	if err != nil {
		return errors.Wrap(err, "syncGhibliPeopleCreation")
	}

	joinsCreateClient, err := filmsCli.CreateJoinPeopleFilm(ctx)
	if err != nil {
		return errors.Wrap(err, "CreateJoinPeopleFilm")
	}
	err = syncGhibliPeopleJoins(ctx, joinsCreateClient, ppl)
	if err != nil {
		return errors.Wrap(err, "syncGhibliPeopleJoins")
	}
	return nil
}

func syncGhibliPeopleCreation(ctx context.Context, filmsCli pbfilms.Films_CreatePeopleClient, ppl []*sdk.People) error {
	for _, p := range ppl {
		err := filmsCli.Send(&pbfilms.CreatePeopleRequest{Data: p.ToPb()})
		if err != nil {
			return errors.Wrap(err, "Send")
		}
	}
	return nil
}

func syncGhibliPeopleJoins(ctx context.Context, filmsCli pbfilms.Films_CreateJoinPeopleFilmClient, ppl []*sdk.People) error {
	for _, ppl := range ppl {
		for _, filmId := range ppl.FilmIds() {
			err := filmsCli.Send(&pbfilms.CreateJoinPeopleFilmRequest{PeopleExternalReference: ppl.Id, FilmExternalReference: filmId})
			if err != nil {
				return errors.Wrap(err, "Send")
			}
		}
	}
	return nil
}
