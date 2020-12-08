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

func PeopleFromPb(pbpeople *pbfilms.PeopleData) *People {
	return &People{
		Name:              pbpeople.Name,
		ExternalReference: pbpeople.ExternalReference,
		Slug:              newSlug(pbpeople.Name),
		HairColor:         pbpeople.HairColor,
		EyeColor:          pbpeople.EyeColor,
		Age:               pbpeople.Age,
	}
}

func (p *People) ToPb() *pbfilms.PeopleData {
	return &pbfilms.PeopleData{
		Name:              p.Name,
		ExternalReference: p.ExternalReference,
		Slug:              p.Slug,
		HairColor:         p.HairColor,
		EyeColor:          p.EyeColor,
		Age:               p.Age,
		Id:                uint64(p.Id),
	}
}

func GetPeopleIdByExternalReference(ctx context.Context, db *sqlx.DB, ref string) (uint64, error) {
	row := db.QueryRowxContext(ctx, `SELECT id FROM people WHERE external_reference = $1`, ref)
	var id uint64
	err := row.Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "Scan")
	}
	return id, nil
}

func GetPeopleByIds(ctx context.Context, db *sqlx.DB, peopleIds ...uint64) ([]*People, error) {
	if len(peopleIds) == 0 {
		return nil, nil
	}
	query, args, err := sqlx.In(`SELECT * FROM people WHERE id IN(?)`, peopleIds)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx.In")
	}
	query = db.Rebind(query)
	rows, err := db.QueryxContext(ctx, query, args...)
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
