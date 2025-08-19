package agent_info

import (
	common "morpher-controller/common/types"
	"sync"
)

type AgentInfoRepository struct {
	agentInfoMap map[string]common.AgentSystemInfo
	mutex        sync.RWMutex
}

func NewAgentInfoRepository() *AgentInfoRepository {
	return &AgentInfoRepository{
		agentInfoMap: make(map[string]common.AgentSystemInfo),
	}
}

func (repository *AgentInfoRepository) Save(agentInfo common.AgentSystemInfo) {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()

	repository.agentInfoMap[agentInfo.ID] = agentInfo
}

func (repository *AgentInfoRepository) GetAll() []common.AgentSystemInfo {
	repository.mutex.Lock()
	defer repository.mutex.Unlock()

	var agentSystemInfoArr []common.AgentSystemInfo
	for _, value := range repository.agentInfoMap {
		agentSystemInfoArr = append(agentSystemInfoArr, value)
	}

	return agentSystemInfoArr
}
