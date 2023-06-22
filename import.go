package main

import (
	"context"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"log"
)

func ImportCelebrity() {

}

func ImportArticles() {

}

func GenerateArticles() {

	for i := 0; i < 1000; i++ {

		exc := ""
		title := ""
		url := ""
		src := int64(1)
		tags := models.Strings{"celebrity1"}

		new := models.Article{
			ArticleData: models.ArticleData{
				Title:           &title,
				Excerpt:         &exc,
				ArticleSourceId: &src,
				Url:             &url,
				Tags:            &tags,
			},
		}

		err := store.NewArticle(context.TODO(), &new)
		if err != nil {
			log.Fatalf("Unable to insert article: %s", *new.Title)
		}
	}
}
