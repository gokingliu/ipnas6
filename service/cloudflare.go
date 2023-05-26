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
		_ = util.JSONResponse(res, entity.ClientGetError, nil)
		return
	}

	var dnsId string
	dnsStatus := result["success"].(bool)
	dnsResult := result["result"].([]interface{})
	if dnsStatus && len(dnsResult) > 0 {
		fmt.Println(1111, dnsResult)
		for _, dns := range dnsResult {
			dnsMap := dns.(map[string]interface{})
			if dnsMap["name"] == "router.crotaliu.top" {
				fmt.Println(2222, dnsMap["id"])
				dnsId = dnsMap["id"].(string)
			}
		}
		_ = util.JSONResponse(res, entity.ResOk, dnsResult)
	} else {
		_ = util.JSONResponse(res, entity.ClientGetDNSError, nil)
	}

	fmt.Println(3333, dnsId)
}
