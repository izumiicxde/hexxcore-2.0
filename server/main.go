package main

import (
	"hexxcore/cmd/api"
	"hexxcore/config"
	"hexxcore/storage"
)

func main() {
	db := storage.NewPostgresStorage()
	storage.AutoMigrate(db)
	storage.InsertPredefinedSchedule(db)

	api := api.NewAPIServer(config.Envs.PORT, db)

	api.Run()
}
