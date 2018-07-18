package crawler

import "testing"

func TestCrawle(t *testing.T) {
	type args struct {
		url   string
		title string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample",
			args: args{
				url:   "http://web-ext.u-aizu.ac.jp/~k-asai/welcome.html",
				title: "k-asai",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Crawle(tt.args.url, tt.args.title); (err != nil) != tt.wantErr {
				t.Errorf("Crawle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
