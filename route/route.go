package route

import (
	"github.com/gin-gonic/gin"
	controller "morpher-controller/controller/aws_controller"
	service "morpher-controller/service/aws"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	ec2PriceService := service.NewGetEc2AwsPriceService()
	AwsPriceController := controller.NewAwsPriceController(ec2PriceService)

	api := r.Group("/api/v1")
	awsApi := api.Group("/aws")
	{
		awsApi.GET("/ec2", AwsPriceController.GetEc2AwsPrice)
	}

	return r
}
