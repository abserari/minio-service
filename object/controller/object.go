package controller

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/yhyddr/minio-service/object/minio"
)

// RegisterRouter
func RegisterRouter(r gin.IRouter, JWT *jwt.GinJWTMiddleware) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	{
		r.POST("/upload", uploadFile)
		r.POST("/dowload", dowloadFile)
	}
}

// upload file
func uploadFile(c *gin.Context) {
	var (
		req struct {
			bucketName  string `json:"bucketname"      binding:"required"`
			objectName  string `json:"objectname"      binding:"required"`
			filePath    string `json:"filepath"        binding:"required"`
			contentType string `json:"contenttype"     binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = minio.UploadFile(req.bucketName, req.objectName, req.filePath, req.contentType)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// dowload file
func dowloadFile(c *gin.Context) {
	var (
		req struct {
			bucketName string `json:"bucketname"      binding:"required"`
			objectName string `json:"objectname"      binding:"required"`
			filePath   string `json:"filepath"        binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = minio.DowloadFile(req.bucketName, req.objectName, req.filePath)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
