package seeder

import (
	"learning-database/index/migration/model"
	"strconv"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	customers []*model.Customer
)

func Customer(db *gorm.DB) {

	for i := 0; i < numOfInserteddata; i++ {
		customer := &model.Customer{
			FirstName: faker.Name().FirstName(),
			LastName:  faker.Name().LastName(),
			Email:     faker.Internet().Email(),
			IsActive:  strconv.Itoa(faker.RandomInt(0, 1)),
			CreatedAt: faker.Date().Backward(time.Hour * 24 * 365 * 10),
		}

		customers = append(customers, customer)
	}

	db.CreateInBatches(customers, batchInsertSize)
}
