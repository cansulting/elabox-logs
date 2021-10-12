package main

import "github.com/cansulting/elabox-system-tools/foundation/logger"

func RetrieveLogWithLimit(offset int64, limit int) {
	LogReader.LoadLimit(offset, limit, func(l logger.Log) bool {
		if filterLog(l) {
			onClientRecievedLog(0, l)
		}
		return true
	})
}

func ApplyFilter(json map[string]interface{}) {
	
}

// retrieve the latest ending offset in log file
func RetrieveLatestOffset() {

}

func onClientRecievedLog(chunkI int, l logger.Log) {

}

// use to filter log
// @l - the log data
// @return - true if the log will be included
func filterLog(l logger.Log) bool {
	return true
}
