package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"

	"backend/pkg/config"
	"backend/pkg/middleware"
)

func reverseProxy(target string) gin.HandlerFunc {
	u, err := url.Parse(target)
	if err != nil {
		log.Fatalf("invalid target %s: %v", target, err)
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	cfg := config.Load()

	r := gin.Default()
	r.Use(middleware.CORS())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "gateway", "status": "ok"})
	})

	// Route per bounded context.
	r.Any("/api/auth/*proxyPath", reverseProxy(cfg.AuthServiceURL))
	r.Any("/api/places", reverseProxy(cfg.PlacesServiceURL))
	r.Any("/api/places/*proxyPath", reverseProxy(cfg.PlacesServiceURL))
	r.Any("/api/categories", reverseProxy(cfg.PlacesServiceURL))
	r.Any("/api/categories/*proxyPath", reverseProxy(cfg.PlacesServiceURL))
	r.Any("/api/reviews", reverseProxy(cfg.ReviewServiceURL))
	r.Any("/api/reviews/*proxyPath", reverseProxy(cfg.ReviewServiceURL))

	addr := ":" + cfg.GatewayPort
	log.Printf("api-gateway listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
