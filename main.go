package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"ipnas6/logic"
	"ipnas6/service"
	"net/http"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func cronTimer() {
	i := 0
	cronObject := newWithSeconds()
	spec := "0 */1 * * * ?" //一分钟运行一次
	cronObject.AddFunc(spec, func() {
		i++
		IPMap := logic.GetIPLogic()
		if IPMap["IPv6"] != "" {
			dnsId, err := logic.GetCloudflareDNSId()
			if err == nil {
				_ = logic.UpdateCloudflareDNS(dnsId, IPMap["IPv6"])
			}
		}
	})
	cronObject.Start()

	select {}
}

func main() {
	cronTimer()

	service.Register()

	err := http.ListenAndServe(":8688", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v", err)
	}
}
