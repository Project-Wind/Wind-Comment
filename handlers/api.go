package handlers

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/project-wind/wind-comment/api"
    "github.com/project-wind/wind-comment/models"
)

// GetComments handles requests to get comments.
func GetComments(c *gin.Context) {
    articleID := c.Param("articleID")
    var query api.GetCommentsQuery
    err := c.ShouldBindQuery(&query)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err.Error(),
        })
        return
    }

    var comments models.Comments
    comments, err = comments.Get(articleID, query)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, comments)
}

// PostComment handles requests to post a comment.
func PostComment(c *gin.Context) {
    var comment models.Comment
    err := c.ShouldBindJSON(&comment)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "msg":  err.Error(),
        })
        return
    }

    comment.Date = time.Now().Format("2006-01-02 15:04:05")
    err = comment.Create()
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "msg":  err.Error(),
        })
        return
    }
}
