package main

import (
	"sync"

	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

type Summary struct {
	Levels     map[string]uint16        `json:"levels"`
	Packages   map[string]uint16        `json:"packages"`
	Categories map[string]uint16        `json:"categories"`
}

// load log summary
func LoadLogSummary() Summary {
	filter := Summary{}
	filter.Levels = make(map[string]uint16)
	filter.Packages = make(map[string]uint16)
	filter.Categories = make(map[string]uint16)
	mutex := &sync.RWMutex{}

	var level string
	var pkg string
	var cat interface{}
	// iterate all logs
	LogReader.Load(0, -1, func(i int, l logger.Log) bool {
		mutex.Lock()
		level = l["level"].(string)
		pkg = l["package"].(string)
		cat = l["category"]
		filter.Levels[level]++
		filter.Packages[pkg]++
		if cat != nil {
			filter.Categories[cat.(string)]++
		}
		mutex.Unlock()
		return true
	})
	return filter
}
