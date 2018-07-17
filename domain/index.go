package domain

import (
	"github.com/aizu-go-kapro/web-int-searcher/lib/mongoutil"
)

type (
	// IndexはWordがどこに存在するかをまとめた索引
	Index struct {
		// 検索の単位となる単語
		Word string
		// Wordが存在するページのユニークなIDのリスト
		PageIDs []string
	}

	IndexRepository interface {
		Save(ms mongoutil.Session, i Index) error
		Get(ms mongoutil.Session, i Index) (Index, error)
		Update(ms mongoutil.Session, i Index) error
	}
)

func NewIndex(word string, ids []string) *Index {
	return &Index{
		Word:    word,
		PageIDs: ids,
	}
}
