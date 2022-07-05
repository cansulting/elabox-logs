package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

var tmpLogFile *os.File

const maintenancefile = "/tmp/tmplog"

// maintenance schedule in millisec. by default this is weekly
var defaultSched int64 = 7 * 24 * 60 * 60 * 1000

// call back when starting log maintenance
// @startRange in millisecods. only persist old logs within specific range.
// if value <= 0 then only persist logs that are within a week. example value is
// 2 * 24 * 60 * 60 * 1000 it means persist logs in last 2 days
func OnMaintenance(startRange int64) error {
	logger.GetInstance().Info().Msg("start log maintenance...")
	ctime := time.Now().UTC()
	if startRange <= 0 {
		startRange = defaultSched
	}

	var tmptime time.Time
	var err error
	dirty := false
	LogReader.LoadSeq(0, -1, logger.OLD_FIRST, func(l logger.Log) bool {
		if l["time"] == nil {
			println("invalid log")
			return false
		}
		tmptime, err = time.Parse(time.RFC3339, l["time"].(string))
		if err != nil {
			return true
		}
		diff := ctime.UnixMilli() - tmptime.UnixMilli()
		if diff < startRange {
			// all logs are fine. then no need to maintain
			if !dirty {
				return false
			}
			if err = write(l); err != nil {
				return false
			}
			return true
		}
		dirty = true
		return true
	})
	// remove old log file
	if dirty {
		if tmpLogFile != nil {
			tmpLogFile.Close()
			tmpLogFile = nil
			if err := logger.DeleteLogFile(); err != nil {
				return errors.New("failed deleting log file")
			}
			if err := os.Rename(maintenancefile, logger.LOG_FILE); err != nil {
				return err
			}
			logger.Reinit()
		}
		logger.GetInstance().Info().Msg("Log house keeping finalization...")
	} else {
		logger.GetInstance().Info().Msg("nothing to clean in log. maintenance skipped")
	}
	return nil
}

// write to tmp file
func write(l logger.Log) error {
	// step: create tmp log file if not yet created
	if tmpLogFile == nil {
		var err error
		tmpLogFile, err = os.OpenFile(maintenancefile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
	}
	// step: deserialize
	bytes, err := json.Marshal(l)
	if err != nil {
		return nil
	}
	if _, err := tmpLogFile.Write(bytes); err != nil {
		return err
	}
	if _, err := tmpLogFile.WriteString("\n"); err != nil {
		return err
	}
	return nil
}
