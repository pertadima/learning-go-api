package seeder

import (
	"learning-database/index/migration/model"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	actors []*model.Actor
)

func Actor(db *gorm.DB) {

	for i := 0; i < numOfInserteddata; i++ {
		actor := &model.Actor{
			FirstName: faker.Name().FirstName(),
			LastName:  faker.Name().LastName(),
			Birthday:  faker.Date().Birthday(9, 80),
			CreatedAt: faker.Date().Backward(time.Hour * 24 * 365 * 10),
		}

		actors = append(actors, actor)
	}

	db.CreateInBatches(actors, batchInsertSize)
}
