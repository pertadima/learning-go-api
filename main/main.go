package main

import (
	ldb "learning-database"
	"learning-database/index/migration/model"
	"learning-database/index/migration/seeder"
)

func main() {
	conn, err := ldb.InitDB(ldb.TableIndex)
	if err != nil {
		panic(err)
	}

	// start migration
	conn.AutoMigrate(&model.Film{}, &model.Genre{}, &model.Actor{}, &model.Customer{}, &model.FilmGenre{}, &model.FilmActor{}, &model.CustomerFilmHistory{}, &model.CustomerFilmPayment{})
	// end migration

	// start seeder
	seeder.Film(conn)
	seeder.Genre(conn)
	seeder.Actor(conn)
	seeder.Customer(conn)
	seeder.FilmGenre(conn)
	seeder.FilmActor(conn)
	seeder.CustomerFilmHistory(conn)
	seeder.CustomerFilmPayment(conn)
}
