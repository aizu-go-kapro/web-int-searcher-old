package application

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	mgo "gopkg.in/mgo.v2"
)

type BuildingMachine {
	DBSession *mgo.Session
    IndexRepo domain.IndexRepository
	PageRepo domain.PageRepository
}

func NewBuildingMachine(sess *mgo.Session, indexRepo domain.IndexRepository, pageRepo domain.PageRepository) {
	return &struct{
		sess,
		indexRepo,
		pageRepo
	}
}

func (bm *BuildingMachine) Run() {
	log.Println("start:\tindex_building_machine")
	log.Println("moving:\tcrawler")
	pages, err := crawler.Crawler()
	if err != nil {
		panic(err)
	}

	log.Println("finished:\tcrawler")
	log.Println("start:\tindex_maker")
	err = MakeIndex(pages)
	if err != nil {
		panic(err)
	}
	log.Println("finished:\tindex_maker")
	log.Println("start:\tdocuments_maker")

	return nil
}

func MakeIndex(pages []page.Page) error {
	for i, page := range pages {
		log.Println(i*100/int(len(pages)), "%")
		page_words, err := ParseText(page.Text)
		if err != nil {
			return err
		}

		for _, page_word := range page_words {
			index, err := domain.GetByWord(ms, page_word)
			if err != nil {
				index.Word = page_word
				index.PageIDs = append(index.PageIDs, page.Id)
				err = domain.(index)
				if err != nil {
					return err
				}
			} else {
				index.PageIDs = append(index.PageIDs, page.Id)

				err = mongo.UpdateIdex(index)
				if err != nil {
					return err
				}
			}

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ParseText(text string) ([]string, error) {
	m, err := generateMeCab()
	if err != nil {
		return nil, err
	}
	node, err := parseToNode(text, m)
	if err != nil {
		return nil, err
	}

	words := extractSurface(node)
	return words, nil
}

func generateMeCab() (*mecab.MeCab, error) {
	m, err := mecab.New("-d /usr/local/lib/mecab/dic/mecab-ipadic-neologd")
	if err != nil {
		return nil, err
	}
	return m, nil
}

func parseToNode(text string, m *mecab.MeCab) (*mecab.Node, error) {
	tg, err := m.NewTagger()
	if err != nil {
		return nil, err
	}
	lt, err := m.NewLattice(text)
	if err != nil {
		return nil, err
	}
	node := tg.ParseToNode(lt)
	return node, nil
}

func extractSurface(node *mecab.Node) []string {
	var words []string
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" || features[0] == "動詞" {
			words = append(words, node.Surface())
		}
		if node.Next() != nil {
			break
		}
	}

	return words
}
