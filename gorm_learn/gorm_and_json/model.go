package gorm_and_json

import (
	"database/sql/driver"
	"encoding/json"
)

type jsonTest struct {
	Pk                    uint64                 `gorm:"column:pk;primaryKey"`
	JsonDefaultMysqlNULL  *jsonDefaultMysqlNULL  `gorm:"column:json_default_mysql_NULL"`
	JsonDefaultJsonNull   *jsonDefaultJsonNull   `gorm:"column:json_default_json_null"`
	JsonDefaultJsonObject *jsonDefaultJsonObject `gorm:"column:json_default_json_object"`
}

func (jsonTest) TableName() string {
	return "json_test"
}

type jsonDefaultMysqlNULL struct {
	A int    `json:"a"`
	B string `json:"b"`
}

func (j *jsonDefaultMysqlNULL) Scan(src any) error {
	return json.Unmarshal(src.([]byte), j)
}

func (j *jsonDefaultMysqlNULL) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type jsonDefaultJsonNull struct {
	A int    `json:"a"`
	B string `json:"b"`
}

func (j *jsonDefaultJsonNull) Scan(src any) error {
	return json.Unmarshal(src.([]byte), j)
}

func (j *jsonDefaultJsonNull) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type jsonDefaultJsonObject struct {
	A int    `json:"a"`
	B string `json:"b"`
}

func (j *jsonDefaultJsonObject) Scan(src any) error {
	return json.Unmarshal(src.([]byte), j)
}

func (j *jsonDefaultJsonObject) Value() (driver.Value, error) {
	return json.Marshal(j)
}
