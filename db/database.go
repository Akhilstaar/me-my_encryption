package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PuppyDb struct {
	*gorm.DB
}

func InitDB(dbString string) (*PuppyDb, error) {
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to Database: %v", err)
	}
    fmt.Println("Connected Successfully !!")
    db.Close()
	return &PuppyDb{db}, nil
}

func (pdb *PuppyDb) GetTable(table string) *gorm.DB {
	return pdb.DB.Table(table)
}