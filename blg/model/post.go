package model

import "time"

//   CREATE TABLE `t_posts` (
//   id INT(11) PRIMARY KEY AUTO_INCREMENT,
//   title VARCHAR(100) NOT NULL,
//   content TEXT NOT NULL,
//   is_public TINYINT(1) NOT NULL,
//   author_id BIGINT NOT NULL,
//   created_at DATETIME NOT NULL,
//   updated_at DATETIME NOT NULL,
//   FOREIGN KEY (author_id) REFERENCES t_user(id)
// );

type Post struct {
	Id         int       `gorm:"column:id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	Is_public  int8      `gorm:"column:is_public"`
	Author_id  int64     `gorm:"column:author_id"`
	Created_at time.Time `gorm:"column:created_at"`
	Updated_at time.Time `gorm:"column:updated_at"`
}

func (p *Post) TableName() string {
	return "t_posts"
}
