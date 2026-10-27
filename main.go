package main

import (
	"errors"
	"os"
	"syscall"
	"time"

	"github.com/nikooo777/blaster/valve"
	"github.com/sirupsen/logrus"
)

const (
	exitOnline = iota
	exitOffline
	exitMalfunction
	exitBadArguments
)

const queryTimeout = 2 * time.Second

func main() {
	if len(os.Args) != 2 {
		logrus.Errorln("you must provide the host:port as 1 argument. Example: ./srcds-watchdog 54.37.245.51:27045")
		os.Exit(exitBadArguments)
	}
	os.Exit(run(os.Args[1]))
}

func run(serverHost string) int {
	querier, err := valve.NewServerQuerier(serverHost, queryTimeout)
	if err != nil {
		logrus.Errorf("failed to create querier for %s: %s", serverHost, err)
		return exitMalfunction
	}
	defer querier.Close()

	info, err := querier.QueryInfo()
	if err != nil {
		if isOffline(err) {
			logrus.Errorf("%s is offline", serverHost)
			return exitOffline
		}
		logrus.Errorf("failed to query %s: %s", serverHost, err)
		return exitMalfunction
	}
	logrus.Infof("%s is online", serverHost)
	logrus.Infof("Hostname: %s", info.Name)
	logrus.Infof("Players: %d/%d", info.Players, info.MaxPlayers)
	logrus.Infof("Game: %s", info.Game)
	if info.Ext != nil {
		logrus.Infof("Version: %s", info.Ext.GameVersion)
	}
	return exitOnline
}

func isOffline(err error) bool {
	return errors.Is(err, os.ErrDeadlineExceeded) || errors.Is(err, syscall.ECONNREFUSED)
}
