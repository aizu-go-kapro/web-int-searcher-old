package indexmaker

import (
	"reflect"
	"testing"
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
	// _, err := ParseText("noahは異常者です。なぜなら異常行動をするからです。")
	// if err != nil {
	// 	panic(err)
	// }
}
