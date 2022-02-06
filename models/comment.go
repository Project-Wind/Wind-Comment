package models

import (
    "fmt"
    "strconv"

    "github.com/project-wind/wind-comment/api"
)

type Comment struct {
    ID        int    `json:"id"`
    ParentID  int    `json:"parentID"`
    ArticleID string `json:"articleID"`
    Avatar    string `json:"avatar"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Content   string `json:"content"`
    Date      string `json:"date"`
}

type Comments []Comment

// Get reads comments from the database and returns them.
func (comments Comments) Get(articleID string, query api.GetCommentsQuery) (Comments, error) {
    db := DB

    if query.ParentID != "" {
        parentID, err := strconv.Atoi(query.ParentID)
        if err != nil {
            return nil, err
        }
        db = db.Where("article_id = ? AND parent_id = ?", articleID, parentID)
    } else {
        db = db.Where("article_id = ?", articleID)
    }

    if query.OrderBy != "" && query.Sort != "" {
        db = db.Order(fmt.Sprintf("%v %v", query.Sort, query.OrderBy))
    }

    err := db.Find(&comments).Error
    if err != nil {
        return nil, err
    }
    return comments, nil
}

// Create inserts a comment into the database.
func (comment Comment) Create() error {
    err := DB.Create(&comment).Error
    if err != nil {
        return err
    }
    return nil
}
