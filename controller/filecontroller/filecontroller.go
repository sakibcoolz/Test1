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
	GetCountDB(c *gin.Context)
	GetCount(c *gin.Context)
}

func (s Controller) GetCountDB(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	sk := s.Service.FileOperation.ServiceOperation(reqBody)
	c.JSON(http.StatusOK, sk)
}

func (s Controller) GetCount(c *gin.Context) {
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	sk := s.Service.FileOperation.StringServiceOperations(reqBody)
	c.JSON(http.StatusOK, sk)
}
