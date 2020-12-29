package main

import (
	"University-Information-Website/model"
	"University-Information-Website/routes"
)

func main() {
	model.InitDb()
	model.InitModel()
	routes.InitRouter()
}
