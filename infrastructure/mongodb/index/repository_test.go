package index

import (
	"reflect"
	"testing"

	"github.com/aizu-go-kapro/web-int-searcher/domain"
	"gopkg.in/mgo.v2"
)

func TestRepository_Get(t *testing.T) {
	type args struct {
		ms   *mgo.Session
		word string
	}
	tests := []struct {
		name    string
		r       *Repository
		args    args:
		want    *domain.Index
		wantErr bool
	}{
		{
			name: "sample",
			r:    &Repository{},
			args: args{
				ms: di.InjectDB(),
				word: "test"
			},
			want:*domain.Index,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{}
			got, err := r.Get(tt.args.ms, tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
