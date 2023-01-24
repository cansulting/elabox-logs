// Copyright 2021 The Elabox Authors
// This file is part of elabox-logs library.

// elabox-logs library is under open source LGPL license.
// If you simply compile or link an LGPL-licensed library with your own code,
// you can release your application under any license you want, even a proprietary license.
// But if you modify the library or copy parts of it into your code,
// you’ll have to release your application under similar terms as the LGPL.
// Please check license description @ https://www.gnu.org/licenses/lgpl-3.0.txt

// file contains global values for log and helper functions

package main

import (
	"github.com/cansulting/elabox-system-tools/foundation/app"
	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

// the current log reader
var LogReader, _ = logger.NewReader("")
var AppController *app.Controller

const CHUNK_SIZE_PER_PAGE = logger.CHUNK_SIZE / 10

const PKID = "ela.logs"

// actions
const LOAD_FILTERS_AC = PKID + ".LOAD_FILTERS_ACTION"
const LOAD_LATEST_AC = PKID + ".LOAD_LATEST_ACTION"
const LOAD_RANGE_AC = PKID + ".LOAD_RANGE_ACTION"
const DELETE_LOG_FILE_AC = PKID + ".DELETE_LOG_FILE_ACTION"

// action broadcast
const SUMMARY_AC = PKID + ".broadcast.LOG_SUMMARY"

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
