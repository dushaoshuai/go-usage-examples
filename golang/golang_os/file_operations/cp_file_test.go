package file_operations

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCpFile(t *testing.T) {
	temp, err := os.MkdirTemp("", "*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(temp)

	type args struct {
		dst string
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "file",
			args: args{
				dst: filepath.Join(temp, "cp_file_test.go"),
				src: "./cp_file_test.go",
			},
			wantErr: false,
		},
		{
			name: "directory",
			args: args{
				dst: filepath.Join(temp, "dstfile"),
				src: ".",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CpFile(tt.args.dst, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("CpFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
