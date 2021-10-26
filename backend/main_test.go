// Copyright 2021 The Elabox Authors
// This file is part of elabox-logs library.

// elabox-logs library is under open source LGPL license.
// If you simply compile or link an LGPL-licensed library with your own code,
// you can release your application under any license you want, even a proprietary license.
// But if you modify the library or copy parts of it into your code,
// you’ll have to release your application under similar terms as the LGPL.
// Please check license description @ https://www.gnu.org/licenses/lgpl-3.0.txt

package main

import (
	"encoding/json"
	"testing"
)

// test log retrieval
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

// test filter retrieval
func TestRetrieveFilter(t *testing.T) {
	summary := LoadLogSummary()
	json, err := json.Marshal(summary)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json)
}
