package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewArticle(ctx context.Context, data *models.Article) error {
	err := graphql.NewArticle(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetArticle(ctx context.Context, id int64) (models.Article, error) {
	data, err := graphql.GetArticle(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteArticle(ctx context.Context, id int64) error {
	err := graphql.DeleteArticle(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateArticle(ctx context.Context, data models.Article) error {
	err := graphql.UpdateArticle(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetArticlesByTagAndDescription(ctx context.Context, q string) ([]models.Article, error) {
	var out []models.Article
	return out, nil
}

func GetArticlesPersonalised(ctx context.Context) ([]models.Article, error) {
	var out []models.Article
	return out, nil
}

func GetArticlesNewest(ctx context.Context) ([]models.Article, error) {
	var out []models.Article
	return out, nil
}
