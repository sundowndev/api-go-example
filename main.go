/**
 * Created by fbonhomm
 * Email: flo-github@outlook.fr
 * License: MIT
 */

package main

import (
    "log"

    "github.com/joho/godotenv"

    "github.com/sundowndev/api-go-example/src/routers"
    "github.com/sundowndev/api-go-example/src/services"
)

// init
func init() {
    // Initialize environment variable contained in .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file found")
    }
}

// main
func main() {
    // Initialize Database
    services.Database()

    // Initialize Jwt
    services.Jwt()

    // Initialize router
    var routerEngine = routers.RouterInitialize()

    // Launch API
    if err := routerEngine.Run(); err != nil {
        log.Fatal(err)
    }

    services.Db.Close()
}
