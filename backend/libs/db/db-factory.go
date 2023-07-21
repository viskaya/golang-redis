package db

import "gorm.io/gorm"

type DBFactory struct {
	DB    *gorm.DB
	Cache CacheDB
}

func NewDBFactory() *DBFactory {
	return &DBFactory{
		DB:    InitDB(),
		Cache: InitCache(),
	}
}
