package handlers

import (
	"fmt"
	"gameon-twotwentyk-api/feed"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "article_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetArticle(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.Article{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewArticle(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := models.Article{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateArticle(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "article_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteArticle(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func Truncate(text string, width int) (string, error) {
	if width < 0 {
		return "", fmt.Errorf("invalid width size")
	}

	r := []rune(text)
	trunc := r[:width]
	return string(trunc) + "...", nil
}

func SearchArticles(w http.ResponseWriter, r *http.Request) {
	var err error
	var out []models.Article

	// q := chi.URLParam(r, "q")
	// if q != "" {
	// 	out, err = store.SearchArticles(ctx, q)
	// 	if err != nil {
	// 		ServeError(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// } else {
	// 	out, err = store.ListArticles(ctx)
	// 	if err != nil {
	// 		ServeError(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	rows, err := feed.GetGlobalFeed()
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, row := range rows {
		// max := len(row.Text)
		// if max > 100 {
		// 	max = 100
		// }

		// excerpt, err := Truncate(row.Text, 200)
		// if err != nil {
		// 	ServeError(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		excerpt := row.Text

		out = append(out, models.Article{
			ArticleData: models.ArticleData{
				Id:           &row.Id,
				Title:        &row.Title,
				Url:          &row.Url,
				CreatedAt:    row.Time,
				ThumbnailSrc: &row.Image,
				Excerpt:      &excerpt,
			},
		})
	}

	ServeJSON(w, out)
}

func GetArticlesPersonalised(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//get nft identity cards
	//get list of names from identity cards
	//get nft prediction cards
	//get list of events and names from nft prediction cards

	// q := ""

	out, err := store.ListArticles(ctx)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, out)
}
