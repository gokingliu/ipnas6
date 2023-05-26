package service

import (
	"ipnas6/entity"
	"ipnas6/logic"
	"ipnas6/util"
	"net/http"
)

func updateCFIPv6(res http.ResponseWriter, req *http.Request) {
	IPMap := logic.GetIPLogic()

	if IPMap["IPv6"] == "" {
		_ = util.JSONResponse(res, entity.ClientIPError, nil)
		return
	}

	dnsId, err := logic.GetCloudflareDNSId()
	if err != nil {
		_ = util.JSONResponse(res, entity.ClientGetDNSError, nil)
		return
	}

	result := logic.UpdateCloudflareDNS(dnsId, IPMap["IPv6"])

	if result {
		_ = util.JSONResponse(res, entity.ResOk, nil)
	} else {
		_ = util.JSONResponse(res, entity.ClientPutDNSError, nil)
	}
}
