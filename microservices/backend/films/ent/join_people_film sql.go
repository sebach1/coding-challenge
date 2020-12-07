package ent

import "database/sql/driver"

func (j *JoinPeopleFilm) GetId() uint64 {
	return j.Id
}

func (j *JoinPeopleFilm) SetId(id uint64) {
	j.Id = id
}

func (j *JoinPeopleFilm) SQLTable() string {
	return "people"
}

func (j *JoinPeopleFilm) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":        j.Id,
		"film_id":   j.FilmId,
		"people_id": j.PeopleId,
	}
}
