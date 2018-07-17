package application

import "testing"

func Test_indexMakerApp(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "sample",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := indexMakerApp(); (err != nil) != tt.wantErr {
				t.Errorf("indexMakerApp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
