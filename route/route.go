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

func corsConfig(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
