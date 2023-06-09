package types

type ReqUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// Id         int       `gorm:"column:id"`
// Title      string    `gorm:"column:title"`
// Content    string    `gorm:"column:content"`
// Is_public  int8      `gorm:"column:is_public"`
// Author_id  int64     `gorm:"column:author_id"`
// Created_at time.Time `gorm:"column:create_at"`
// Updated_at time.Time `gorm:"column:update_at"`
type ReqCreatePost struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsPublic int8   `json:"ispublic"`
}

type ReqUpdatePost struct {
	PostId int `json:"postid"`
	ReqCreatePost
}
