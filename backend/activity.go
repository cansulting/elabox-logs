// Copyright 2021 The Elabox Authors
// This file is part of elabox-logs library.

// elabox-logs library is under open source LGPL license.
// If you simply compile or link an LGPL-licensed library with your own code,
// you can release your application under any license you want, even a proprietary license.
// But if you modify the library or copy parts of it into your code,
// you’ll have to release your application under similar terms as the LGPL.
// Please check license description @ https://www.gnu.org/licenses/lgpl-3.0.txt

// Implementation of elabox activity

package main

import (
	"os"
	"encoding/json"

	"github.com/cansulting/elabox-system-tools/foundation/app/rpc"
	"github.com/cansulting/elabox-system-tools/foundation/event/data"
	"github.com/cansulting/elabox-system-tools/foundation/event/protocol"
)

type Activity struct {
}

// callback when activity started
func (instance *Activity) OnStart(action *data.Action) error {
	// recieved requests from client
	AppController.RPC.OnRecieved(LOAD_FILTERS_AC, instance.OnAction_LoadFilters)
	AppController.RPC.OnRecieved(LOAD_LATEST_AC, instance.OnAction_LoadLatest)
	AppController.RPC.OnRecieved(LOAD_RANGE_AC, instance.OnAction_LoadRange)
	AppController.RPC.OnRecieved(DELETE_LOG_FILE_AC, instance.OnAction_DeleteLogFile)
	return nil
}

func (instance *Activity) IsRunning() bool {
	return true
}
func (instance *Activity) OnEnd() error {
	return nil
}

// callback from client. this delete the log file
func (instance *Activity) OnAction_DeleteLogFile(client protocol.ClientInterface, data data.Action) string {
	e := os.Remove(ELA_LOG_FILE_LOC)
    if e != nil {
		return rpc.CreateResponseQ(rpc.SYSTEMERR_CODE, e.Error(), false)
    }
	_, err := os.Create(ELA_LOG_FILE_LOC)
	if err != nil {
		return rpc.CreateResponseQ(rpc.SYSTEMERR_CODE, e.Error(), false)
    }
	return rpc.CreateResponseQ(rpc.SUCCESS_CODE, "success", false)
	

}


// callback from client. this load the filters
func (instance *Activity) OnAction_LoadFilters(client protocol.ClientInterface, data data.Action) string {
	summary := LoadLogSummary()
	output, err := json.Marshal(summary)
	if err != nil {
		return rpc.CreateResponseQ(rpc.SYSTEMERR_CODE, err.Error(), false)
	}
	return rpc.CreateResponseQ(rpc.SUCCESS_CODE, string(output), false)
}

// callback from client. this load the latest logs
// @optional map - contains filters and offset. Offset is the start of retrieval
func (instance *Activity) OnAction_LoadLatest(client protocol.ClientInterface, data data.Action) string {
	filters, err := data.DataToMap(nil)
	if err != nil {
		return err.Error()
	}
	ApplyFilter(filters)
	var offset int64 = 0
	if filters["offset"] != nil {
		val := filters["offset"].(float64)
		offset = int64(val)
	}
	res, err := RetrieveFromOffset(offset)
	if err != nil {
		return rpc.CreateResponseQ(rpc.SYSTEMERR_CODE, err.Error(), false)
	}
	return rpc.CreateResponseQ(rpc.SUCCESS_CODE, res, false)
}

// callback from client. this load the latest logs
// @optional map - contains filters, offset and limit. Offset is the start of retrieval. limit is the end offset
func (instance *Activity) OnAction_LoadRange(client protocol.ClientInterface, data data.Action) string {
	filters, err := data.DataToMap(nil)
	if err != nil {
		return err.Error()
	}
	ApplyFilter(filters)
	var offset int64 = 0
	var length int64 = 0
	if filters["offset"] != nil {
		val := filters["offset"].(float64)
		offset = int64(val)
	}
	if filters["limit"] != nil {
		val := filters["limit"].(float64)
		length = int64(val)
	}
	res, err := RetrieveFromRange(offset, length)
	if err != nil {
		return rpc.CreateResponseQ(rpc.SYSTEMERR_CODE, err.Error(), false)
	}
	return rpc.CreateResponseQ(rpc.SUCCESS_CODE, res, false)
}
