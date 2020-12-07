package ent

import "database/sql/driver"

func (p *People) GetId() uint64 {
	return p.Id
}

func (p *People) SetId(id uint64) {
	p.Id = id
}

func (p *People) SQLTable() string {
	return "people"
}

func (p *People) SQLMap() map[string]driver.Value {
	return map[string]driver.Value{
		"id":                 p.Id,
		"name":               p.Name,
		"gender":             p.Gender,
		"age":                p.Age,
		"hair_color":         p.HairColor,
		"eye_color":          p.EyeColor,
		"slug":               p.Slug,
		"external_reference": p.ExternalReference,
	}
}
