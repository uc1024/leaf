package cluster

import (
	"math"
	"time"

	"github.com/uc1024/leaf/network"
)

type IClusterMaster interface {
	Init()
	Destroy()
}

type ClusterMasterOptions struct {
	ListenAddr string   `json:"listen_addr" yaml:"listenAddr"`
	ConnAddrs  []string `json:"conn_addrs" yaml:"connAddrs"`

	PendingWriteNum int `json:"pending_write_num" yaml:"pendingWriteNum"`

	ConnectInterval time.Duration `json:"connect_interval" yaml:"connectInterval"`
}

type ClusterMaster struct {
	options ClusterMasterOptions
	server  *network.TCPServer
	clients []*network.TCPClient
}

func NewClusterMaster(options ClusterMasterOptions) IClusterMaster {
	return &ClusterMaster{
		options: options,
	}
}

func (master *ClusterMaster) Init() {
	if master.options.ListenAddr != "" {
		master.server = new(network.TCPServer)
		master.server.Addr = master.options.ListenAddr
		master.server.MaxConnNum = int(math.MaxInt32)
		master.server.PendingWriteNum = master.options.PendingWriteNum
		master.server.LenMsgLen = 4
		master.server.MaxMsgLen = math.MaxUint32
		master.server.NewAgent = newAgent

		master.server.Start()
	}

	for _, addr := range master.options.ConnAddrs {
		client := new(network.TCPClient)
		client.Addr = addr
		client.ConnNum = 1
		client.ConnectInterval = master.options.ConnectInterval
		client.PendingWriteNum = master.options.PendingWriteNum
		client.LenMsgLen = 4
		client.MaxMsgLen = math.MaxUint32
		client.NewAgent = newAgent

		client.Start()
		master.clients = append(master.clients, client)
	}
}

func (master *ClusterMaster) Destroy() {
	if master.server != nil {
		master.server.Close()
	}

	for _, client := range master.clients {
		client.Close()
	}
}
