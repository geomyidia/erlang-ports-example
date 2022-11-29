package echo

import (
	log "github.com/sirupsen/logrus"

	"github.com/geomyidia/erlcmd/pkg/messages"
)

func SendResult(resultMsg string) {
	r, err := messages.NewResultResponse(messages.Result(resultMsg))
	if err != nil {
		SendError(err)
		return
	}
	r.Send()
}

func SendError(err error) {
	if err == nil {
		return
	}
	log.Error(err)
	resp, _ := messages.NewErrorResponse(messages.Err(err.Error()))
	resp.Send()
}
