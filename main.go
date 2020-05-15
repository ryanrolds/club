package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/ryanrolds/club/signaling"

	"github.com/sirupsen/logrus"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	if env == "prod" {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Setup logging
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return funcName, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
		},
	})

	logrus.Infof("Log level: %s", logrus.GetLevel())

	http.Handle("/room", &signaling.Server{})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	logrus.Info("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		logrus.Fatal(err)
	}
}
