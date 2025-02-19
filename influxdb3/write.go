package influxdb3

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/v2/influxdb3"
	"github.com/influxdata/line-protocol/v2/lineprotocol"
)

// https://docs.influxdata.com/influxdb3/core/write-data/client-libraries/#construct-points-and-write-line-protocol
func Write() error {
	url := os.Getenv("INFLUX_HOST")
	token := os.Getenv("INFLUX_TOKEN")
	database := os.Getenv("INFLUX_DATABASE")

	// To instantiate a client, call New() with InfluxDB credentials.
	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: database,
	})
	if err != nil {
		panic(err)
	}

	/** Use a deferred function to ensure the client is closed when the
	 * function returns.
	**/
	defer func(client *influxdb3.Client) {
		err = client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	/** Use the NewPoint method to construct a point.
	 * NewPoint(measurement, tags map, fields map, time)
	**/
	point := influxdb3.NewPoint("home",
		map[string]string{
			"room": "Living Room",
		},
		map[string]any{
			"temp": 24.5,
			"hum":  40.5,
			"co":   15i},
		time.Now(),
	)

	/** Use the NewPointWithMeasurement method to construct a point with
	 * method chaining.
	**/
	point2 := influxdb3.NewPointWithMeasurement("home").
		SetTag("room", "Living Room").
		SetField("temp", 23.5).
		SetField("hum", 38.0).
		SetField("co", 16i).
		SetTimestamp(time.Now())

	fmt.Println("Writing points")
	points := []*influxdb3.Point{point, point2}

	/** Write points to InfluxDB.
	 * You can specify WriteOptions, such as Gzip threshold,
	 * default tags, and timestamp precision. Default precision is lineprotocol.Nanosecond
	**/
	err = client.WritePoints(context.Background(), points,
		influxdb3.WithPrecision(lineprotocol.Nanosecond))
	return nil
}
