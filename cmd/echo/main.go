package main

import (
	"github.com/geomyidia/erlang-port-examples/echo"
	"github.com/geomyidia/erlang-port-examples/port"
)

func main() {
	port.SetupLogging()
	port.SetupRandom()
	port.ProcessPortMessages(echo.ProcessEchoCommand)
}
