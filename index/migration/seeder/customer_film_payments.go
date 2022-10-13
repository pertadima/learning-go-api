package seeder

import (
	"learning-database/index/migration/model"
	"math/rand"
	"time"

	"gorm.io/gorm"
	"syreclabs.com/go/faker"
)

var (
	customer_film_payments []*model.CustomerFilmPayment
)

func CustomerFilmPayment(db *gorm.DB) {
	numOfPopulatedData := 1000
	numOfInserteddataPayments := 2000

	for i := 0; i < numOfInserteddataPayments; i++ {
		customerID := uint64(faker.RandomInt(1, numOfInserteddataPayments))
		for j := 0; j < numOfPopulatedData; j++ {
			paymentType := []string{"direct", "thirdparty"}
			customer_film_payment := &model.CustomerFilmPayment{
				FilmID:      uint64(faker.RandomInt(1, numOfInserteddataPayments)),
				CustomerId:  customerID,
				Price:       float64(faker.RandomInt(10000, 1000000)),
				PaymentDate: faker.Date().Backward(time.Hour * 24 * 365 * 10),
				PaymentType: paymentType[rand.Intn(len(paymentType))],
				CreatedAt:   faker.Date().Backward(time.Hour * 24 * 365 * 10),
			}

			customer_film_payments = append(customer_film_payments, customer_film_payment)
		}
	}

	db.CreateInBatches(customer_film_payments, batchInsertSize)
}
