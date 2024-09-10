package main

import (
	"testing"
)

func Test_collect(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collect()
			if (err != nil) != tt.wantErr {
				t.Errorf("collect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			t.Logf("%s\n", got)
		})
	}
}
