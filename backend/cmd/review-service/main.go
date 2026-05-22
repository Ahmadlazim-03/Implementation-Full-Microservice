package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	reviewCmd "backend/internal/review/application/command"
	reviewQuery "backend/internal/review/application/query"
	reviewPersistence "backend/internal/review/infrastructure/persistence"
	reviewHTTP "backend/internal/review/interfaces/http"
	reviewHandler "backend/internal/review/interfaces/http/handler"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/middleware"
)

// review-service — backed by MongoDB.
// Rationale: append-mostly writes, document model fits rating+comment naturally.
func main() {
	cfg := config.Load()

	ctx := context.Background()
	mongoClient, err := database.NewMongoClient(ctx, cfg.MongoReviewURI)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(cfg.MongoDBReview)

	reviewRepo := reviewPersistence.NewMongoReviewRepository(db)

	createUC := reviewCmd.NewCreateReviewHandler(reviewRepo)
	listUC := reviewQuery.NewListReviewsByPlaceHandler(reviewRepo)

	h := reviewHandler.NewReviewHandler(createUC, listUC)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "review", "db": "mongodb", "status": "ok"})
	})
	reviewHTTP.RegisterRoutes(r, h)

	addr := ":" + cfg.ReviewServicePort
	log.Printf("review-service (mongo) listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
