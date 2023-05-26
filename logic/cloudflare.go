package logic

import (
	"errors"
	"fmt"
	"ipnas6/util"
)

var token = util.LoaderGet("cloudflare.token")
var zoneId = util.LoaderGet("cloudflare.zoneId")
var hostName = util.LoaderGet("cloudflare.hostName")

var getUri = fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", zoneId)
var headers = make(map[string]string)

func GetCloudflareDNSId() (string, error) {
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	headers["Content-Type"] = "application/json"
	result, err := util.Get(getUri, headers, nil)
	if err != nil {
		return "", err
	}

	var dnsId string
	dnsStatus := result["success"].(bool)
	dnsResult := result["result"].([]interface{})
	if dnsStatus && len(dnsResult) > 0 {
		for _, dns := range dnsResult {
			dnsMap := dns.(map[string]interface{})
			if dnsMap["name"] == hostName {
				dnsId = dnsMap["id"].(string)
			}
		}

		return dnsId, nil
	}

	return "", errors.New("DNS 获取失败")
}

func UpdateCloudflareDNS(dnsId string, IPv6 string) bool {
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	headers["Content-Type"] = "application/json"
	putUri := fmt.Sprintf("%s/%s", getUri, dnsId)

	data := map[string]interface{}{
		"type":    "AAAA",
		"name":    hostName,
		"content": IPv6,
		"proxied": false,
		"ttl":     60,
	}
	result, err := util.Put(putUri, headers, data)
	if err != nil {
		return false
	}

	dnsStatus := result["success"].(bool)

	return dnsStatus
}
