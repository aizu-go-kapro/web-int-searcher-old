package index

import (
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	"gopkg.in/mgo.v2"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

// TODO: domainのindexに定義してあるIndexRepositoryインターフェースを実装する。(domain/index.go#16~)
// イメージこのGetみたいな感じ。引数でmongoutil.Session渡されてるのが抽象化されてなくて微妙だけど、一旦これで
func (r *Repository) Get(ms *mgo.Session, word string) (domain.Index, error) {
	const errtag = "Repository.Get failed"

	var index domain.Index
	return index, nil
}
