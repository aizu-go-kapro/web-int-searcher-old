package index

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	"github.com/aizu-go-kapro/web-int-searcher/lib/mongoutil"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

// TODO: domainのindexに定義してあるIndexRepositoryインターフェースを実装する。(domain/index.go#16~)
// イメージこのGetみたいな感じ。引数でmongoutil.Session渡されてるのが抽象化されてなくて微妙だけど、一旦これで
func (r *Repository) Get(ms mongoutil.Session, i domain.Index) (Index, error) {
	const errtag = "Repository.Get failed"
	var index domain.Index
	err := ms.C("index").Find(bson.M{"word": word}).One(&index)
	if err != nil {
		return index, err
	}
	return index, nil
}

func (r *Repository) Save(ms mongoutil.Session, i domain.Index) error {
	c := ms.C("index")
	err := c.Insert(&i)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ms mongoutil.Session, i domain.Index) error {
	c := ms.C("index")
	selecter := bson.M{"word": i.word}
	err := c.Update(selector, index)
	if err != nil {
		return err
	}
	return nil
}
