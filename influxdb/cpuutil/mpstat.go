package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type cpuStatistic struct {
	Timestamp string `json:"timestamp"`
	CpuLoad   []struct {
		Cpu    string  `json:"cpu"`
		Core   string  `json:"core"`
		Socket string  `json:"socket"`
		Node   string  `json:"node"`
		Usr    float64 `json:"usr"`
		Nice   float64 `json:"nice"`
		Sys    float64 `json:"sys"`
		Iowait float64 `json:"iowait"`
		Irq    float64 `json:"irq"`
		Soft   float64 `json:"soft"`
		Steal  float64 `json:"steal"`
		Guest  float64 `json:"guest"`
		Gnice  float64 `json:"gnice"`
		Idle   float64 `json:"idle"`
	} `json:"cpu-load"`
}

// mpstat -P ALL -T -o JSON
type mpstat struct {
	SysStat struct {
		Hosts []struct {
			NodeName     string         `json:"nodename"`
			SysName      string         `json:"sysname"`
			Release      string         `json:"release"`
			Machine      string         `json:"machine"`
			NumberOfCpus int            `json:"number-of-cpus"`
			Date         string         `json:"date"`
			Statistics   []cpuStatistic `json:"statistics"`
		} `json:"hosts"`
	} `json:"sysstat"`
}

func collect() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// mpstat -P ALL -o JSON -T 1 60
	cmd := exec.CommandContext(ctx, "mpstat", "-P", "ALL", "-o", "JSON", "-T", "1", "10")
	cmd.Env = append(os.Environ(), "S_TIME_FORMAT=ISO")
	return cmd.Output()
}

func parse(rawMpstat []byte) ([]cpuStatistic, error) {
	var statistics mpstat
	err := json.Unmarshal(rawMpstat, &statistics)
	if err != nil {
		return nil, err
	}

	var result []cpuStatistic

	for i := range statistics.SysStat.Hosts {
		for j := range statistics.SysStat.Hosts[i].Statistics {
			statistics.SysStat.Hosts[i].Statistics[j].Timestamp = statistics.SysStat.Hosts[i].Date +
				" " + statistics.SysStat.Hosts[i].Statistics[j].Timestamp
			result = append(result, statistics.SysStat.Hosts[i].Statistics[j])
		}
	}

	fmt.Println(result)

	return result, nil
}
