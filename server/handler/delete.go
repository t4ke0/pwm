package handler

import (
	"fmt"
	"github.com/TaKeO90/pwm/server/dialer"
	"net/http"
)

func DeleteStuffHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%s] %s\n", r.Method, r.URL.Path)
	if r.URL.Path != "/delete/" {
		fmt.Fprintf(w, "404 Not Found", http.StatusNotFound)
	}
	dialer.ServeDelete(w, r)
}
