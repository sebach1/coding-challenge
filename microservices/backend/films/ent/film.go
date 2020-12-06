package ent

import (
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Film struct {
	Id     uint64 `json:"id,omitempty"`
Title string
Description []byte
DirectorName string
ProducerName string
Rating int
ReleaseYear int
Slug string
ExternalReference string
}


// func (p *People) ToPb() *pbusers.Identification {
// 	return &pbusers.Identification{
// 		Verified: i.Verified,
// 		UserId:   i.UserId,
// 		Dni:      uint64(i.DNI),
// 		Name:     string(i.Name),
// 		Surname:  string(i.Surname),
// 		License:  i.License,
// 		Tome:     i.Tome,
// 		Folio:    i.Folio,
// 		Cue:      i.CUE,
// 	}
// }

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


