package agent

import (
	common "morpher-controller/common/types"
	agentrepository "morpher-controller/repository/agent_info"
)

type CreateAgentInfoService struct {
	agentInfoRepository *agentrepository.AgentInfoRepository
}

func NewCreateAgentInfoService(
	agentInfoRepository *agentrepository.AgentInfoRepository,
) *CreateAgentInfoService {
	return &CreateAgentInfoService{
		agentInfoRepository: agentInfoRepository,
	}
}

func (service *CreateAgentInfoService) Save(agentInfo common.AgentSystemInfo) {
	service.agentInfoRepository.Save(agentInfo)
}
