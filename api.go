package main

import (
	"log"
	"net/http"

	// "newjwt/handler"

	"github.com/Waire214/newjwt/handler"
	"github.com/go-chi/chi/v5"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// // var err error
// var response ResponseObject

func routerHandler() http.Handler {
	router := chi.NewRouter()
	router.Post("/Signup", handler.SignUpHandler)
	router.Post("/signin", handler.SignInHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
	return router
}
func main() {
	routerHandler()
}
