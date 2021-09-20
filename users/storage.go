package users

import "gorm.io/gorm"

type Storage struct {
	DB *gorm.DB
}

func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		DB: db,
	}
}
