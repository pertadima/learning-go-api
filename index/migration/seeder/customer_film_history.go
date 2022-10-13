package seeder

import (
	"learning-database/index/migration/model"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	customer_film_histories []*model.CustomerFilmHistory
)

func CustomerFilmHistory(db *gorm.DB) {
	numOfPopulatedData := 1000
	numOfInserteddataHistories := 2000
	for i := 0; i < numOfInserteddataHistories; i++ {
		customerID := uint64(faker.RandomInt(1, numOfInserteddataHistories))
		for j := 0; j < numOfPopulatedData; j++ {
			customer_film_history := &model.CustomerFilmHistory{
				FilmID:               uint64(faker.RandomInt(1, numOfInserteddataHistories)),
				CustomerId:           customerID,
				WatchDurationSeconds: uint(faker.RandomInt(1000, 3600)),
				CreatedAt:            faker.Date().Backward(time.Hour * 24 * 365 * 10),
			}

			customer_film_histories = append(customer_film_histories, customer_film_history)
		}
	}

	db.CreateInBatches(customer_film_histories, batchInsertSize)
}
