package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"
	"strings"

	"github.com/ryanrolds/club/pkg/signaling"

	"github.com/sirupsen/logrus"
)

const (
	configFilename = "club.yaml"
)

func main() {
	config, err := GetConfig(configFilename)
	if err != nil {
		log.Fatal("problem reading club.yaml")
	}

	if config.Environment == EnvironmentProduction {
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
	room.StartReaper(config.ReaperInterval)

	err = room.AddGroup(signaling.NewGroup(signaling.RoomDefaultGroupID, config.DefaultGroupLimit))
	if err != nil {
		logrus.Fatal(err)
	}

	for _, groupConfig := range config.Groups {
		err = room.AddGroup(signaling.NewGroup(groupConfig.ID, groupConfig.Limit))
		if err != nil {
			logrus.Fatal(err)
		}
	}

	http.Handle("/room", signaling.NewServer(room))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", NoCache(fs))

	logrus.Info("Listening on :3001...")
	err = http.ListenAndServe(":3001", nil)
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
