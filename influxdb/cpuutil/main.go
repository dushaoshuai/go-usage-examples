package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
	raw, err := collect()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	statistics, err := parse(raw)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	cli := getInfluxdbCli()
	ctx := context.Background()
	measurement := "cpu_utilization"

	for _, stat := range statistics {
		for _, load := range stat.CpuLoad {
			t, err := time.ParseInLocation(time.DateTime, stat.Timestamp, time.Local)
			if err != nil {
				slog.Error(err.Error())
				return
			}

			point := write.NewPointWithMeasurement(measurement).
				AddTag("cpu", load.Cpu).
				AddTag("core", load.Core).
				AddTag("socket", load.Socket).
				AddTag("node", load.Node).
				AddField("%usr", load.Usr).
				AddField("%nice", load.Nice).
				AddField("%sys", load.Sys).
				AddField("%iowait", load.Iowait).
				AddField("%irq", load.Irq).
				AddField("%soft", load.Soft).
				AddField("%steal", load.Steal).
				AddField("%guest", load.Guest).
				AddField("%gnice", load.Gnice).
				AddField("%idle", load.Idle).
				SetTime(t) // TODO: time is incorrect

			err = cli.WritePoint(ctx, point)
			if err != nil {
				slog.Error(err.Error())
				return
			}

			slog.Info("write succeed...")
		}
	}
}
