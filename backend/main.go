package main

import (
	"log"

	"crypto/tls"

	_ "github.com/lib/pq"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/yourusername/URLShorten/config"
	"github.com/yourusername/URLShorten/internal/handler"
	"github.com/yourusername/URLShorten/internal/repository"
	"github.com/yourusername/URLShorten/internal/services"
)

func main() {
	cfg := config.Load()
	db, err := sqlx.Connect("postgres", cfg.GetDBConnectionString())

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:      cfg.GetRedisAddress(),
		Password:  cfg.RedisPassword,
		TLSConfig: &tls.Config{},
	})
	repo := repository.New(db, rdb)
	service := services.New(repo)
	routeHandler := handler.New(service)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))
	r.POST("/api/shorten", routeHandler.ShortenURL)
	r.GET("/:code", routeHandler.GetURLHandler)
	r.GET("/api/stats/:code", routeHandler.GetStatsHandler)

	r.Run(":" + cfg.ServerPort)

}
