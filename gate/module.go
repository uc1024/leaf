package gate

import (
	"time"

	"github.com/uc1024/leaf/chanrpc"
	"github.com/uc1024/leaf/network"
)

type GateModuleOptions struct {
	MaxConnNum      int
	PendingWriteNum int
	MaxMsgLen       uint32
	WSAddr          string
	HTTPTimeout     time.Duration
	CertFile        string
	KeyFile         string
	TCPAddr         string
	LenMsgLen       int
	LittleEndian    bool
	Processor       network.Processor
	AgentChanRPC    *chanrpc.Server
}

type GateModule struct {
	options GateModuleOptions
	*Gate
}

func NewGateModule(options GateModuleOptions) *GateModule {
	return &GateModule{
		options: options,
	}
}

func (m *GateModule) OnInit() {
	// log.Debug("gate OnInit")
	m.Gate = &Gate{
		MaxConnNum:      m.options.MaxConnNum,
		PendingWriteNum: m.options.PendingWriteNum,
		MaxMsgLen:       m.options.MaxMsgLen,
		WSAddr:          m.options.WSAddr,
		HTTPTimeout:     m.options.HTTPTimeout,
		CertFile:        m.options.CertFile,
		KeyFile:         m.options.KeyFile,
		TCPAddr:         m.options.TCPAddr,
		LenMsgLen:       m.options.LenMsgLen,
		LittleEndian:    m.options.LittleEndian,
		Processor:       m.options.Processor,
		AgentChanRPC:    m.options.AgentChanRPC,
	}
}
