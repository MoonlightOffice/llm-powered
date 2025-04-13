package repositoryImpl

import (
	"errors"
	"example/client/repository"
	"example/model/article"
	"example/model/cterr"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository() repository.IArticleRepository {
	return &articleRepository{
		db: InitSQLite(),
	}
}

func (ar *articleRepository) FindById(id string) (*article.Article, error) {
	var a article.Article
	result := ar.db.First(&a, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, cterr.ErrorNotExisted
		}
		return nil, result.Error
	}
	return &a, nil
}

func (ar *articleRepository) Save(a article.Article) error {
	result := ar.db.Save(&a)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
