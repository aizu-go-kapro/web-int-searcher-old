package page

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/yuin/charsetutil"
)

const (
	ToppagePath = "toppage.json"
)

var sawPages []string

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

func NewPage(url, title string) Page {
	id := ""
	return Page{
		id, url, title, "", nil, nil,
	}
}

func (p *Page) GetText() error {
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
	text, err := charsetutil.DecodeString(doc.Find("body").Text(), "EUC-JP")
	if err != nil {
		return "", err
	}
	result := strings.Join(
		strings.Split(
			strings.TrimSpace(
				text,
			),
			"\n",
		), " ",
	)
	return result, nil
}

func geturlfrompage(url string) ([]string, error) {
	sawPages = append(sawPages, url)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err, doc)
		return nil, err
	}
	var result []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		link, _ := s.Attr("href")

		r := regexp.MustCompile(`web-ext`)
		r2 := regexp.MustCompile(`web-int`)
		r3 := regexp.MustCompile(`u-aizu`)
		r4 := regexp.MustCompile(`html`)
		if (r.MatchString(link) == true || r2.MatchString(link) == true || r3.MatchString(link) == true) && contains(sawPages, link) != true && r4.MatchString(link) == true {
			fmt.Println("link", link)
			result = append(
				result, link,
			)
		}
	})
	return result, nil
}
func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
