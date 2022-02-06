package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/project-wind/wind-comment/handlers"
)

// Init initializes routers of Gin.
func Init(r *gin.Engine) {
    r.GET("/api/v0/comments/:articleID", handlers.GetComments)
    r.POST("/api/v0/comments", handlers.PostComment)
}
