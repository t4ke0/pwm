package handler

import (
	"fmt"
	"github.com/TaKeO90/pwm/server/dialer"
	"net/http"
)

func UpdateStuffHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
	if r.URL.Path != "/update/" {
		fmt.Fprintf(w, "404 Not Found", http.StatusNotFound)
	}
	dialer.ServeUpdate(w, r)
}
