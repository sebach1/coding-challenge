package ent

import (
	"encoding/json"
)

type People struct {
	Id                uint64 `json:"id,omitempty"`
	Name              string
	Gender            string
	Age               string
	HairColor         string
	EyeColor          string
	Slug              string
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

func (p *People) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}
