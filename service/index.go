package service

import "net/http"

func Register() {
	http.HandleFunc("/GetIP", getIP)
}
