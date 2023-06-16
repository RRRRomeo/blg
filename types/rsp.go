package types

import "time"

type RspArticleCategory struct {
	CategoryId int   `json:"category_id,omitempty"`
	TagId      []int `json:"tag_id,omitempty"`
}

//	article: {
//	  id: '',
//	  title: '',
//	  commentCounts: 0,
//	  viewCounts: 0,
//	  summary: '',
//	  author: {},
//	  tags: [],
//	  category:{},
//	  createDate: '',
//	  editor: {
//	    value: '',
//	    toolbarsFlag: false,
//	    subfield: false,
//	    defaultOpen: 'preview'
//	  }
//	},
type RspArticleBody struct {
	Content     string `json:"content,omitempty"`
	ContentHtml string `json:"content_html,omitempty"`
}

type RspArticle struct {
	Id            int            `json:"id,omitempty"`
	Title         string         `json:"title,omitempty"`
	CommentCounts int            `json:"comment_counts,omitempty"`
	ViewCount     int            `json:"view_count,omitempty"`
	Summary       string         `json:"summary,omitempty"`
	Author        string         `json:"author,omitempty"`
	Tags          []string       `json:"tags,omitempty"`
	Category      string         `json:"category,omitempty"`
	CreateDate    time.Time      `json:"create_date,omitempty"`
	Body          RspArticleBody `json:"body,omitempty"`
}
