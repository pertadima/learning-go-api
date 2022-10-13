package seeder

import (
	"learning-database/index/migration/model"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	film_genres []*model.FilmGenre
)

func FilmGenre(db *gorm.DB) {
	numOfPopulatedData := 1000
	numOfInserteddataGenres := 2000
	for i := 0; i < numOfInserteddataGenres; i++ {
		genreID := uint64(faker.RandomInt(1, numOfInserteddataGenres))
		for j := 0; j < numOfPopulatedData; j++ {
			film_genre := &model.FilmGenre{
				FilmID:    uint64(faker.RandomInt(1, numOfInserteddataGenres)),
				GenreId:   genreID,
				CreatedAt: faker.Date().Backward(time.Hour * 24 * 365 * 10),
			}

			film_genres = append(film_genres, film_genre)
		}
	}

	db.CreateInBatches(film_genres, batchInsertSize)
}
