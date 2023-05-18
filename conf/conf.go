package conf

import "time"

var (
	LenStackBuf = 4096

	// log
	LogLevel string
	LogPath  string
	LogFlag  int

	// console
	ConsolePort   int
	ConsolePrompt string = "Leaf# "
	ProfilePath   string

	// cluster
	ListenAddr      string
	ConnAddrs       []string
	PendingWriteNum int
)

type ServerConfig struct {
	LogLevel    string
	WSAddr      string `json:"ws_addr" yaml:"wsAddr"`           // ws: 地址
	CertFile    string `json:"cert_file" yaml:"certFile"`       // 证书路径
	KeyFile     string `json:"key_file" yaml:"keyFile"`         // key 路径
	TCPAddr     string `json:"tcp_addr" yaml:"tcpAddr"`         // tcp:地址
	MaxConnNum  int    `json:"max_conn_num" yaml:"maxConnNum"`  // 最大连接数
	ConsolePort int    `json:"console_port" yaml:"consolePort"` // 控制台端口
	ProfilePath string `json:"profile_path" yaml:"profilePath"`

	// gate config
	PendingWriteNum int           `json:"pending_write_num" yaml:"pendingWriteNum"` // 2000
	MaxMsgLen       uint32        `json:"max_msg_len" yaml:"maxMsgLen"`             // 4096
	HTTPTimeout     time.Duration `json:"http_timeout" yaml:"httpTimeout"`          // 10 * time.Second
	LenMsgLen       int           `json:"len_msg_len" yaml:"lenMsgLen"`             // 2
	LittleEndian    bool          `json:"little_endian" yaml:"littleEndian"`        // false

	// skeleton conf
	GoLen              int `json:"go_len" yaml:"goLen"` // 10000
	TimerDispatcherLen int `json:"timer_dispatcher_len" yaml:"timerDispatcherLen"`
	AsynCallLen        int `json:"asyn_call_len" yaml:"asynCallLen"` // 10000
	ChanRPCLen         int `json:"chan_rpc_len" yaml:"chanRPCLen"`   // 10000

	// cluster conf
	ListenAddr      string        `json:"listen_addr" yaml:"listenAddr"` // listen addr for clusterin
	ConnAddrs       []string      `json:"conn_addrs" yaml:"connAddrs"`
	ConnectInterval time.Duration `json:"connect_interval" yaml:"connectInterval"`
}
