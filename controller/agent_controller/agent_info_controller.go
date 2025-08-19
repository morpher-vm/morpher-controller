package agent_controller

import (
	"github.com/gin-gonic/gin"
	common "morpher-controller/common/types"
	"morpher-controller/service/agent"
	"net/http"
)

type AgentInfoController struct {
	agentWriter *agent.CreateAgentInfoService
	agentReader *agent.GetAgentInfoService
}

func NewAgentInfoController(
	createAgentInfoService *agent.CreateAgentInfoService,
	getAgentInfoService *agent.GetAgentInfoService,
) *AgentInfoController {
	return &AgentInfoController{
		agentWriter: createAgentInfoService,
		agentReader: getAgentInfoService,
	}
}

func (controller *AgentInfoController) CreateAgentInfo(c *gin.Context) {
	var requestAgentInfo common.AgentSystemInfo
	if err := c.ShouldBindJSON(&requestAgentInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}

	controller.agentWriter.Save(requestAgentInfo)
}

func (controller *AgentInfoController) GetAgentInfo(c *gin.Context) {
	agentInfos := controller.agentReader.GetAllAgentInfo()

	c.JSON(200, gin.H{
		"result": agentInfos,
	})
}
