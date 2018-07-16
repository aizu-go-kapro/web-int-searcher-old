package mongo

import (
	"reflect"
	"testing"
)

func TestSaveIndex(t *testing.T) {
	type args struct {
		index Index
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "insert sample",
			args: args{
				Index{"Test", []string{"UUID"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveIndex(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("SaveIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetIndex(t *testing.T) {
	TestSaveIndex(t)
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		want    Index
		wantErr bool
	}{
		{
			name: "get index test",
			args: args{
				"Test",
			},
			want: Index{
				"Test", []string{"UUID"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIndex(tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
