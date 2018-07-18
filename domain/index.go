package domain

import mgo "gopkg.in/mgo.v2"

type (
	// IndexはWordがどこに存在するかをまとめた索引
	Index struct {
		// 検索の単位となる単語
		Word string
		// Wordが存在するページのユニークなIDのリスト
		PageIDs []string
	}

	IndexRepository interface {
		Save(session *mgo.Session, i Index) error
		Get(session *mgo.Session, i Index) (Index, error)
		GestByWord(session *mgo.Session, word string) (Index, error)
		Update(session *mgo.Session, i Index) error
	}
)

func NewIndex(word string, ids []string) *Index {
	return &Index{
		Word:    word,
		PageIDs: ids,
	}
}
