package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	authCmd "backend/internal/auth/application/command"
	authHTTP "backend/internal/auth/interfaces/http"
	authHandler "backend/internal/auth/interfaces/http/handler"
	authPersistence "backend/internal/auth/infrastructure/persistence"
	authSecurity "backend/internal/auth/infrastructure/security"
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

	db := mongoClient.Database(cfg.MongoDBAuth)

	// Wire dependencies (Composition Root).
	userRepo, err := authPersistence.NewMongoUserRepository(db)
	if err != nil {
		log.Fatalf("init user repo: %v", err)
	}
	hasher := authSecurity.NewBcryptHasher()
	tokens := authSecurity.NewJWTService(cfg.JWTSecret, cfg.JWTExpiresHours)

	registerUC := authCmd.NewRegisterUserHandler(userRepo, hasher, tokens)
	loginUC := authCmd.NewLoginUserHandler(userRepo, hasher, tokens)

	h := authHandler.NewAuthHandler(registerUC, loginUC)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "auth", "status": "ok"}) })
	authHTTP.RegisterRoutes(r, h)

	addr := ":" + cfg.AuthServicePort
	log.Printf("auth-service listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
