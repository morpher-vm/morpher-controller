package agent_info

import (
	common "morpher-controller/common/types"
)

type AgentInfoRepository struct {
	agentInfoMap map[string]common.AgentSystemInfo
}

func NewAgentInfoRepository() *AgentInfoRepository {
	return &AgentInfoRepository{
		agentInfoMap: make(map[string]common.AgentSystemInfo),
	}
}

func (repository *AgentInfoRepository) Save(agentInfo common.AgentSystemInfo) {
	repository.agentInfoMap[agentInfo.ID] = agentInfo
}

func (repository *AgentInfoRepository) GetAll() []common.AgentSystemInfo {
	var agentSystemInfoArr []common.AgentSystemInfo
	for _, value := range repository.agentInfoMap {
		agentSystemInfoArr = append(agentSystemInfoArr, value)
	}

	return agentSystemInfoArr
}
