package seeder

import (
	"learning-database/index/migration/model"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	film_actors []*model.FilmActor
)

func FilmActor(db *gorm.DB) {
	numOfPopulatedData := 1000
	numOfInserteddataActors := 2000
	for i := 0; i < numOfInserteddataActors; i++ {
		actorID := uint64(faker.RandomInt(1, numOfInserteddataActors))
		for j := 0; j < numOfPopulatedData; j++ {
			film_actor := &model.FilmActor{
				FilmID:    uint64(faker.RandomInt(1, numOfInserteddataActors)),
				ActorId:   actorID,
				CreatedAt: faker.Date().Backward(time.Hour * 24 * 365 * 10),
			}

			film_actors = append(film_actors, film_actor)
		}
	}

	db.CreateInBatches(film_actors, batchInsertSize)
}
