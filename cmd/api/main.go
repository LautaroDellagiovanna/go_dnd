package main

import (
	"go_dnd/internal/handlers"
	"go_dnd/pkg/config"
	"go_dnd/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
)

// Se necesita gin para esta aplicación
// Se debe ejecutar el siguiente comando para que
// el programa sea detectado por GO:
//		go mod init example/go_dnd
// Luego se debe instalar gin, el cual es un paquete de github con el siguiente comando
//		go get github.com/gin-gonic/gin
// Ahora sí se puede crear la aplicación.

// Convertir la struct a JSON. Para eso se utilizan los `json:"id"` que se encuentran al lado de la struct

// Crear el server
// 	router := gin.Default()
//	router.Run("localhost:9090")

func main() {
	router := gin.Default()

	cfg := config.LoadConfig()

	log.Printf("%s\n", cfg.DatabaseURL)
	log.Printf("%s\n", cfg.ServerAddress)
	db, err := database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v.", err)
		return
	}

	userHandler := handlers.NewUserHandler(db)

	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.AddUser)

	router.Run(cfg.ServerAddress)

}
