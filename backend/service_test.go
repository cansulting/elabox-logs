package main

import (
	"testing"

	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

func TestMaintenance(t *testing.T) {
	logger.InitFromFile("ela.testing", "")
	OnMaintenance()
}
