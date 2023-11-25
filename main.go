package main

import (
	"challenge/config/database"
	"challenge/interfaces/entries"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {

	_, err := database.ConectarBanco()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	entries.Rotas(r.Group("/api"))
	r.Run(":8080")

}
