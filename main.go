package main

import (
	"log"

	"github.com/alwyalhaddad/belajar-golang-post/config"
	"github.com/alwyalhaddad/belajar-golang-post/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := config.ConnectDatabase()
	router := gin.Default()
	routes.MainRoutes(router, db)
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run Gin server: %v", err)
	}
}
