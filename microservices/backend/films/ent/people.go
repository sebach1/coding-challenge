package ent

import (
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
)

type People struct {
	Id                uint64 `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Gender            string `json:"gender,omitempty"`
	Age               string `json:"age,omitempty"`
	HairColor         string `json:"hair_color,omitempty"`
	EyeColor          string `json:"eye_color,omitempty"`
	Slug              string `json:"slug,omitempty"`
	ExternalReference string `json:"external_reference,omitempty"`
}

func (p *People) ToPb() *pbfilms.PeopleData {
	return &pbfilms.PeopleData{
		Name:              p.Name,
		ExternalReference: p.ExternalReference,
		Slug:              p.Slug,
		HairColor:         p.HairColor,
		EyeColor:          p.EyeColor,
		Age:               p.Age,
	}
}

func GetPeopleByIds(ctx context.Context, db *sqlx.DB, peopleIds ...uint64) ([]*People, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM people WHERE id=in($1)`, peopleIds)
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
	return people, nil
}

func (p *People) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}

func SlicePeopleByIds(people []*People, ids ...uint64) (slicedPeople []*People) {
	for _, film := range people {
		for _, id := range ids {
			if id == film.Id {
				slicedPeople = append(slicedPeople, film)
			}
		}
	}
	return
}
