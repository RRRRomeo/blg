package types

import "time"

type RspArticleCategory struct {
	CategoryId int   `json:"category_id,omitempty"`
	TagId      []int `json:"tag_id,omitempty"`
}

type RspArticleBody struct {
	Content     string `json:"content,omitempty"`
	ContentHtml string `json:"content_html,omitempty"`
}

type RspAuthor struct {
	Name string `json:"name"`
}

type RspArticle struct {
	Id            int            `json:"id,omitempty"`
	Title         string         `json:"title,omitempty"`
	CommentCounts int            `json:"comment_counts,omitempty"`
	ViewCount     int            `json:"view_count,omitempty"`
	Summary       string         `json:"summary,omitempty"`
	Author        RspAuthor      `json:"author,omitempty"`
	Tags          []string       `json:"tags,omitempty"`
	Category      string         `json:"category,omitempty"`
	CreateDate    time.Time      `json:"create_date,omitempty"`
	Body          RspArticleBody `json:"body,omitempty"`
}

type RspArticleCard struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type RspTags struct {
	Id      int    `json:"id,omitempty"`
	TagName string `json:"tagname,omitempty"`
}
