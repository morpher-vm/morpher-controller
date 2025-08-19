package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"morpher-controller/controller/agent_controller"
	awsController "morpher-controller/controller/aws_controller"
	infoController "morpher-controller/controller/info_controller"
	agentRepository "morpher-controller/repository/agent_info"
	agentInfoService "morpher-controller/service/agent"
	awsService "morpher-controller/service/aws"
	infoService "morpher-controller/service/info"
	"time"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	corsConfig(r)

	getInfoService := infoService.NewGetServerInfoService()
	serverInfoController := infoController.NewServerInfoController(getInfoService)

	getEc2PriceService := awsService.NewGetEc2AwsPriceService()
	AwsPriceController := awsController.NewAwsPriceController(getEc2PriceService)

	agentInfoRepository := agentRepository.NewAgentInfoRepository()
	agentWriterService := agentInfoService.NewCreateAgentInfoService(agentInfoRepository)
	agentReaderService := agentInfoService.NewGetAgentInfoService(agentInfoRepository)
	agentController := agent_controller.NewAgentInfoController(agentWriterService, agentReaderService)

	r.GET("/info", serverInfoController.GetServerInfo)
	r.POST("/agent/v1/system", agentController.CreateAgentInfo)

	api := r.Group("/api/v1")
	awsApi := api.Group("/aws")
	{
		awsApi.POST("/ec2", AwsPriceController.GetEc2AwsPrice)
	}

	agentInfoApi := api.Group("/agents")
	{
		agentInfoApi.GET("", agentController.GetAgentInfo)
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
