package main

import (
	"go_gorm_graphql/db"
	"go_gorm_graphql/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
