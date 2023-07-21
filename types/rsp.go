package types

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

type RspArticleItem struct {
	Id            int       `json:"id,omitempty"`
	Weight        int       `json:"weight,omitempty"`
	Title         string    `json:"title,omitempty"`
	CommentCounts int       `json:"commentCounts,omitempty"`
	ViewCount     int       `json:"viewCounts,omitempty"`
	Summary       string    `json:"summary,omitempty"`
	Author        RspAuthor `json:"author,omitempty"`
	Tags          []RspTags `json:"tags,omitempty"`
	CreateDate    string    `json:"createDate,omitempty"`
	// Body          RspArticleBody `json:"body,omitempty"`
}

type RspCategory struct {
	Id           int    `json:"id,omitempty"`
	CategoryName string `json:"categoryname,omitempty"`
}

type RspArticleDetail struct {
	Id            int            `json:"id,omitempty"`
	Title         string         `json:"title,omitempty"`
	CommentCounts int            `json:"commentCounts,omitempty"`
	ViewCount     int            `json:"viewCounts,omitempty"`
	Summary       string         `json:"summary,omitempty"`
	Author        RspAuthor      `json:"author,omitempty"`
	Tags          []RspTags      `json:"tags,omitempty"`
	Category      RspCategory    `json:"category,omitempty"`
	CreateDate    string         `json:"createDate,omitempty"`
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
