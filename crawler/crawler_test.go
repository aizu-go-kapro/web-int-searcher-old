package crawler

import (
	"reflect"
	"testing"
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
