package logic

import (
	"fmt"
	"net/http"
)

func GetIPV6(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
