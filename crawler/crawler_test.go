package crawler

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestGetDoc(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "hoge",
			args: args{
				url: "http://web-int.u-aizu.ac.jp/~nisim/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := GetDoc(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("GetDoc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetALLURL(t *testing.T) {
	doc, err := GetDoc("http://web-int.u-aizu.ac.jp/~nisim/")
	if err != nil {
		t.Errorf("GetDoc() error = %v", err)
	}
	type args struct {
		doc *goquery.Document
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "hoge",
			args: args{
				doc: doc,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetALLURL(tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("GetALLURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
