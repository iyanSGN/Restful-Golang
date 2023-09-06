package main

import (
	"new-project2/db"
	"new-project2/routes"
)

func main() {
	db.ConnDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}