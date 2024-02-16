package main

import (
	"final_projek_go/database"
	"final_projek_go/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
