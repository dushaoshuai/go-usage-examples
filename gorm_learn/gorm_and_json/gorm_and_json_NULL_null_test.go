package gorm_and_json

import (
	"api-examples/gorm_learn"
)

// 测试目的：
// MySQL 中的 JSON 值为 MySQL NULL, JSON null 时，取出来的结构体指针是不是 nil 呢？
//                           是         不是

func Example_gorm_and_json_NULL_null() {
	db := gorm_learn.DefaultGormDB()

	var testData jsonTest
	err := db.Model(&jsonTest{}).
		Where("pk = ?", 1).
		Find(&testData).
		Error
	if err != nil {
		panic(err)
	}

	gorm_learn.FmtStruct(testData)
	// Output:
	// struct info:
	//
	// name: jsonTest
	//
	// field info:
	//
	// name: Pk
	// value: 1
	//
	// name: JsonDefaultMysqlNULL
	// value: <nil>
	//
	// name: JsonDefaultJsonNull
	// value: {A:0 B:}
	//
	// name: JsonDefaultJsonObject
	// value: {A:0 B:}
}
