package indexmaker

import (
	"reflect"
	"testing"

	"github.com/aizu-go-kapro/web-int-searcher/crawler/page"
)

func TestParseText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "noahは異常者",
			args: args{
				text: "noahは異常者です。なぜなら異常行動をするからです。",
			},
			want: []string{
				"noah", "異常者", "異常", "行動", "する",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseText(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeIndex(t *testing.T) {
	type args struct {
		pages []page.Page
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample",
			args: args{
				[]page.Page{
					{
						"hoge",
						"example.html",
						"noah",
						"noahは異常者",
						[]string{},
						[]string{},
					},
					{
						"fuga",
						"hato.html",
						"hato",
						"noahは異常者です。なぜなら異常行動をするからです。",
						[]string{},
						[]string{},
					},
					{
						"piyo",
						"hato.html",
						"hato",
						"はとバスと八坂さんのドキドキデート",
						[]string{},
						[]string{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeIndex(tt.args.pages); (err != nil) != tt.wantErr {
				t.Errorf("MakeIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
