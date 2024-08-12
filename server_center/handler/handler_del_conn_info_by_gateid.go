package handler

import (
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/mjczz/chess/pb/center"
	"github.com/mjczz/chess/server_center/conn_info"
)

var delConnInfoByGateidResp *center.DelConnInfoByGateidResp = &center.DelConnInfoByGateidResp{}

func HandleDelConnInfoByGateid(client io.Writer, req proto.Message) error {
	delConnInfoByGateidReq, ok := req.(*center.DelConnInfoByGateidReq)
	if !ok || delConnInfoByGateidReq.Gateid == 0 {
		return nil
	}

	conn_info.DelByGateid(delConnInfoByGateidReq.Gateid)
	sendDelConnInfoByGateidNotify(delConnInfoByGateidReq.Gateid, client)
	return sendResp(client, delConnInfoByGateidResp)
}
