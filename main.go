package main

import (
	"os"
	"strings"
	"time"

	"github.com/rumblefrog/go-a2s"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) != 2 {
		logrus.Errorln("you must provide the host:port as 1 argument. Example: ./srcds-watchdog 151.80.111.130:27015")
		os.Exit(3)
	}
	serverHost := os.Args[1]
	client, err := a2s.NewClient(serverHost, a2s.TimeoutOption(time.Second*2))
	if err != nil {
		panic(err)
	}
	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	info, err := client.QueryInfo() // QueryInfo, QueryPlayer, QueryRules
	if err != nil {
		if strings.Contains(err.Error(), "i/o timeout") || strings.Contains(err.Error(), "connection refused") {
			logrus.Errorf("%s is offline", os.Args[1])
			os.Exit(1)
		}
		panic(err)
	}
	logrus.Printf("%s is online", os.Args[1])
	logrus.Printf("Hostname: %s", info.Name)
	logrus.Printf("Players: %d/%d", info.Players, info.MaxPlayers)
	logrus.Printf("Game: %s", info.Game)
	logrus.Printf("Version: %s", info.Version)
}
