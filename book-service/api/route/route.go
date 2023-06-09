package route

import (
	"time"

	"github.com/temryakov/go-backend-book-app/book-service/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(config *bootstrap.Config, db *gorm.DB, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("/api")
	NewBookRouter(config, db, timeout, publicRouter)
}
