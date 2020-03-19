package controller

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/yhyddr/minio-service/bucket/minio"
)

// RegisterRouter
func RegisterRouter(r gin.IRouter, JWT *jwt.GinJWTMiddleware) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	{
		r.POST("/create", makeBucket)
	}
}

// make a new bucket
func makeBucket(c *gin.Context) {
	var (
		req struct {
			bucketName string `json:"bucketname"         binding:"required"`
			location   string `json:"location"           binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = minio.MakeBucket(req.bucketName, req.location)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
