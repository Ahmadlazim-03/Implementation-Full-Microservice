package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	placesCmd "backend/internal/places/application/command"
	placesQuery "backend/internal/places/application/query"
	placesCache "backend/internal/places/infrastructure/cache"
	placesPersistence "backend/internal/places/infrastructure/persistence"
	placesHTTP "backend/internal/places/interfaces/http"
	placesHandler "backend/internal/places/interfaces/http/handler"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/middleware"
)

// places-service — backed by MongoDB (geospatial 2dsphere) + Redis cache.
// Rationale: $nearSphere for "places near me", flexible doc schema.
func main() {
	cfg := config.Load()

	ctx := context.Background()
	mongoClient, err := database.NewMongoClient(ctx, cfg.MongoPlacesURI)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(cfg.MongoDBPlaces)

	rdb, err := database.NewRedisClient(ctx, cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Fatalf("redis connect: %v", err)
	}
	defer rdb.Close()

	// Real repositories (Mongo).
	mongoPlaceRepo, err := placesPersistence.NewMongoPlaceRepository(ctx, db)
	if err != nil {
		log.Fatalf("init place repo (2dsphere index): %v", err)
	}
	mongoCategoryRepo := placesPersistence.NewMongoCategoryRepository(db)

	// Decorate with Redis cache — domain code unaware.
	placeRepo := placesCache.NewCachedPlaceRepository(mongoPlaceRepo, rdb, 5*time.Minute)
	categoryRepo := placesCache.NewCachedCategoryRepository(mongoCategoryRepo, rdb, time.Hour)

	createPlaceUC := placesCmd.NewCreatePlaceHandler(placeRepo)
	listPlacesUC := placesQuery.NewListPlacesHandler(placeRepo)
	nearbyUC := placesQuery.NewFindNearbyPlacesHandler(placeRepo)
	listCategoriesUC := placesQuery.NewListCategoriesHandler(categoryRepo)

	placeH := placesHandler.NewPlaceHandler(createPlaceUC, listPlacesUC, nearbyUC)
	categoryH := placesHandler.NewCategoryHandler(listCategoriesUC)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "places", "db": "mongodb+redis", "status": "ok"})
	})
	placesHTTP.RegisterRoutes(r, placeH, categoryH)

	addr := ":" + cfg.PlacesServicePort
	log.Printf("places-service (mongo+redis) listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
