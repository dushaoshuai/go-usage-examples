package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)

	org := "initorg"
	bucket := "initbucket"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for i := range 5 {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]any{
			"field1": i,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(time.Second)

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}

	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "initbucket")
				|> range(start: -10m)
				|> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err = results.Err(); err != nil {
		log.Fatal(err)
	}

	query = `from(bucket: "initbucket")
              |> range(start: -10m)
              |> filter(fn: (r) => r._measurement == "measurement1")
              |> mean()`
	results, err = queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err = results.Err(); err != nil {
		log.Fatal(err)
	}
}
