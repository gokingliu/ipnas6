package service

import (
	"fmt"
	"ipnas6/entity"
	"ipnas6/util"
	"net/http"
)

func updateCFIPv6(res http.ResponseWriter, req *http.Request) {
	token := util.LoaderGet("cloudflare.token")
	zoneId := util.LoaderGet("cloudflare.zoneId")
	uri := fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", zoneId)
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	headers["Content-Type"] = "application/json"

	result, err := util.Get(uri, headers, nil)
	if err != nil {
		fmt.Println("111")
	}
	fmt.Println(result)
	_ = util.JSONResponse(res, entity.ResOk, result)
}
