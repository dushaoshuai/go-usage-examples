package file_operations

import (
	"testing"
)

func TestIsDir(t *testing.T) {
	tests := []struct {
		name    string
		dirName string
		want    bool
		wantErr bool
	}{
		{"file", "./link_test.go", false, false},
		{"dir", ".", true, false},
		{"non-exist", "/tmp/non-exist", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDir(tt.dirName)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}
