package main

import (
	"fmt"
	"log"
	"net/http"

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

	initLogging(config)

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

	// In production this service is just a websocket service
	if config.Environment != EnvironmentProduction {
		fs := http.FileServer(http.Dir("./static"))
		http.Handle("/", NoCache(fs))
	}

	logrus.Infof("Listening on :%d...", config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
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
