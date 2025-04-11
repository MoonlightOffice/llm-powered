package article

import (
	"example/util"
	"time"
)

type Article struct {
	ArticleId   string
	Title       string
	Content     string
	PublishedAt int64 // UnixMilli time
}

func NewArticle(title string, content string) Article {
	return Article{
		ArticleId:   util.GenId("article"),
		Title:       title,
		Content:     content,
		PublishedAt: time.Now().UnixMilli(),
	}
}
