package controller

import (
	"Test1/controller/filecontroller"
	"Test1/service"

	"gorm.io/gorm"
)

type Control struct {
	Controller filecontroller.IController
}

func GiveControl(db *gorm.DB) Control {
	return Control{
		Controller: filecontroller.Controller{
			Service: service.GiveService(db),
		},
	}
}
