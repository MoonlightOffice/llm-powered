package repository

import "example/model/article"

type IArticleRepository interface {
	FindById(id string) (*article.Article, error)
	Save(article article.Article) error
}
