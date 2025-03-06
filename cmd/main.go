package main

import (
	"log"
	"os"

	"github.com/cestevezing/veloces/internal/core/common/router"
	"github.com/cestevezing/veloces/internal/core/common/utils"
	"github.com/cestevezing/veloces/internal/infra/redis_service"
	"github.com/cestevezing/veloces/internal/infra/repository"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func main() {
	db, err := repository.InitDB()
	if err != nil {
		panic(err)
	}
	utils.InitValidator()
	redisClient := redis_service.NewRedisClient()
	app := SetupApp(db, redisClient)
	port := os.Getenv("PORT")
	log.Println(app.Listen(":" + port))
}

func SetupApp(database *gorm.DB, redisClient *redis.Client) *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "Api Veloces",
	})
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} \n",
	}))
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "../docs/swagger.json",
		Path:     "/",
		Title:    "Swagger API Documentation",
	}
	app.Use(swagger.New(cfg))
	routes := router.NewRouter(database, app, redisClient)
	routes.InitRoutes()
	return app
}
