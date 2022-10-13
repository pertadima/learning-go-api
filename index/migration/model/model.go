package model

import "time"

type Film struct {
	ID            uint64    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Title         string    `gorm:"size:255"`
	Description   string    `gorm:"type:text"`
	ReleaseYear   int       `gorm:"type:int(4)"`
	LengthSeconds uint      `gorm:"type:int(11)"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type Genre struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name      string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type Actor struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	FirstName string    `gorm:"size:255"`
	LastName  string    `gorm:"size:255"`
	Birthday  time.Time `gorm:"type:date;"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type Customer struct {
	ID        uint64    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	FirstName string    `gorm:"size:255"`
	LastName  string    `gorm:"size:255"`
	Email     string    `gorm:"size:255"`
	IsActive  string    `gorm:"type:enum('0','1')"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type FilmGenre struct {
	ID        uint64 `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	FilmID    uint64
	GenreId   uint64
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type FilmActor struct {
	ID        uint64 `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	FilmID    uint64
	ActorId   uint64
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type CustomerFilmHistory struct {
	ID                   uint64 `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	FilmID               uint64
	CustomerId           uint64
	WatchDurationSeconds uint      `gorm:"type:int(11)"`
	CreatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time `gorm:"type:timestamp;"`
}

type CustomerFilmPayment struct {
	ID          uint64 `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	FilmID      uint64
	CustomerId  uint64
	Price       float64   `gorm:"type:decimal(20,2)"`
	PaymentDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	PaymentType string    `gorm:"type:enum('direct','thirdparty')"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}
