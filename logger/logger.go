package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
    // "github.com/Franco-Poveda/logrus-splunk-hook"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	// "github.com/weekface/mgorus"
	lib "github.com/paulsjohnson91/challenge-accepted/shared"
)

//Logger hook
func Logger() *log.Logger {
	if err := godotenv.Load(); err != nil {
		log.Info("Error loading .env file! Try get a path...")
		if err2 := godotenv.Load(lib.GetPath() + "/.env"); err2 != nil {
			log.Info("Fail...")
			//os.Exit(1)
		}
	}
	logger := log.New()
	// hooker, err := mgorus.NewHookerWithAuth("127.0.0.1:27017", "gorest", "logs", "adminuser", "admpass1")
	// if err == nil {
	// 	logger.Hooks.Add(hooker)
	// } else {
	// 	fmt.Print(err)
	// }
	logger.SetReportCaller(true)
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {

			repopath := fmt.Sprintf("%s/src/github.com/paulsjohnson91/challenge-accepted/", os.Getenv("GOPATH"))
			filename := strings.Replace(f.File, repopath, "", -1)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	// logger.SetFormatter(&log.JSONFormatter{})


	// splunkClient := splunk.NewClient(nil, cfg.SplunkUrl, cfg.SplunkToken, cfg.SplunkSource, cfg.SplunkSourcetype, cfg.SplunkIndex)
	// splunkHook := splunk.NewHook(splunkClient, []log.Level{InfoLevel})
	// log.AddHook(splunkHook)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// logger.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// logger.SetLevel(log.InfoLevel)
	return logger
}

// func (h *Hook) Fire(entry *logrus.Entry) error {
// 	event := map[string]interface{}{
// 		"message": entry.Message,
// 		"time": entry.Time.String(),
// 		"level": entry.Level.String(),

// 	}
// 	for k, v := range entry.Data {
// 		event[k] = v
// 	}

// 	err := h.Client.Log(
// 		event,
// 	)
// 	return err
// }