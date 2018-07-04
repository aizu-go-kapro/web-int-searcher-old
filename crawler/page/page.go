package page

import (
	"encoding/json"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
)

const (
	ToppagePath = "toppage.json"
)

type TopPages struct {
	Url   string `json: "url"`
	Title string `json: "title"`
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

type Page struct {
	Id       string   `json: "id"`
	Url      string   `json: "url"`
	Title    string   `json: "title"`
	Text     string   `json:	"text"`
	Tolink   []string `json: "tolink"`
	ToBelink []string `json: "tobelink"`
}

func NewPage(url , title string) Page {
	id := ""
	return Page{
		id, url, title, "", nil, nil,
	}
}

func (p *Page) GetText() error{
	text, err := gettext(p.Url)
	if err != nil {
		return err
	}
	p.Text = text
	return nil
}

func (p *Page) GetLink() error {
	links, err := geturlfrompage(p.Url)
	if err != nil {
		return err
	}
	p.Tolink = links
	return nil
}

func gettext(url string) (string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}
	result := strings.Join(
		strings.Split(
			strings.TrimSpace(
				doc.Find("body").Text(),
			),
			"\n",
		), " ",
	)
	return result, nil
}

func geturlfrompage(url string) ([]string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err, doc)
		return nil, err
	}
	var result []string{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		result = append(
			result, link,
		)
	})
	return result, nil
}