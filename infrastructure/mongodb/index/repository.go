package index

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

// TODO: domainのindexに定義してあるIndexRepositoryインターフェースを実装する。(domain/index.go#16~)
// イメージこのGetみたいな感じ。引数で*mgo.Session渡されてるのが抽象化されてなくて微妙だけど、一旦これで
func (r *Repository) Get(ms *mgo.Session, word string) (*domain.Index, error) {
	const errtag = "Repository.Get failed"
	var index domain.Index
	err := ms.DB("test").C("index").Find(bson.M{"word": word}).One(&index)
	if err != nil {
		return &index, err
	}
	return &index, nil
}

func (r *Repository) Save(ms *mgo.Session, i domain.Index) error {
	c := ms.DB("test").C("index")
	err := c.Insert(&i)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ms *mgo.Session, i domain.Index) error {
	c := ms.DB("test").C("index")
	selecter := bson.M{"word": i.Word}
	err := c.Update(selecter, i)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByWord(ms *mgo.Session, word string) (domain.Index, error) {
	var index domain.Index
	err := ms.DB("test").C("index").Find(bson.M{"word": word}).One(&index)
	if err != nil {
		return index, err
	}
	return index, nil
}
