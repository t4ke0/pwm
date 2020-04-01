package main

import (
	"fmt"
	"log"
	"net/http"

	"./handler"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", handler.ReqHandler)
	r.Get("/register", handler.ReqHandler)
	r.Post("/register", handler.ReqHandler)
	r.Get("/login", handler.ReqHandler)
	r.Post("/login", handler.ReqHandler)
	r.Get("/show", handler.ReqHandler)
	r.Post("/show", handler.ReqHandler)
	r.Get("/update", handler.ReqHandler)
	r.Post("/update", handler.ReqHandler)
	r.Get("/add", handler.ReqHandler)
	r.Post("/add", handler.ReqHandler)
	r.Get("/delete", handler.ReqHandler)
	r.Post("/delete", handler.ReqHandler)
	r.Get("/logout", handler.ReqHandler)
	fmt.Println("Running Server on 127.0.0.1:8080")
	err := (http.ListenAndServe(":8080", r))
	if err != nil {
		log.Fatal(err)
	}
}
