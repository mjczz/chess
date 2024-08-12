package handler

import (
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/mjczz/chess/pb/center"
	"github.com/mjczz/chess/server_center/conn_info"
)

var delConnInfoResp *center.DelConnInfoResp = &center.DelConnInfoResp{}

func HandleDelConnInfo(client io.Writer, req proto.Message) error {
	delConnInfoReq, ok := req.(*center.DelConnInfoReq)
	if !ok {
		return nil
	}

	gateid := delConnInfoReq.Gateid
	connid := delConnInfoReq.Connid
	if userid, ok := conn_info.Del(gateid, connid); ok {
		sendDelConnInfoNotify(&center.ConnInfo{userid, gateid, connid}, client)
	}

	return sendResp(client, delConnInfoResp)
}
