package main

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	bucket "github.com/yhyddr/minio-service/bucket/controller"
	object "github.com/yhyddr/minio-service/object/controller"
)

var (
	// JWTMiddleware should be exported for user authentication.
	JWTMiddleware *jwt.GinJWTMiddleware
)

func main() {
	router := gin.Default()

	bucket.RegisterRouter(router.Group("/api/v1/bucket"), JWTMiddleware)

	object.RegisterRouter(router.Group("/api/v1/object"), JWTMiddleware)

	router.Run(":8000")
}
