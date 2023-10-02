package main

import (
	"stock-api/api-portal/routes"
	"stock-api/global"
	"stock-api/repo"
)

func main() {
	// fetch env
	global.FetchEnvs()
	// init db
	repo.Init()

	// migrate db
	repo.DoMigration()

	// init routes
	routes.Init()
}
