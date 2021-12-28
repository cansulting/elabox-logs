package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/cansulting/elabox-system-tools/foundation/logger"
)

var tmpLogFile *os.File

const maintenancefile = "tmplog"

// call back when starting log maintenance
func OnMaintenance() error {
	//logger.GetInstance().Info().Msg("start log maintenance...")
	ctime := time.Now().UTC()
	var weeklyMil int64 = 7 * 24 * 60 * 60 * 1000
	var tmptime time.Time
	var err error
	dirty := false
	LogReader.LoadSeq(0, -1, logger.OLD_FIRST, func(l logger.Log) bool {
		tmptime, err = time.Parse(time.RFC3339, l["time"].(string))
		if err != nil {
			return true
		}
		diff := ctime.UnixMilli() - tmptime.UnixMilli()
		if diff < weeklyMil {
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
	// remove the old log file
	if dirty {
		println("Needs to clean logs")
		if tmpLogFile != nil {
			tmpLogFile.Close()
			tmpLogFile = nil
			if err := logger.DeleteLogFile(); err != nil {
				return errors.New("failed deleting log file")
			}
			if err := os.Rename(maintenancefile, logger.LOG_FILE); err != nil {
				return err
			}
		}
	}
	return nil
}

func write(l logger.Log) error {
	if tmpLogFile == nil {
		var err error
		tmpLogFile, err = os.OpenFile(maintenancefile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
		if err != nil {
			return err
		}
	}
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
