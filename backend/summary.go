// Copyright 2021 The Elabox Authors
// This file is part of elabox-logs library.

// elabox-logs library is under open source LGPL license.
// If you simply compile or link an LGPL-licensed library with your own code,
// you can release your application under any license you want, even a proprietary license.
// But if you modify the library or copy parts of it into your code,
// you’ll have to release your application under similar terms as the LGPL.
// Please check license description @ https://www.gnu.org/licenses/lgpl-3.0.txt

// handles data for current log summary

package main

import (
	"sync"

	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

type Summary struct {
	Levels     map[string]uint16 `json:"levels"`     // count for each log level
	Packages   map[string]uint16 `json:"packages"`   // count for all packages found in log
	Categories map[string]uint16 `json:"categories"` // count for all categoriest found in log
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
	LogReader.Load(0, -1, logger.LATEST_FIRST, func(i int, l logger.Log) bool {
		mutex.Lock()
		level = "debug"
		if l["level"] != nil {
			level = l["level"].(string)
		}
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
