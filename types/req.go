package types

type ReqUser struct {
	Account  string `json:"account,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Password string `json:"password,omitempty"`
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

// commit('SET_ACCOUNT', data.data.account)
// commit('SET_NAME', data.data.nickname)
// commit('SET_AVATAR', data.data.avatar)
// commit('SET_ID', data.data.id)
type RespGetCurrentUser struct {
	Account string `json:"account,omitempty"`
	Nciname string `json:"nciname,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Id      int64  `json:"id,omitempty"`
}

// title: this.articleForm.title,
// summary: this.articleForm.summary,
// category: this.articleForm.category,
// tags: tags,
//
//	body: {
//	  content: this.articleForm.editor.value,
//	  contentHtml: this.articleForm.editor.ref.d_render
//	}

type ReqArticleBody struct {
	Content     string `json:"content,omitempty"`
	ContentHtml string `json:"content_html,omitempty"`
}

type ReqArticle struct {
	Id       int            `json:"id,omitempty"`
	Title    string         `json:"title,omitempty"`
	Summary  string         `json:"summary,omitempty"`
	Category string         `json:"category,omitempty"`
	Tags     []string       `json:"tags,omitempty"`
	Body     ReqArticleBody `json:"body,omitempty"`
}
