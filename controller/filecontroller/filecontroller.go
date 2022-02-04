package filecontroller

import (
	"Test1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service service.Services
}

type IController interface {
	GetTest(c *gin.Context)
}

func (s Controller) GetTest(c *gin.Context) {
	data := c.Query("message")
	sk := s.Service.FileOperation.FileServiceOperation(data)
	c.JSON(http.StatusOK, sk)
}
