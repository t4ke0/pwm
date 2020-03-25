package main

import (
	"./handler"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", handler.HomeHandle)
	r.Get("/register", handler.RegisterHandle)
	r.Post("/register", handler.RegisterHandle)
	r.Get("/login", handler.LoginHandle)
	r.Post("/login", handler.LoginHandle)
	r.Get("/show", handler.ShowStuffHandle)
	r.Post("/show", handler.ShowStuffHandle)
	r.Get("/update", handler.UpdateStuffHandle)
	r.Post("/update", handler.UpdateStuffHandle)
	r.Get("/add", handler.AddStuffHandle)
	r.Post("/add", handler.AddStuffHandle)
	r.Get("/delete", handler.DeleteStuffHandle)
	r.Post("/delete", handler.DeleteStuffHandle)
	r.Get("/logout", handler.LogoutHandle)
	fmt.Println("Running Server on 127.0.0.1:8080")
	err := (http.ListenAndServe(":8080", r))
	if err != nil {
		log.Fatal(err)
	}
}
