package handler

import (
	"github.com/mjczz/chess/common"
	"github.com/mjczz/chess/game/server"
)

func HandleEcho(userid uint32, connid uint32, msgBody []byte) {
	server.SendResp(userid, MsgidEchoResp, common.ResultSuccess, msgBody)
}
