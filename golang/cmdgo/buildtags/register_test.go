package buildtags

import (
	"reflect"
	"slices"
	"testing"

	"github.com/dushaoshuai/go-usage-examples/golang/cmdgo/buildtags/registry"
)

//go:generate go test -tags a,b,c -test.run ^TestAllPkgs$
func TestAllPkgs(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "test",
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := registry.AllPkgs()
			slices.Sort(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPkgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:generate go test -tags a -test.run ^TestAllPkgs_a$
func TestAllPkgs_a(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "test",
			want: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := registry.AllPkgs()
			slices.Sort(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPkgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:generate go test -tags b -test.run ^TestAllPkgs_b$
func TestAllPkgs_b(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "test",
			want: []string{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := registry.AllPkgs()
			slices.Sort(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPkgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:generate go test -tags c -test.run ^TestAllPkgs_c$
func TestAllPkgs_c(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "test",
			want: []string{"c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := registry.AllPkgs()
			slices.Sort(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPkgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go:generate go test -tags a,c -test.run ^TestAllPkgs_ac$
func TestAllPkgs_ac(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "test",
			want: []string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := registry.AllPkgs()
			slices.Sort(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPkgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
