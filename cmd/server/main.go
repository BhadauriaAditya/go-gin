package main

import (
    "github.com/aditya/go-gin/configs"
    "github.com/aditya/go-gin/internal/database"
)

func main() {
    configs.LoadEnv()
    database.InitPostgres()
}
