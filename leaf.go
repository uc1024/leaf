package leaf

import (
	"os"
	"os/signal"

	"github.com/uc1024/leaf/chanrpc"
	"github.com/uc1024/leaf/cluster"
	"github.com/uc1024/leaf/conf"
	"github.com/uc1024/leaf/console"
	"github.com/uc1024/leaf/gate"
	"github.com/uc1024/leaf/log"
	"github.com/uc1024/leaf/module"
	"github.com/uc1024/leaf/network"
)

type Server struct {
	conf          conf.ServerConfig
	gate          *gate.GateModule
	modulesMaster module.IModulesMaster
	clusterMaster cluster.IClusterMaster
	consoleMaster console.IConsoleMaster
}

func NewServer(conf conf.ServerConfig, processor network.Processor, agentChanRPC *chanrpc.Server) *Server {
	s := new(Server)
	s.conf = conf
	s.modulesMaster = module.NewModulesMaster()
	s.clusterMaster = cluster.NewClusterMaster(cluster.ClusterMasterOptions{
		ListenAddr:      conf.ListenAddr,
		ConnAddrs:       conf.ConnAddrs,
		PendingWriteNum: conf.PendingWriteNum,
		ConnectInterval: conf.ConnectInterval,
	})
	s.consoleMaster = console.NewConsoleMaster(console.ConsoleMasterOptions{
		ConsolePort: conf.ConsolePort,
		ProfilePath: conf.ProfilePath,
	})
	s.gate = gate.NewGateModule(gate.GateModuleOptions{
		MaxConnNum:      conf.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.CertFile,
		KeyFile:         conf.KeyFile,
		TCPAddr:         conf.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       processor,
		AgentChanRPC:    agentChanRPC,
	})

	return s
}

func (s *Server) Run(mods ...module.Module) {
	// logger
	if s.conf.LogLevel != "" {
		logger, err := log.New(s.conf.LogLevel, conf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}

	log.Release("Leaf %v starting up", version)

	// module
	mods = append(mods, s.gate)
	for i := 0; i < len(mods); i++ {
		s.modulesMaster.Register(mods[i])
	}
	s.modulesMaster.Init()

	// cluster
	s.clusterMaster.Init()
	// console
	s.consoleMaster.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Release("Leaf closing down (signal: %v)", sig)
	s.clusterMaster.Destroy()
	s.consoleMaster.Destroy()
	s.modulesMaster.Destroy()
}