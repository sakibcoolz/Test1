package service

import (
	"Test1/dao"
	"Test1/service/fileoperations"

	"gorm.io/gorm"
)

type Services struct {
	FileOperation fileoperations.IFileService
}

func GiveService(Db *gorm.DB) Services {
	return Services{
		FileOperation: fileoperations.FileService{
			Filestore: dao.DatabaseConnection(Db),
		},
	}
}
