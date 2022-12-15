package main

import (
	"log"

	"github.com/jesusafd/SRU/connections"
	"github.com/jesusafd/SRU/handlers"
)

func main() {
	connections.NewConnection()
	defer connections.DBConn.Close()
	if !connections.CheckConnection() {
		log.Fatal("Error al conectar con la base de datos")
	}
	handlers.Handler()
}
