package main

// mpstat -P ALL -T -o JSON
type mpstat struct {
	SysStat struct {
		Hosts []struct {
			NodeName     string `json:"nodename"`
			SysName      string `json:"sysname"`
			Release      string `json:"release"`
			Machine      string `json:"machine"`
			NumberOfCpus int    `json:"number-of-cpus"`
			Date         string `json:"date"`
			Statistics   []struct {
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
			} `json:"statistics"`
		} `json:"hosts"`
	} `json:"sysstat"`
}

// mpstat -P ALL -o JSON -T 1 60
