package index

import "github.com/aizu-go-kapro/web-int-searcher/domain"

//  | RDBでの呼称  | MongoDBでの呼称
//  | database    | database
//  | table 	  | collection
//  | row 	      | document
//  | column 	  | field
//  | index 	  | index
//  | primary key | _id field

// IndexCollectionはMongoDB用のIndexのためのDTS
type IndexCollection struct {
	Word    string   `json:"word"`
	PageIDs []string `json:"pageIDs"`
}

// NewIndexCollectionはdomain.IndexをNewIndexCollectionに変換します。
func NewIndexCollection(i *domain.Index) (*IndexCollection, error) {
	const errtag = " mongodb/index/NewIndexCollection failed"

	return &IndexCollection{
		Word:    i.Word,
		PageIDs: i.PageIDs,
	}, nil
}

// Domainはdomain.Indexに変換します。
func (i *IndexCollection) Domain() *domain.Index {
	return &domain.Index{
		Word:    i.Word,
		PageIDs: i.PageIDs,
	}
}
