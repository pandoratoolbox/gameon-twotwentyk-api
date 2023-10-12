package feed

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/lib/pq"

	// "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

type FeedRow struct {
	Id                 int64
	Title              string
	Url                string
	Time               *time.Time
	Text               string
	Image              string
	Category           string
	Publisher          string
	DirectoryUrl       string
	IdentifiedText     []string //celebrity names in article body
	IdentifiedTitle    []string //celebrity names in article title
	IdentifiedTitleGpt []string //trigger descriptions
}

const (
	DB_NAME     = "twotwenty"
	DB_USER     = "postgres"
	DB_PASSWORD = "Vzc|#kENc(RVokACn7xa?Lv8XC$X"
	DB_HOST     = "twotwenty.cluster-cfkyofjse5qs.us-east-1.rds.amazonaws.com"
	DB_PORT     = "5432"
)

func Init() error {
	var err error

	host := DB_HOST
	port := DB_PORT
	user := DB_USER
	name := DB_NAME
	password := url.QueryEscape(DB_PASSWORD)

	q := "postgres://%s:%s@%s:%s/%s"
	connectionString := fmt.Sprintf(q, user, password, host, port, name)

	db, err = sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatalf("Postgres error: %s", err)
	}

	// db, err = pgx.Connect(context.Background(), connectionString)
	// if err != nil {
	// 	log.Fatalf("Postgres error: %s", err)
	// }

	retries := 5

	for r := 0; r < retries; r++ {
		err := db.Ping()
		if err == nil {
			break
		}

		if r == retries {
			log.Fatalf("Unable to establish connection to Postgres: %s", err.Error())
		}

		time.Sleep(10 * time.Second)
	}
	return nil

}

func GetGlobalFeed() ([]FeedRow, error) {
	var feed []FeedRow
	q := `SELECT * FROM news_analysis ORDER BY time DESC LIMIT 100;`
	rows, err := db.Query(q)
	if err != nil {
		return feed, err
	}
	defer rows.Close()

	for rows.Next() {
		var row FeedRow

		err = rows.Scan(&row.Id, &row.Title, &row.Url, &row.Time, &row.Text, &row.Image, &row.Category, &row.Publisher, &row.DirectoryUrl, pq.Array(&row.IdentifiedText), pq.Array(&row.IdentifiedTitle), pq.Array(&row.IdentifiedTitleGpt))
		if err != nil {
			return feed, err
		}
		feed = append(feed, row)
	}

	return feed, nil
}

// func GetPersonalisedFeed(celebrity_names []string) ([]FeedRow, error) {
// 	var feed []FeedRow
// 	q := `SELECT * FROM news_analysis WHERE lower(identified_text::text)::text[] && lower(($1)::text)::text[] OR lower(identified_text::text)::text[] && lower(($1)::text)::text[] ORDER BY time DESC LIMIT 100;` //$1 = array of celeb names
// 	rows, err := db.Query(context.Background(), q, celebrity_names)
// 	if err != nil {
// 		return feed, err
// 	}

// 	for rows.Next() {
// 		var row FeedRow
// 		err = rows.Scan(&row.Id, &row.Title, &row.Url, &row.Time, &row.Text, &row.Image, &row.Category, &row.Publisher, &row.DirectoryUrl, &row.IdentifiedText, &row.IdentifiedTitle, &row.IdentifiedTitleGpt)
// 		if err != nil {
// 			return feed, err
// 		}
// 		feed = append(feed, row)
// 	}

// 	return feed, nil
// }
