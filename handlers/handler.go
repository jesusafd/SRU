package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jesusafd/SRU/middlewares"
	"github.com/jesusafd/SRU/routes"
	"github.com/rs/cors"
)

func Handler() {
	router := mux.NewRouter()

	router.HandleFunc("/user", middlewares.CheckDB(routes.InsertUser)).Methods("POST")

	PORT := "8080"

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
