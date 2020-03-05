package main

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/yhyddr/minio-service/bucket/controller"
)

var (
	// JWTMiddleware should be exported for user authentication.
	JWTMiddleware *jwt.GinJWTMiddleware
)

func main() {
	router := gin.Default()

	controller.RegisterRouter(router.Group("/api/v1/bucket"), JWTMiddleware)

	router.Run(":8000")
}
