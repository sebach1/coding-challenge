package ent

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type JoinPeopleFilm struct {
	Id       uint64 `json:"id,omitempty"`
	FilmId   uint64 `json:"film_id,omitempty"`
	PeopleId uint64 `json:"people_id,omitempty"`
}

func GetJoinsByFilmIds(ctx context.Context, db *sqlx.DB, filmIds ...uint64) ([]*JoinPeopleFilm, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM join_people_films WHERE film_id=in($1)`, filmIds)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()

	var joins []*JoinPeopleFilm
	for rows.Next() {
		join := &JoinPeopleFilm{}
		err = rows.StructScan(join)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		joins = append(joins, join)
	}
	return joins, nil
}

func JoinsToMap(joins ...*JoinPeopleFilm) map[uint64][]uint64 {
	joinsMap := make(map[uint64][]uint64)
	for _, join := range joins {
		peopleIds := joinsMap[join.FilmId]
		peopleIds = append(peopleIds, join.PeopleId)
		joinsMap[join.FilmId] = peopleIds
	}
	return joinsMap
}
