package unixtime_test

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dushaoshuai/go-usage-examples/gorm"
)

type Ints struct {
	Int8  int8  `gorm:"serializer:unixtime"`
	Int16 int16 `gorm:"serializer:unixtime"`
	Int32 int32 `gorm:"serializer:unixtime"`
	Int64 int64 `gorm:"serializer:unixtime"`
	Int   int   `gorm:"serializer:unixtime"`
}

type Uints struct {
	Uint8  uint8  `gorm:"serializer:unixtime"`
	Uint16 uint16 `gorm:"serializer:unixtime"`
	Uint32 uint32 `gorm:"serializer:unixtime"`
	Uint64 uint64 `gorm:"serializer:unixtime"`
	Uint   uint   `gorm:"serializer:unixtime"`
}

func Test_unixtime(t *testing.T) {
	db := gorm.PostgreSQLDB(context.Background())

	if err := db.Migrator().AutoMigrate(&Ints{}, &Uints{}); err != nil {
		t.Fatal(err)
	}

	err := db.WithContext(context.Background()).Create(&Ints{
		Int8:  math.MaxInt8,
		Int16: math.MaxInt16,
		Int32: math.MaxInt32,
		Int64: math.MaxInt32,
		Int:   math.MaxInt32,
	}).Error
	require.NoError(t, err)

	require.Panics(t, func() {
		db.WithContext(context.Background()).Create(&Uints{
			Uint8:  math.MaxUint8,
			Uint16: math.MaxUint16,
			Uint32: math.MaxUint32,
			Uint64: math.MaxUint32,
			Uint:   math.MaxUint32,
		})
	})

	// Output:
	// 2025/10/11 11:41:24 go-usage-examples/gorm_learn/unixtime/unixtime_test.go:32
	// [59.046ms] [rows:1] SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'ints' AND table_type = 'BASE TABLE'
	//
	// 2025/10/11 11:41:24 go-usage-examples/gorm_learn/unixtime/unixtime_test.go:32
	// [20.329ms] [rows:0] CREATE TABLE "ints" ("int8" smallint,"int16" smallint,"int32" integer,"int64" bigint,"int" bigint)
	//
	// 2025/10/11 11:41:24 go-usage-examples/gorm_learn/unixtime/unixtime_test.go:32
	// [1.062ms] [rows:1] SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'uints' AND table_type = 'BASE TABLE'
	//
	// 2025/10/11 11:41:24 go-usage-examples/gorm_learn/unixtime/unixtime_test.go:32
	// [1.559ms] [rows:0] CREATE TABLE "uints" ("uint8" smallint,"uint16" integer,"uint32" bigint,"uint64" bigint,"uint" bigint)
	//
	// 2025/10/11 11:41:24 go-usage-examples/gorm_learn/unixtime/unixtime_test.go:36 failed to encode args[0]: unable to encode &schema.serializer{Field:(*schema.Field)(0xc0004ce800), Serializer:schema.SerializerInterface(nil), SerializeValuer:schema.UnixSecondSerializer{}, Destination:reflect.Value{typ_:(*abi.Type)(0x6f129a0), ptr:(unsafe.Pointer)(0xc000266120), flag:0x199}, Context:context.backgroundCtx{emptyCtx:context.emptyCtx{}}, value:interface {}(nil), fieldValue:127} into binary format for int2 (OID 21): invalid field type 127 for UnixSecondSerializer, only int, uint supported
	// [0.814ms] [rows:0] INSERT INTO "ints" ("int8","int16","int32","int64","int") VALUES (NULL,'1970-01-01 09:06:07','2038-01-19 03:14:07','2038-01-19 03:14:07','2038-01-19 03:14:07')
	//    unixtime_test.go:43:
	//        	Error Trace:	go-usage-examples/gorm_learn/unixtime/unixtime_test.go:43
	//        	Error:      	Received unexpected error:
	//        	            	failed to encode args[0]: unable to encode &schema.serializer{Field:(*schema.Field)(0xc0004ce800), Serializer:schema.SerializerInterface(nil), SerializeValuer:schema.UnixSecondSerializer{}, Destination:reflect.Value{typ_:(*abi.Type)(0x6f129a0), ptr:(unsafe.Pointer)(0xc000266120), flag:0x199}, Context:context.backgroundCtx{emptyCtx:context.emptyCtx{}}, value:interface {}(nil), fieldValue:127} into binary format for int2 (OID 21): invalid field type 127 for UnixSecondSerializer, only int, uint supported
	//        	Test:       	Test_unixtime
	// --- FAIL: Test_unixtime (0.10s)
	//
	// FAIL
}
