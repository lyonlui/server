package users

import (
	"github.com/name5566/leaf/gate"
)

var agents = make(map[int]*(gate.Agent))
var userCount int = 0

func AddAgent(uid int, agent *(gate.Agent)) {
	agents[uid] = agent
	userCount++

}
func DeleteAgent(uid int) {
	delete(agents, uid)
	userCount--
}

func GetAgent(uid int) *(gate.Agent) {
	return agents[uid]
}

func AgentExit(uid int) bool {

	_, ok := agents[uid]

	return ok
}

func GetAgents() map[int]*(gate.Agent) {
	return agents
}

func GetAgentCounts() int {
	return userCount
}
