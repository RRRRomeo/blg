package model

import "time"

type Comment struct {
	Id         int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id,omitempty"`
	Content    string    `gorm:"column:content;type:varchar;size:255" json:"content,omitempty"`
	CreateDate time.Time `gorm:"column:create_date;type:datetime" json:"create_date,omitempty"`
	ArticleId  int       `gorm:"column:article_id;type:int" json:"article_id,omitempty"`
	AuthorId   int       `gorm:"column:author_id;type:int" json:"author_id,omitempty"`
	ParentId   int       `gorm:"column:parent_id;type:int" json:"parent_id,omitempty"`
	ToUid      int64     `gorm:"column:to_uid;type:bigint" json:"to_uid,omitempty"`
	Level      byte      `gorm:"column:level;type:varchar;size:1" json:"level,omitempty"`
}

func (c *Comment) TableName() string {
	return "me_comment"
}
