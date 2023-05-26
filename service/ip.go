package service

import (
	"ipnas6/entity"
	"ipnas6/logic"
	"ipnas6/util"
	"net/http"
)

func getIP(res http.ResponseWriter, req *http.Request) {
	IPMap := logic.GetIPLogic()

	if IPMap["IPv4"] == "" && IPMap["IPv6"] == "" {
		_ = util.JSONResponse(res, entity.ClientIPError, nil)
		return
	}

	_ = util.JSONResponse(res, entity.ResOk, IPMap)
}
