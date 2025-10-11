package not_a_string_type_test

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/gorm"

	"github.com/dushaoshuai/go-usage-examples/gorm_learn"
)

type notString string

func (ns notString) String() string {
	return string(ns)
}

const (
	col1 notString = "col1"
)

func Test_not_a_string_type(t *testing.T) {
	db := gorm_learn.SQLiteInMemoryDB(context.Background()).
		Session(&gorm.Session{DryRun: true})

	var val map[string]any

	notStrCond := col1 + " IS NOT NULL"
	err := db.Table("some_table").
		Where(notStrCond).
		Find(&val).
		Error
	if err != nil {
		t.Error(err)
	}
	// [0.030ms] [rows:0] SELECT * FROM `some_table` WHERE `some_table`. = "col1 IS NOT NULL"
	//    not_a_string_type_test.go:35: model value required

	strCond := col1.String() + " IS NOT NULL"
	err = db.Table("some_table").
		Where(strCond).
		Find(&val).
		Error
	if err != nil {
		t.Fatal(err)
	}
	// [0.006ms] [rows:0] SELECT * FROM `some_table` WHERE col1 IS NOT NULL

	fmt.Println()
	// "col1 IS NOT NULL": col1 IS NOT NULL: not_a_string_type_test.notString
	fmt.Printf("%q: %v: %T\n", notStrCond, notStrCond, notStrCond)
	// "col1 IS NOT NULL": col1 IS NOT NULL: string
	fmt.Printf("%q: %v: %T\n", strCond, strCond, strCond)

	fmt.Println()
	var a any = notStrCond
	_, ok := a.(string)
	fmt.Printf("notString is string?: %v\n", ok) // notString is string?: false
}
