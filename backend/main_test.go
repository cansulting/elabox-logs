package main

import (
	"encoding/json"
	"testing"
)

func TestRetrieveLogs(t *testing.T) {
	res := RetrieveLatestOffset()
	output, err := json.Marshal(res)
	ClearLogs(res)
	if err != nil {
		t.Error(output)
		return
	}
	t.Log(output)
}

func TestRetrieveFilter(t *testing.T) {
	summary := LoadLogSummary()
	json, err := json.Marshal(summary)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json)
}
