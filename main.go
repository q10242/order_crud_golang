package main

import (
	"order_crud/database"
	"order_crud/migrations"
	"order_crud/routes"
)

func main() {
	database.ConnectDatabase()
	migrations.Migrate()
	r := routes.SetupRouter()
	r.Run(":8080")
}
