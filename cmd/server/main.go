package main

import (
    "backend/go-gin/configs"
    "backend/go-gin/internal/database"
    "backend/go-gin/api/rest"
)

func main() {
    configs.LoadEnv()
    database.InitPostgres()

    router := rest.InitRoutes()
    router.Run(":8080") // Start on port 8080
}
