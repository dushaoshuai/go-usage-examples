package main

import (
	"os"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func getInfluxdbCli() api.WriteAPIBlocking {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)

	org := "initorg"
	bucket := "initbucket"
	return client.WriteAPIBlocking(org, bucket)
}
