package main

import (
	"github.com/alwyalhaddad/belajar-golang-post/config"
	"github.com/alwyalhaddad/belajar-golang-post/routes"
)

func main() {
	db := config.NewDB()
	router := routes.ApiRouter()
}
