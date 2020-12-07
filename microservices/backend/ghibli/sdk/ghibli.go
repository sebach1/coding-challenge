package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sebach1/coding-challenge/microservices/pb/pbfilms"
)

type Ghibli struct {
	http.Client
}

func (gh *Ghibli) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return gh.Client.Do(req)
}

type Film struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	Producer    string `json:"producer"`
	ReleaseDate string `json:"release_date"`
	RtScore     string `json:"rt_score"`
}

func (f *Film) ToPb() (*pbfilms.FilmData, error) {
	intrt, err := f.IntRtScore()
	if err != nil {
		return nil, errors.Wrap(err, "IntRtScore")
	}
	intrd, err := f.IntReleaseDate()
	if err != nil {
		return nil, errors.Wrap(err, "IntReleaseDate")
	}
	return &pbfilms.FilmData{
		ExternalReference: f.Id,
		Title:             f.Title,
		Description:       f.Description,
		DirectorName:      f.Director,
		ProducerName:      f.Producer,
		ReleaseYear:       uint64(intrd),
		Rating:            uint64(intrt),
	}, nil
}

func (f *Film) IntRtScore() (int, error) {
	intrt, err := strconv.Atoi(f.RtScore)
	if err != nil {
		return 0, errors.Wrap(err, "strconv.Atoi")
	}
	return intrt, nil
}

func (f *Film) IntReleaseDate() (int, error) {
	intrd, err := strconv.Atoi(f.ReleaseDate)
	if err != nil {
		return 0, errors.Wrap(err, "strconv.Atoi")
	}
	return intrd, nil
}

type People struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Gender    string   `json:"gender"`
	Age       string   `json:"age"`
	EyeColor  string   `json:"eye_color"`
	HairColor string   `json:"hair_color"`
	FilmURIs  []string `json:"films"`
}

func (p *People) ToPb() *pbfilms.PeopleData {
	return &pbfilms.PeopleData{
		ExternalReference: p.Id,
		Name:              p.Name,
		Age:               p.Age,
		Gender:            p.Gender,
		HairColor:         p.HairColor,
		EyeColor:          p.EyeColor,
	}
}

func (p *People) FilmIds() (ids []string) {
	for _, uri := range p.FilmURIs {
		ids = append(ids, filepath.Base(uri))
	}
	return
}

func routeTo(resource string, params url.Values) (string, error) {
	base := "https://ghibliapi.herokuapp.com"
	base += resource
	URL, err := url.Parse(base)
	if err != nil {
		return "", errors.Wrap(err, "url.Parse")
	}
	URL.RawQuery = params.Encode()
	base = URL.String()

	return base, nil
}

func (gh *Ghibli) GetFilms(ctx context.Context, limit uint64) ([]*Film, error) {
	params := url.Values{"limit": []string{strconv.Itoa(int(limit))}}
	route, err := routeTo("/films", params)
	if err != nil {
		return nil, errors.Wrap(err, "routeTo")
	}
	resp, err := gh.Get(route)
	if err != nil {
		return nil, errors.Wrap(err, "gh.Get")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("invalid request: %v", resp.StatusCode)
	}
	var films []*Film
	err = json.NewDecoder(resp.Body).Decode(&films)
	if err != nil {
		return nil, errors.Wrap(err, "json.Decode")
	}
	return films, nil
}

func (gh *Ghibli) GetPeople(ctx context.Context, limit int) ([]*People, error) {
	params := url.Values{"limit": []string{strconv.Itoa(limit)}}
	route, err := routeTo("/people", params)
	if err != nil {
		return nil, errors.Wrap(err, "routeTo")
	}
	resp, err := gh.Get(route)
	if err != nil {
		return nil, errors.Wrap(err, "gh.Get")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("invalid request: %v", resp.StatusCode)
	}
	var people []*People
	err = json.NewDecoder(resp.Body).Decode(&people)
	if err != nil {
		return nil, errors.Wrap(err, "json.Decode")
	}
	return people, nil
}
