package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	placesv1 "backend/gen/proto/kampusmap/places/v1"
	placesCmd "backend/internal/places/application/command"
	placesQuery "backend/internal/places/application/query"
	placesCache "backend/internal/places/infrastructure/cache"
	placesPersistence "backend/internal/places/infrastructure/persistence"
	placesGrpc "backend/internal/places/interfaces/grpc"
	placesHTTP "backend/internal/places/interfaces/http"
	placesHandler "backend/internal/places/interfaces/http/handler"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/middleware"
)

// places-service — MongoDB (geo) + Redis (cache).
// Mengekspos DUA protocol:
//   - HTTP/JSON  :8082  (public, lewat gateway)
//   - gRPC/proto :9082  (internal, service-to-service)
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

	// Repositories.
	mongoPlaceRepo, err := placesPersistence.NewMongoPlaceRepository(ctx, db)
	if err != nil {
		log.Fatalf("init place repo (2dsphere index): %v", err)
	}
	mongoCategoryRepo := placesPersistence.NewMongoCategoryRepository(db)

	placeRepo := placesCache.NewCachedPlaceRepository(mongoPlaceRepo, rdb, 5*time.Minute)
	categoryRepo := placesCache.NewCachedCategoryRepository(mongoCategoryRepo, rdb, time.Hour)

	// Use cases.
	createPlaceUC := placesCmd.NewCreatePlaceHandler(placeRepo)
	listPlacesUC := placesQuery.NewListPlacesHandler(placeRepo)
	nearbyUC := placesQuery.NewFindNearbyPlacesHandler(placeRepo)
	getPlaceUC := placesQuery.NewGetPlaceHandler(placeRepo)
	listCategoriesUC := placesQuery.NewListCategoriesHandler(categoryRepo)

	// ===== Run gRPC server in background (internal protobuf API) =====
	go func() {
		lis, err := net.Listen("tcp", ":"+cfg.PlacesGrpcPort)
		if err != nil {
			log.Fatalf("grpc listen: %v", err)
		}
		grpcSrv := grpc.NewServer()
		placesv1.RegisterPlacesServiceServer(grpcSrv, placesGrpc.NewPlacesServer(placeRepo, getPlaceUC))
		log.Printf("places-service gRPC listening on :%s (internal/protobuf)", cfg.PlacesGrpcPort)
		if err := grpcSrv.Serve(lis); err != nil {
			log.Fatalf("grpc serve: %v", err)
		}
	}()

	// ===== HTTP server (public REST/JSON) =====
	placeH := placesHandler.NewPlaceHandler(createPlaceUC, listPlacesUC, nearbyUC)
	categoryH := placesHandler.NewCategoryHandler(listCategoriesUC)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "places", "db": "mongodb+redis", "status": "ok"})
	})
	placesHTTP.RegisterRoutes(r, placeH, categoryH)

	addr := ":" + cfg.PlacesServicePort
	log.Printf("places-service HTTP listening on %s (public/JSON)", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
