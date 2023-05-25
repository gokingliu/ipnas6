package logic

import (
	"ipnas6/util"
)

func GetIPLogic() map[string]string {
	IPv4 := util.GetIPv4()
	IPv6 := util.GetIPv6()

	result := make(map[string]string)
	result["IPv4"] = IPv4
	result["IPv6"] = IPv6

	return result
}
