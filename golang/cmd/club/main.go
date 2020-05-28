package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/ryanrolds/club/pkg/signaling"

	"github.com/sirupsen/logrus"
)

const reaperInterval = time.Second * 15

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

	var room = signaling.NewRoom()
	room.StartReaper(reaperInterval)

	http.Handle("/room", signaling.NewServer(room))

	logrus.Info("Listening on :3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		logrus.Fatal(err)
	}
}

func NoCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=0, public, must-revalidate, proxy-revalidate")
		h.ServeHTTP(w, r)
	})
}
