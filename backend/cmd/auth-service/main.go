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

// auth-service — backed by PostgreSQL (polyglot persistence).
// Rationale: ACID + unique constraints + strong consistency for credentials.
func main() {
	cfg := config.Load()

	ctx := context.Background()
	pool, err := database.NewPostgresPool(ctx, cfg.PostgresAuthDSN)
	if err != nil {
		log.Fatalf("postgres connect: %v", err)
	}
	defer pool.Close()

	if err := authPersistence.Migrate(ctx, pool); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// Composition root — wire dependencies in one place.
	userRepo := authPersistence.NewPostgresUserRepository(pool)
	hasher := authSecurity.NewBcryptHasher()
	tokens := authSecurity.NewJWTService(cfg.JWTSecret, cfg.JWTExpiresHours)

	registerUC := authCmd.NewRegisterUserHandler(userRepo, hasher, tokens)
	loginUC := authCmd.NewLoginUserHandler(userRepo, hasher, tokens)

	h := authHandler.NewAuthHandler(registerUC, loginUC)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"service": "auth", "db": "postgres", "status": "ok"})
	})
	authHTTP.RegisterRoutes(r, h)

	addr := ":" + cfg.AuthServicePort
	log.Printf("auth-service (postgres) listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
