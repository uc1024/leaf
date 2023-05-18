package cluster

import (
	"github.com/uc1024/leaf/network"
)

type Agent struct {
	conn *network.TCPConn
}

func newAgent(conn *network.TCPConn) network.Agent {
	a := new(Agent)
	a.conn = conn
	return a
}

func (a *Agent) Run() {}

func (a *Agent) OnClose() {}
