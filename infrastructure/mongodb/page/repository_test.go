package page

import (
	"reflect"
	"testing"

	"github.com/aizu-go-kapro/web-int-searcher/domain"
)

func TestRepository_GetFromCrawler(t *testing.T) {
	tests := []struct {
		name    string
		r       *Repository
		want    []*domain.Page
		wantErr bool
	}{
		{
			name:    "sampke",
			r:       &Repository{},
			want:    []*domain.Page{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{}
			got, err := r.GetFromCrawler()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetFromCrawler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetFromCrawler() = %v, want %v", got, tt.want)
			}
		})
	}
}
