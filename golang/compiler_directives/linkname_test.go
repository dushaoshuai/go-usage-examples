package compiler_directives_test

import (
	_ "database/sql" // import link package
	"runtime"
	_ "runtime" // import link package
	"testing"
	"time"
	_ "unsafe" // required by //go:linkname
)

//go:linkname convertAssign database/sql.convertAssign
func convertAssign(dest, src any) error

// go test -ldflags=-checklinkname=0 -run Test_linkname_func
func Test_linkname_func(t *testing.T) {
	var dest string
	src := time.Date(2024, 10, 1, 15, 11, 30, 3333, time.Local)
	want := "2024-10-01T15:11:30.000003333+08:00"

	err := convertAssign(&dest, src)
	if err != nil {
		t.Fatalf("convertAssign() err = %v", err)
	}
	if dest != want {
		t.Fatalf("convertAssign(), got = %v, want %v", dest, want)
	}
}

var (
	//go:linkname buildVersion runtime.buildVersion
	buildVersion string
)

// go test -ldflags=-checklinkname=0 -run Test_linkname_var
func Test_linkname_var(t *testing.T) {
	wantBuildVersion := runtime.Version()

	if buildVersion != wantBuildVersion {
		t.Errorf("buildVersion = %s, want %s", buildVersion, wantBuildVersion)
	}
}
