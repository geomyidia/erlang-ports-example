package port

import (
	"math/rand"
	"time"

	"github.com/geomyidia/zylog/logger"
)

// SetupApp ...
func SetupLogging() {
	logger.SetupLogging(&logger.ZyLogOptions{
		Colored:      false,
		Level:        "fatal",
		Output:       "stderr",
		ReportCaller: true,
	})

}

func SetupRandom() {
	rand.Seed(time.Now().Unix())
}
