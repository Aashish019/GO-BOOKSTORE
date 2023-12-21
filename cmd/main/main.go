package main

import (
	"log"
	"net/http"

	"github.com/Aashish019/GO-BOOKSTORE/pkg/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	router.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
