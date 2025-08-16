package aws_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	service "morpher-controller/service/aws"
)

type AwsPriceController struct {
	getEc2AwsPriceService *service.GetEc2AwsPriceService
}

func NewAwsPriceController(
	getEc2AwsPriceService *service.GetEc2AwsPriceService,
) *AwsPriceController {
	return &AwsPriceController{getEc2AwsPriceService: getEc2AwsPriceService}
}

func GetAllAwsPrice(c *gin.Context) {
	c.JSON(200, gin.H{
		"result": "all",
	})
}

func (controller *AwsPriceController) GetEc2AwsPrice(c *gin.Context) {

	var request GetEc2PriceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query := service.GetEc2PriceQuery{
		Region:      request.Region,
		MinVCpu:     request.MinVCpu,
		MinMemoryGB: request.MinMemoryGB,
		OS:          request.OS,
		MaxResults:  request.MaxResults,
	}
	price, err := controller.getEc2AwsPriceService.GetEc2Price(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"result": price,
	})
}

type GetEc2PriceRequest struct {
	Region      string  `json:"region"`
	MinVCpu     int     `json:"minVcpu"`
	MinMemoryGB float64 `json:"minMemoryGB"`
	OS          string  `json:"os"`
	MaxResults  int32   `json:"maxResults"`
}
