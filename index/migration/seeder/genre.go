package seeder

import (
	"learning-database/index/migration/model"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	genres []*model.Genre
)

func Genre(db *gorm.DB) {

	for i := 0; i < numOfInserteddata; i++ {
		genre := &model.Genre{
			Name:      faker.Name().Name(),
			CreatedAt: faker.Date().Backward(time.Hour * 24 * 365 * 10),
		}

		genres = append(genres, genre)
	}

	db.CreateInBatches(genres, batchInsertSize)
}
