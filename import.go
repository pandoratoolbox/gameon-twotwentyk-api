package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ImportCelebrities() error {
	var err error
	file, err := os.Open("celebrities.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		fmt.Println(line)

		id, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		name := line[1]

		c_raw := line[2]
		categories := strings.Split(c_raw, "|")

		b_raw := line[3]
		b_split := strings.Split(b_raw, "-")
		if len(b_split) < 3 {
			continue
		}

		year, err := strconv.ParseInt(b_split[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		month, err := strconv.ParseInt(b_split[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		day, err := strconv.ParseInt(b_split[2], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		ccategories := models.Strings(categories)

		store.NewCelebrity(context.TODO(), &models.Celebrity{
			CelebrityData: models.CelebrityData{
				Id:         &id,
				Name:       &name,
				BirthMonth: &month,
				BirthYear:  &year,
				BirthDay:   &day,
				Categories: &ccategories,
			},
		})

		// if err != nil {
		// 	log.Fatal(err)
		// }

		for _, c := range categories {
			new := models.Category{
				CategoryData: models.CategoryData{
					Name: &c,
				},
			}

			store.NewCategory(context.TODO(), &new)
		}

	}

	return nil
}

func ImportArticles() {

}
