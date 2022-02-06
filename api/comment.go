package api

type GetCommentsQuery struct {
    ParentID string `form:"parentID"`
    Sort     string `form:"sort"`
    OrderBy  string `form:"orderBy"`
}
