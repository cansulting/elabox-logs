package main

import "github.com/cansulting/elabox-system-tools/foundation/event/data"

func BroadcastSummary(summary interface{}) {
	AppController.RPC.CallBroadcast(data.NewAction(SUMMARY_AC, "", summary))
}
