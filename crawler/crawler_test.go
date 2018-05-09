package crawler

import (
	"reflect"
	"testing"
	"time"
)

func TestGetALLURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []Link
		wantErr bool
	}{
		{
			name: "example.com",
			args: args{
				url: "http://example.com/",
			},
			want: []Link{
				Link{Url: "http://www.iana.org/domains/example"},
			},
			wantErr: false,
		},
		{
			name: "get_example",
			args: args{
				url: "http://web-int.u-aizu.ac.jp/~nisim/",
			},
			want: []Link{
				Link{Url: "logic_design/index.html"},
				Link{Url: "comparch/index.html"},
				Link{Url: "cg_gpu/index.html"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetALLURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetALLURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetALLURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchDataFromURL(t *testing.T) {
	type args struct {
		link Link
	}
	tests := []struct {
		name    string
		args    args
		want    Link
		wantErr bool
	}{
		{
			name: "sample",
			args: args{
				link: Link{
					Url: "http://example.com",
				},
			},
			want: Link{
				Url:   "http://example.com",
				Title: "Example Domain",
				Date:  time.Now().Unix(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchDataFromURL(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchDataFromURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchDataFromURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllText(t *testing.T) {
	type args struct {
		link Link
	}
	tests := []struct {
		name    string
		args    args
		want    Link
		wantErr bool
	}{
		{
			name: "sample",
			args: args{
				link: Link{
					Url: "http://example.com",
				},
			},
			want: Link{
				Url:  "http://example.com",
				Text: "Example Domain     This domain is established to be used for illustrative examples in documents. You may use this     domain in examples without prior coordination or asking for permission.     More information...",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllText(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllText() = %v, want %v", got, tt.want)
			}
		})
	}
}
