package main

import (
	"strings"
	"sync"

	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

const LIMIT = 20

var filter map[string]interface{}
var filterLevels map[string]interface{}
var filterPackages map[string]interface{}
var filterCategories map[string]interface{}
var filterConditions []interface{}

// resuable pool of logs
var logPool = sync.Pool{
	New: func() interface{} {
		vals := make([]logger.Log, LIMIT)
		for i := 0; i < LIMIT; i++ {
			vals[i] = logger.Log{}
		}
		return vals
	},
}

func RetrieveLogWithLimit(offset int64) []logger.Log {
	output := logPool.Get().([]logger.Log)
	total := 0
	// load logs
	LogReader.LoadLimit(offset, LIMIT,
		// function callback when log was retrieved
		func(l logger.Log) bool {
			if filterLog(l) {
				s := output[total]
				CopyLog(s, l)
				total++
				return true
			}
			return false
		})
	// clear the unused indexes
	for i := LIMIT - 1; i >= total; i-- {
		ResuseLog(output[i])
	}
	return output
}

func ClearLogs(logs []logger.Log) {
	for i := 0; i < len(logs); i++ {
		ResuseLog(logs[i])
	}
	logPool.Put(logs)
}

func ApplyFilter(newFilter map[string]interface{}) {
	if len(newFilter) == 0 {
		filter = nil
		return
	}

	filter = newFilter
	filterLevels = nil
	filterPackages = nil
	filterCategories = nil
	filterConditions = nil
	if filter["levels"] != nil {
		filterLevels = filter["levels"].(map[string]interface{})
	}
	if filter["packages"] != nil {
		filterPackages = filter["packages"].(map[string]interface{})
	}
	if filter["categories"] != nil {
		filterCategories = filter["categories"].(map[string]interface{})
	}
	if filter["conditions"] != nil {
		filterConditions = filter["conditions"].([]interface{})
	}
	//filterLevels = filter["levels"].(map[string]interface{})
}

// retrieve the latest logs
func RetrieveLatestOffset() []logger.Log {
	return RetrieveLogWithLimit(0)
}

func onClientRecievedLog(chunkI int, l logger.Log) {

}

// use to check if key is toggled/true in map
func checkIfToggle(key string, _map map[string]interface{}, l logger.Log) bool {
	if l[key] == nil {
		return true
	}
	key2 := l[key].(string)
	return _map == nil || _map[key2] == nil || _map[key2].(bool)
}

// return true if all conditions are satisfied
func checkConditions(l logger.Log) bool {
	length := len(filterConditions)
	for i := 0; i < length; i++ {
		filterCon := filterConditions[i].(map[string]interface{})
		// if filterCon != nil && !filterCon["on"].(bool) {
		// 	continue
		// }
		// is field available?
		field := filterCon["key"].(string)
		if l[field] == nil {
			continue
		}
		fieldVal := strings.ToLower(l[field].(string))
		filterVal := strings.ToLower(filterCon["value"].(string))
		// do the operation
		switch filterCon["operator"] {
		case "contains":
			if !strings.Contains(fieldVal, filterVal) {
				return false
			}
		case "==":
			if fieldVal != filterVal {
				return false
			}
		case "!=":
			if fieldVal == filterVal {
				return false
			}
		case "not contains":
			if strings.Contains(fieldVal, filterVal) {
				return false
			}
		}
	}
	return true
}

// use to filter log
// @l - the log data
// @return - true if the log will be included
func filterLog(l logger.Log) bool {
	if filter != nil {
		// check level
		if !checkIfToggle("level", filterLevels, l) ||
			!checkIfToggle("package", filterPackages, l) ||
			!checkIfToggle("category", filterCategories, l) {
			return false
		}
		if !checkConditions(l) {
			return false
		}
	}
	return true
}
