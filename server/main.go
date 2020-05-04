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
	//r.Get("/", handler.ReqHandler)
	r.Options("/user", handler.ReqHandler)
	r.Get("/user", handler.ReqHandler)
	r.Post("/register", handler.ReqHandler)
	r.Options("/register", handler.ReqHandler)
	r.Post("/login", handler.ReqHandler)
	r.Options("/login", handler.ReqHandler)
	r.Options("/show", handler.ReqHandler)
	r.Post("/show", handler.ReqHandler)
	r.Options("/update", handler.ReqHandler)
	r.Post("/update", handler.ReqHandler)
	r.Options("/add", handler.ReqHandler)
	r.Post("/add", handler.ReqHandler)
	r.Options("/delete", handler.ReqHandler)
	r.Post("/delete", handler.ReqHandler)
	r.Options("/forgot", handler.ReqHandler)
	r.Post("/forgot", handler.ReqHandler)
	fmt.Println("Running Server on 127.0.0.1:8080")
	err := (http.ListenAndServe(":8080", r))
	if err != nil {
		log.Fatal(err)
	}
}
