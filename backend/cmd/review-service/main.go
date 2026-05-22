package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	reviewCmd "backend/internal/review/application/command"
	reviewQuery "backend/internal/review/application/query"
	reviewGrpc "backend/internal/review/infrastructure/grpc"
	reviewPersistence "backend/internal/review/infrastructure/persistence"
	reviewHTTP "backend/internal/review/interfaces/http"
	reviewHandler "backend/internal/review/interfaces/http/handler"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/middleware"
)

// review-service — MongoDB untuk reviews.
// Komunikasi ke places-service via gRPC (internal protobuf) untuk validasi place_id.
func main() {
	cfg := config.Load()

	ctx := context.Background()
	mongoClient, err := database.NewMongoClient(ctx, cfg.MongoReviewURI)
	if err != nil {
		log.Fatalf("mongo connect: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	db := mongoClient.Database(cfg.MongoDBReview)

	// gRPC client ke places-service — Anti-Corruption Layer.
	verifier, err := reviewGrpc.NewGrpcPlaceVerifier(cfg.PlacesGrpcAddr)
	if err != nil {
		log.Fatalf("grpc place verifier: %v", err)
	}
	defer verifier.Close()
	log.Printf("review-service connected to places gRPC at %s", cfg.PlacesGrpcAddr)

	reviewRepo := reviewPersistence.NewMongoReviewRepository(db)

	createUC := reviewCmd.NewCreateReviewHandler(reviewRepo, verifier)
	listUC := reviewQuery.NewListReviewsByPlaceHandler(reviewRepo)

	h := reviewHandler.NewReviewHandler(createUC, listUC)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "review", "db": "mongodb", "status": "ok"})
	})
	reviewHTTP.RegisterRoutes(r, h)

	addr := ":" + cfg.ReviewServicePort
	log.Printf("review-service HTTP listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
