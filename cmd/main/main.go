package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/juanmachuca95/classdigitalhouse/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar las variables de entorno")
	}

	//cacheMemory := NewCacheMemory()
	routes := routes.NewRoutes()

	//log.Println("Cache", cacheMemory.cache)
	log.Printf("API - %s:%s", os.Getenv("HOSTNAME"), os.Getenv("PORT"))
	http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal("No se puede inicializar el servidor - error: ", err)
	}
}
