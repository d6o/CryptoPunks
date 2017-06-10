package main

import (
	"github.com/sirupsen/logrus"
	"github.com/kelseyhightower/envconfig"
	"fmt"
	"flag"
	"os"
)

var (
	app App
)

func init() {
	logrus.Info("Starting CriptoPunks")
	app = App{}
	logrus.Info("Loading configs")
	err := envconfig.Process("msb", &app.config)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.Infof("Configs: %v", app.config)
	fmt.Println("CriptoPunks running!")
}

func main() {
	cmd := flag.String("cmd", "", "Command to be executed: (choose)")
	flag.Parse()

	logrus.Infof("Command received: %s", *cmd)

	switch *cmd {
	case "scraper":
		scraperCMD()
	}

	logrus.Fatal("Not good...")
	os.Exit(1)
}
