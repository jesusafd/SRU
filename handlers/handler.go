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
	router.HandleFunc("/user", middlewares.CheckDB(routes.ReadUser)).Methods("GET")
	router.HandleFunc("/user", middlewares.CheckDB(routes.UpdateUser)).Methods("PUT")
	router.HandleFunc("/user", middlewares.CheckDB(routes.ReadUser)).Methods("GET")

	PORT := "8080"

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
