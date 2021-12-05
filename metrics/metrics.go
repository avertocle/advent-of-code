package metrics

import (
	"fmt"
	"time"
)

var metrics *Metrics

type Metrics struct {
	StartTime int64
	EndTime   int64
}

func init() {
	metrics = new(Metrics)
}

func ProgStart() {
	metrics.StartTime = time.Now().Unix()
}

func ProgEnd() {
	metrics.EndTime = time.Now().Unix()
}

func ToString() string {
	return fmt.Sprintf("prog time = %vms", metrics.EndTime-metrics.StartTime)
}
