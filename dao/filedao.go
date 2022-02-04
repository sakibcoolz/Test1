package dao

import (
	"Test1/model"
	"log"

	"gorm.io/gorm"
)

type FileDao struct {
	Db *gorm.DB
}

func DatabaseConnection(db *gorm.DB) FileDao {
	return FileDao{Db: db}
}

type IFileDao interface {
	StringOperations(sql string) (selectData []model.SelectData)
}

func (F FileDao) StringOperations(sql string) (selectData []model.SelectData) {
	if db := F.Db.Raw(sql).Scan(&selectData); db.Error != nil {
		log.Fatal(db.Error)
	}

	return selectData
}
