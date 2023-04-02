package service

import (
	"ipnas6/logic"
	"net/http"
)

func Register() {
	http.HandleFunc("/GetIPV6", logic.GetIPV6)
}
