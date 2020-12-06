package filmstest

import (
	"testing"

	"github.com/sebach1/coding-challenge/microservices/backend/people/ent"
)

func CopyFilm(t *testing.T, o *ent.Film) *ent.Film {
	t.Helper()
	if o == nil {
		t.Fatal("cant copy nil ent")
	}
	cp := ent.Film{}
	cp = *o
	return &cp
}

func CopyPeople(t *testing.T, o *ent.People) *ent.People {
	t.Helper()
	if o == nil {
		t.Fatal("cant copy nil ent")
	}
	cp := ent.People{}
	cp = *o
	return &cp
}
