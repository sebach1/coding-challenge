package ent

import (
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
)

type Film struct {
	Id                uint64 `json:"id,omitempty"`
	Title             string `json:"title,omitempty"`
	Description       []byte `json:"description,omitempty"`
	DirectorName      string `json:"director_name,omitempty"`
	ProducerName      string `json:"producer_name,omitempty"`
	Rating            int    `json:"rating,omitempty"`
	ReleaseYear       int    `json:"release_year,omitempty"`
	Slug              string `json:"slug,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
}

func (f *Film) ToPb() *pbfilms.FilmData {
	return &pbfilms.FilmData{
		Title:             f.Title,
		ExternalReference: f.ExternalReference,
		Slug:              f.Slug,
		DirectorName:      f.DirectorName,
		ProducerName:      f.ProducerName,
		Rating:            uint64(f.Rating),
		ReleaseYear:       uint64(f.ReleaseYear),
		Description:       string(f.Description),
	}
}

func GetFilmsWithPeople(ctx context.Context, db *sqlx.DB, limit uint64) (map[*Film][]*People, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM films LIMIT $1`, limit)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()

	var films []*Film
	var filmIds []uint64
	for rows.Next() {
		film := &Film{}
		err = rows.StructScan(film)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		films = append(films, film)
		filmIds = append(filmIds, film.Id)
	}

	joins, err := GetJoinsByFilmIds(ctx, db, filmIds...)
	if err != nil {
		return nil, errors.Wrap(err, "GetJoinsByFilmIds")
	}

	var peopleIds []uint64
	for _, join := range joins {
		peopleIds = append(peopleIds, join.PeopleId)
	}

	rows, err = db.QueryxContext(ctx, `SELECT * FROM people WHERE id in($1)`, peopleIds)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()

	var people []*People
	for rows.Next() {
		ppl := &People{}
		err = rows.StructScan(ppl)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		people = append(people, ppl)
	}

	joinsMap := JoinsToMap(joins...)
	result := make(map[*Film][]*People)
	for _, film := range films {
		result[film] = SlicePeopleByIds(people, joinsMap[film.Id]...)
	}

	return result, nil
}

func (f *Film) String() string {
	s, _ := json.Marshal(f)
	return string(s)
}

func FindFilm(ctx context.Context, db *sqlx.DB, uid uint64) (*Film, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM films WHERE id=$1`, uid)
	f := &Film{}
	err := row.StructScan(f)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return f, nil
}
