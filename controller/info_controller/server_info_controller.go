package info_controller

import (
	"github.com/gin-gonic/gin"
	service "morpher-controller/service/info"
)

type ServerInfoController struct {
	getServerInfoService *service.GetServerInfoService
}

func NewServerInfoController(
	getServerInfoService *service.GetServerInfoService,
) *ServerInfoController {
	return &ServerInfoController{getServerInfoService: getServerInfoService}
}

func (controller *ServerInfoController) GetServerInfo(c *gin.Context) {
	response := controller.getServerInfoService.GetServerInfo()

	c.JSON(200, gin.H{
		"result": response,
	})
}
