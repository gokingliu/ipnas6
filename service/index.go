package service

import "net/http"

func Register() {
	http.HandleFunc("/GetIP", getIP)
	http.HandleFunc("/UpdateCFIPv6", updateCFIPv6)
}
