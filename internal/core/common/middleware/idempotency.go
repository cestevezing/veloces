package middleware

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/cestevezing/veloces/internal/core/dto/response"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func IdempotencyMiddleware(redisClient *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idempotencyKey := c.Get("Idempotency-Key")
		if idempotencyKey == "" {
			return response.Error(c, fiber.StatusBadRequest, "Idempotency-Key header is required")
		}

		key := "idempotency:" + idempotencyKey

		status, err := redisClient.Get(context.Background(), key).Result()
		if err == nil {
			if status == "COMPLETED" {
				responseJSON, _ := redisClient.Get(context.Background(), key+":response").Result()
				var responseData map[string]any
				if err := json.Unmarshal([]byte(responseJSON), &responseData); err != nil {
					return response.Error(c, fiber.StatusInternalServerError, "Failed to parse response JSON")
				}
				return response.Success(c, "", responseData)
			} else if status == "IN_PROGRESS" {
				return response.Error(c, fiber.StatusConflict, "Conflict: order in progress")
			}
		}
		redisClient.Set(context.Background(), key, "IN_PROGRESS", 24*time.Hour)
		err = c.Next()
		if err != nil {
			log.Printf("Idempotency: error processing order %s: %v\n", idempotencyKey, err)
			return err
		}
		response := c.Locals("response")
		if response != nil {
			responseJSON, _ := json.Marshal(response)
			redisClient.Set(context.Background(), key+":response", string(responseJSON), 24*time.Hour)
			redisClient.Set(context.Background(), key, "COMPLETED", 24*time.Hour)
			log.Printf("Idempotency: order %s completed\n", idempotencyKey)
		}
		return nil
	}
}
