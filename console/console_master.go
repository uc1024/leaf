package console

import (
	"math"
	"strconv"

	"github.com/uc1024/leaf/conf"
	"github.com/uc1024/leaf/network"
)

type IConsoleMaster interface {
	Init()
	Destroy()
}

type ConsoleMasterOptions struct {
	ConsolePort int    `json:"console_port" yaml:"consolePort"`
	ProfilePath string `json:"profile_path" yaml:"profilePath"`
}

type ConsoleMaster struct {
	options ConsoleMasterOptions
	server  *network.TCPServer
}

func NewConsoleMaster(options ConsoleMasterOptions) IConsoleMaster {
	return &ConsoleMaster{
		options: options,
	}
}

func (master *ConsoleMaster) Init() {
	if master.options.ConsolePort == 0 {
		return
	}

	master.server = new(network.TCPServer)
	master.server.Addr = "localhost:" + strconv.Itoa(conf.ConsolePort)
	master.server.MaxConnNum = int(math.MaxInt32)
	master.server.PendingWriteNum = 100
	master.server.NewAgent = newAgent

	master.server.Start()
}

func (master *ConsoleMaster) Destroy() {
	if master.server != nil {
		master.server.Close()
	}
}
