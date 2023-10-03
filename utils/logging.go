package utils

import (
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

var Log = log.New()

func init() {

	Log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	Log.SetLevel(log.DebugLevel)

	Log.Out = os.Stdout
	// log to a file, or fall back to default stderr

	// create log file
	logFile, err := os.OpenFile(filepath.Join("data/logs", time.Now().Format("2006-01-02")+".log"), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Log.WithError(err).Error("failed to create log file")
	}

	Log.Out = logFile

	Log.WithField("log_file", logFile.Name()).Infof("------- Starting Ulsidor CLI %s -------", time.Now().Format("2006-01-02 15:04:05"))

}
