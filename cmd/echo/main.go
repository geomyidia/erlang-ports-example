package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlang-port-examples/internal/app"
	"github.com/geomyidia/erlang-port-examples/pkg/echo"
	"github.com/geomyidia/erlang-port-examples/pkg/port"
)

func main() {
	log.Info("Starting up Go Port example ...")
	app.SetupLogging()
	app.SetupRandom()
	port.ProcessPortMessages(echo.ProcessEchoCommand)
}
