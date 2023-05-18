package module

import (
	"runtime"

	"github.com/uc1024/leaf/conf"
	"github.com/uc1024/leaf/log"
)

type IModulesMaster interface {
	Register(mi Module)
	Destroy()
	Init()
}

type ModulesMaster struct {
	mods []*module
}

func NewModulesMaster() *ModulesMaster {
	return &ModulesMaster{
		mods: []*module{},
	}
}

func (master *ModulesMaster) Register(mi Module) {
	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool, 1)
	master.mods = append(master.mods, m)
}

func (master *ModulesMaster) Destroy() {
	for i := len(master.mods) - 1; i >= 0; i-- {
		m := master.mods[i]
		m.closeSig <- true
		m.wg.Wait()
		master.destroy(m)
	}
}

func (master *ModulesMaster) Init() {
	for i := 0; i < len(master.mods); i++ {
		master.mods[i].mi.OnInit()
	}

	for i := 0; i < len(master.mods); i++ {
		m := master.mods[i]
		m.wg.Add(1)
		go master.run(m)
	}
}

func (master *ModulesMaster) run(m *module) {
	m.mi.Run(m.closeSig)
	m.wg.Done()
}

func (master *ModulesMaster) destroy(m *module) {
	defer func() {
		if r := recover(); r != nil {
			if conf.LenStackBuf > 0 {
				buf := make([]byte, conf.LenStackBuf)
				l := runtime.Stack(buf, false)
				log.Error("%v: %s", r, buf[:l])
			} else {
				log.Error("%v", r)
			}
		}
	}()

	m.mi.OnDestroy()
}
