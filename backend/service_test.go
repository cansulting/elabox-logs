package main

import (
	"testing"
	"time"

	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

var sched int64 = 4 * 24 * 60 * 60 * 1000 // only persist logs that are 4 days old

// test maintenance with async feature
func TestMaintenance(t *testing.T) {
	logger.InitFromFile("ela.testing", "")
	done := false
	go func() {
		if err := OnMaintenance(sched); err != nil {
			t.Error(err)
		}
		done = true
	}()

	for !done {
		logger.GetInstance().Debug().Msg("Testing")
		time.Sleep(time.Millisecond * 100)
	}
}
