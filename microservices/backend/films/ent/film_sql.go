package ent

import "database/sql/driver"

func (f *Film) GetId() uint64 {
	return f.Id
}

func (f *Film) SetId(id uint64) {
	f.Id = id
}

func (f *Film) SQLTable() string {
	return "films"
}

func (f *Film) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":                 f.Id,
		"title":              f.Title,
		"description":        f.Description,
		"director_name":      f.DirectorName,
		"producer_name":      f.ProducerName,
		"rating":             f.Rating,
		"release_year":       f.ReleaseYear,
		"slug":               f.Slug,
		"external_reference": f.ExternalReference,
	}
}
