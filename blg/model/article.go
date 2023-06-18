package model

import "time"

// -- ----------------------------
// --  Table structure for `me_article`
// -- ----------------------------
// DROP TABLE IF EXISTS `me_article`;
// CREATE TABLE `me_article` (
//   `id` int(11) NOT NULL AUTO_INCREMENT,
//   `comment_counts` int(11) DEFAULT NULL,
//   `create_date` datetime DEFAULT NULL,
//   `summary` varchar(100) DEFAULT NULL,
//   `title` varchar(64) DEFAULT NULL,
//   `view_counts` int(11) DEFAULT NULL,
//   `weight` int(11) NOT NULL,
//   `author_id` bigint(20) DEFAULT NULL,
//   `body_id` bigint(20) DEFAULT NULL,
//   `category_id` int(11) DEFAULT NULL,
//   PRIMARY KEY (`id`),
//   KEY `FKndx2m69302cso79y66yxiju4h` (`author_id`),
//   KEY `FKrd11pjsmueckfrh9gs7bc6374` (`body_id`),
//   KEY `FKjrn3ua4xmiulp8raj7m9d2xk6` (`category_id`),
//   CONSTRAINT `FKjrn3ua4xmiulp8raj7m9d2xk6` FOREIGN KEY (`category_id`) REFERENCES `me_category` (`id`),
//   CONSTRAINT `FKndx2m69302cso79y66yxiju4h` FOREIGN KEY (`author_id`) REFERENCES `sys_user` (`id`),
//   CONSTRAINT `FKrd11pjsmueckfrh9gs7bc6374` FOREIGN KEY (`body_id`) REFERENCES `me_article_body` (`id`)
// ) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;

type Article struct {
	Id            int       `gorm:"column:id;type:int;primaryKey;autoIncrement" json:"id,omitempty"`
	CommentCounts int       `gorm:"column:comment_counts;type:int" json:"comment_counts,omitempty"`
	CreateDate    time.Time `gorm:"column:create_date;type:datetime" json:"create_date,omitempty"`
	Summary       string    `gorm:"column:summary;type:varchar;size:255" json:"summary,omitempty"`
	Title         string    `gorm:"column:title;type:varchar;size:255" json:"title,omitempty"`
	ViewCounts    int       `gorm:"column:view_counts;type:int" json:"view_count,omitempty"`
	AuthorId      int64     `gorm:"column:author_id;type:bigint" json:"author_id,omitempty"`
	BodyId        int64     `gorm:"column:body_id;type:bigint" json:"body_id,omitempty"`
	Weight        int       `gorm:"column:weight;type:int" json:"weight,omitempty"`
	CategoryId    int       `gorm:"column:category_id;type:int" json:"category_id,omitempty"`
}

func (a *Article) TableName() string {
	return "me_article"
}

type ArticleCategory struct {
	Id           int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Avatar       string `gorm:"column:avatar;type:varchar"`
	CategoryName string `gorm:"column:categoryname;type:varchar"`
	Description  string `gorm:"column:description;type:varchar"`
}

func (c ArticleCategory) TableName() string {
	return "me_category"
}

type ArticleTag struct {
	Id      int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Avatar  string `gorm:"column:avatar;type:varchar"`
	TagName string `gorm:"column:tagname;type:varchar"`
}

func (t *ArticleTag) TableName() string {
	return "me_tag"
}

type ArticleTagRelation struct {
	ArticleId int `gorm:"column:article_id;type:int"`
	TagId     int `gorm:"column:tag_id;type:int"`
}

func (r *ArticleTagRelation) TableName() string {
	return "me_article_tag"
}

type ArticleBody struct {
	Id          int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Content     string `gorm:"column:content;type:varchar"`
	ContentHtml string `gorm:"column:content_html;type:varchar"`
}

func (b *ArticleBody) TableName() string {
	return "me_article_body"
}
