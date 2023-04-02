package main

import (
	"fmt"
	"ipnas6/service"
	"net/http"
)

func main() {
	service.Register()

	err := http.ListenAndServe(":8688", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v", err)
	}
}
