package main

import (
	"log"
	"os"
	"simple-api/internal/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	srv := server.NewServer()
	log.Printf("Servidor rodando na porta %s...\n", port)
	log.Fatal(srv.Start(":" + port))
}
