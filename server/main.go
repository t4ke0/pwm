package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TaKeO90/pwm/server/handler"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	//r.Get("/", handler.ReqHandler)
	r.Options("/user", handler.ReqHandler)
	r.Get("/user", handler.ReqHandler)
	r.Post("/register", handler.ReqHandler)
	r.Options("/register", handler.ReqHandler)
	r.Post("/login", handler.ReqHandler)
	r.Options("/login", handler.ReqHandler)
	r.Options("/show", handler.ReqHandler)
	r.Post("/show", handler.ReqHandler)
	r.Options("/forgot", handler.ReqHandler)
	r.Post("/forgot", handler.ReqHandler)
	r.Options("/creds", handler.ReqHandler)
	r.Post("/creds", handler.ReqHandler)
	r.Options("/logout", handler.ReqHandler)
	r.Get("/logout", handler.ReqHandler)
	r.Options("/upload", handler.ReqHandler)
	r.Post("/upload", handler.ReqHandler)
	r.Options("/genpw", handler.ReqHandler)
	r.Post("/genpw", handler.ReqHandler)
	fmt.Println("Running Server on 127.0.0.1:8080")
	err := (http.ListenAndServe(":8080", r))
	if err != nil {
		log.Fatal(err)
	}
}
