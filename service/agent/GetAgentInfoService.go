package agent

import (
	common "morpher-controller/common/types"
	agentrepository "morpher-controller/repository/agent_info"
)

type GetAgentInfoService struct {
	agentInfoRepository *agentrepository.AgentInfoRepository
}

func NewGetAgentInfoService(
	agentInfoRepository *agentrepository.AgentInfoRepository,
) *GetAgentInfoService {
	return &GetAgentInfoService{
		agentInfoRepository: agentInfoRepository,
	}
}

func (service *GetAgentInfoService) GetAllAgentInfo() []common.AgentSystemInfo {
	return service.agentInfoRepository.GetAll()
}
