package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/jcp100760/go_gorm/internal/factory"
)

func main() {
	// Construimos la aplicación mediante el factory (ensamblador)
	appFactory := factory.NewFactory()

	// Obtenemos el router ya configurado
	router := appFactory.Router()

	// Opcional: registro de CORS / logging simple
	loggedRouter := handlers.LoggingHandler(log.Writer(), router)

	// Ejecutar servidor
	addr := ":8080"
	log.Printf("Servidor iniciado en %s\n", addr)
	if err := http.ListenAndServe(addr, loggedRouter); err != nil {
		log.Fatalf("error al iniciar servidor: %v", err)
	}
}
