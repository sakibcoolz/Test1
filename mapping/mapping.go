package mapping

import (
	"Test1/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UrlMapping(router *gin.Engine) *gin.Engine {

	controller := controller.GiveControl(func() *gorm.DB {
		dsn := "root:@tcp(127.0.0.1:3306)/test"

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		return db
	}())

	router.GET("/wordcounter", controller.Controller.GetTest)

	return router
}
