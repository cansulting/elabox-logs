package main

import (
	"encoding/json"

	"github.com/cansulting/elabox-system-tools/foundation/app/rpc"
	"github.com/cansulting/elabox-system-tools/foundation/event/data"
	"github.com/cansulting/elabox-system-tools/foundation/event/protocol"
)

type Activity struct {
}

func (instance *Activity) OnStart(action *data.Action) error {
	// recieved requests from client
	AppController.RPC.OnRecieved(LOAD_FILTERS_AC, instance.OnAction_LoadFilters)
	AppController.RPC.OnRecieved(LOAD_LATEST_AC, instance.OnAction_LoadLatest)
	return nil
}

func (instance *Activity) IsRunning() bool {
	return true
}
func (instance *Activity) OnEnd() error {
	return nil
}

// callback from client. this load the filters
func (instance *Activity) OnAction_LoadFilters(client protocol.ClientInterface, data data.Action) string {
	summary := LoadLogSummary()
	output, err := json.Marshal(summary)
	if err != nil {
		return err.Error()
	}
	return rpc.CreateResponseQ(rpc.SUCCESS_CODE, string(output), false)
}

// callback from client. this load the latest logs
func (instance *Activity) OnAction_LoadLatest(client protocol.ClientInterface, data data.Action) string {
	filters, err := data.DataToMap(nil)
	if err != nil {
		return err.Error()
	}
	ApplyFilter(filters)
	res := RetrieveLatestOffset()
	output, err := json.Marshal(res)
	ClearLogs(res)
	if err != nil {
		return err.Error()
	}
	return rpc.CreateResponseQ(rpc.SUCCESS_CODE, string(output), false)
}
