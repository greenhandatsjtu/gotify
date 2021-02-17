package logs

import (
	log "github.com/sirupsen/logrus"
)

//初始化logger
func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}
