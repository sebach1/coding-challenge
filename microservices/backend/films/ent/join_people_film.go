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
	if len(filmIds) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(`SELECT * FROM join_people_films WHERE film_id IN(?);`, filmIds)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.In")
	}
	query = db.Rebind(query)
	rows, err := db.QueryxContext(ctx, query, args...)
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

func (j *JoinPeopleFilm) InsertIntoDB(ctx context.Context, db *sqlx.DB) error {
	ids, err := db.NamedQueryContext(ctx, `INSERT INTO join_people_films(film_id, people_id) VALUES(:film_id, :people_id) RETURNING id`, j)
	if err != nil {
		return errors.Wrap(err, "NamedQueryContext")
	}
	defer ids.Close()
	var i int
	for ids.Next() {
		var id uint64
		err := ids.Scan(&id)
		if err != nil {
			return errors.Wrap(err, "Scan")
		}
		j.SetId(id)
		i += 1
	}
	err = ids.Err()
	if err != nil {
		return errors.Wrap(err, "cursor.Err")
	}
	return nil
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
