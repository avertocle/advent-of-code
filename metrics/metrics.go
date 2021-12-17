package metrics

import (
	"fmt"
	"time"
)

var m *Metrics

type Metrics struct {
	StartTime int64
	EndTime   int64
	InputLen  int
	RecDepth  int
}

func init() {
	m = new(Metrics)
}

func ProgStart() {
	m.StartTime = time.Now().Unix()
}

func ProgEnd() {
	m.EndTime = time.Now().Unix()
}

func InputLen(x int) {
	m.InputLen = x
}

func RecIn() {
	m.RecDepth++
}

func RecOut() {
	m.RecDepth--
}

func I() *Metrics {
	return m
}

func ToString() string {
	return fmt.Sprintf("input-len (%v) | prog-time (%vms)", m.InputLen, m.EndTime-m.StartTime)
}
