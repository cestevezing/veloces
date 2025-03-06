package router

import (
	"github.com/cestevezing/veloces/internal/infra/repository/data"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Router struct {
	DB          *gorm.DB
	App         *fiber.App
	RouterGroup fiber.Router
	RedisClient *redis.Client
}

func NewRouter(db *gorm.DB, app *fiber.App, redisClient *redis.Client) *Router {
	return &Router{
		DB:          db,
		App:         app,
		RedisClient: redisClient,
	}
}

func (r *Router) InitRoutes() {
	r.SetupGroup()
	r.SeedData()
	r.BuildProductRoutes()
	r.BuildOrderRoutes()
}

func (r *Router) SetupGroup() {
	r.RouterGroup = r.App.Group("/api")
}

func (r *Router) SeedData() {
	dataLoader := data.NewDataLoader(r.DB)
	dataLoader.Load()
}
