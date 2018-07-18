package domain

import (
	"gopkg.in/mgo.v2"
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
		Save(ms *mgo.Session, i Index) error
		Get(ms *mgo.Session, word string) (*Index, error)
		Update(ms *mgo.Session, i Index) error
	}
)

func NewIndex(word string, ids []string) *Index {
	return &Index{
		Word:    word,
		PageIDs: ids,
	}
}
