package main

import (
	"github.com/cansulting/elabox-system-tools/foundation/app"
	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

// the current log reader
var LogReader, _ = logger.NewReader("")
var AppController *app.Controller

const PKID = "ela.logs"

// actions
const LOAD_FILTERS_AC = PKID + ".LOAD_FILTERS_ACTION"
const LOAD_LATEST_AC = PKID + ".LOAD_LATEST_ACTION"

// copy the log map to destination
func CopyLog(dst logger.Log, src logger.Log) {
	/* Copy Content from Map1 to Map2*/
	for key, value := range src {
		dst[key] = value
	}
}

// reuse map
func ResuseLog(src logger.Log) {
	for key := range src {
		delete(src, key)
	}
}
