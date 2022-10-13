package seeder

import (
	"learning-database/index/migration/model"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	films []*model.Film
)

func Film(db *gorm.DB) {

	for i := 0; i < numOfInserteddata; i++ {
		film := &model.Film{
			Title:         faker.Name().Name(),
			Description:   faker.Lorem().Sentence(10),
			ReleaseYear:   faker.RandomInt(1900, 2020),
			LengthSeconds: uint(faker.RandomInt(3600, 8000)),
			CreatedAt:     faker.Date().Backward(time.Hour * 24 * 365 * 10),
		}

		films = append(films, film)
	}

	db.CreateInBatches(films, batchInsertSize)
}
