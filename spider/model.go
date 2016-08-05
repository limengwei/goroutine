package spider

import (
	"html/template"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// 通用的内容
type Content struct {
	Type         int
	Title        string
	Markdown     string
	Html         template.HTML
	CommentCount int
	Hits         int // 点击数量
	CreatedAt    time.Time
	CreatedBy    bson.ObjectId
	UpdatedAt    time.Time
	UpdatedBy    string
}

// 文章
type Article struct {
	Id bson.ObjectId `bson:"_id"`
	Content
	OriginalSource string
	OriginalUrl    string
	CategoryId     bson.ObjectId
}

// 文章分类
type ArticleCategory struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string
}

//开发者头条
//伯乐在线
//极客头条
//技术头条
//稀土掘金
//app开发日报
//深度开源
//推酷
