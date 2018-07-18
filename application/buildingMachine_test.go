package application

import (
	"testing"

	"github.com/aizu-go-kapro/web-int-searcher/di"
	"github.com/aizu-go-kapro/web-int-searcher/domain"
	mgo "gopkg.in/mgo.v2"
)

func TestBuildingMachine_Run(t *testing.T) {
	type fields struct {
		IndexRepo domain.IndexRepository
		PageRepo  domain.PageRepository
		DBSession *mgo.Session
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "hoge",
			fields: fields{
				di.InjectIndexRepository(),
				di.InjectPageRepository(),
				di.InjectDB(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bm := &BuildingMachine{
				IndexRepo: tt.fields.IndexRepo,
				PageRepo:  tt.fields.PageRepo,
				DBSession: tt.fields.DBSession,
			}
			bm.Run()
		})
	}
}
