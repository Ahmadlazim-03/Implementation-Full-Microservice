package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	placesCmd "backend/internal/places/application/command"
	placesQuery "backend/internal/places/application/query"
	placesPersistence "backend/internal/places/infrastructure/persistence"
	placesHTTP "backend/internal/places/interfaces/http"
	placesHandler "backend/internal/places/interfaces/http/handler"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/middleware"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()
	mongoClient, err := database.NewMongoClient(ctx, cfg.MongoURI)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(cfg.MongoDBPlaces)

	placeRepo := placesPersistence.NewMongoPlaceRepository(db)
	categoryRepo := placesPersistence.NewMongoCategoryRepository(db)

	createPlaceUC := placesCmd.NewCreatePlaceHandler(placeRepo)
	listPlacesUC := placesQuery.NewListPlacesHandler(placeRepo)
	listCategoriesUC := placesQuery.NewListCategoriesHandler(categoryRepo)

	placeH := placesHandler.NewPlaceHandler(createPlaceUC, listPlacesUC)
	categoryH := placesHandler.NewCategoryHandler(listCategoriesUC)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "places", "status": "ok"}) })
	placesHTTP.RegisterRoutes(r, placeH, categoryH)

	addr := ":" + cfg.PlacesServicePort
	log.Printf("places-service listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
