package main

import (
	"assignment_02/database"
	"assignment_02/routers"
)

func main() {
	var PORT = ":8000"

	database.StartDB()

	routers.StartServer().Run(PORT)
}
