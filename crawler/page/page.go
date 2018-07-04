package page

import (
	"encoding/json"
	"io/ioutil"
)

const (
	ToppagePath = "toppage.json"
)

type TopPages struct {
	Url   string `json: "url"`
	Title string `json: "title"`
}

type Pages struct {
	Id       string   `json: "id"`
	Url      string   `json: "url"`
	Title    string   `json: "title"`
	Text     string   `json:	"text"`
	Tolink   []string `json: "tolink"`
	ToBelink []string `json: "tobelink"`
}

func LoadTopPage() ([]TopPages, error) {
	p := []TopPages{}
	bytes, err := ioutil.ReadFile("./" + ToppagePath)
	if err != nil {
		return p, err
	}
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}


