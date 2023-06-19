package types

type ReqUser struct {
	Account  string `json:"account,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Password string `json:"password,omitempty"`
}

type ReqCreatePost struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsPublic int8   `json:"ispublic"`
}

type ReqUpdatePost struct {
	PostId int `json:"postid"`
	ReqCreatePost
}

type RespGetCurrentUser struct {
	Account string `json:"account,omitempty"`
	Nciname string `json:"nciname,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Id      int64  `json:"id,omitempty"`
}

type ReqArticleBody struct {
	Content     string `json:"content,omitempty"`
	ContentHtml string `json:"content_html,omitempty"`
}

type ReqArticle struct {
	Id          int      `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Summary     string   `json:"summary,omitempty"`
	Category    string   `json:"category,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Content     string   `json:"content,omitempty"`
	ContentHtml string   `json:"content_html,omitempty"`
}

type ReqComment struct {
	ArticleId int    `json:"article_id,omitempty"`
	AuthorId  int    `json:"author_id,omitempty"`
	ParentId  int    `json:"parent_id,omitempty"`
	ToUid     int64  `json:"to_uid,omitempty"`
	Content   string `json:"content,omitempty"`
}
