package listener

import (
	"github.com/pkg6/oamc/sls/config"
	"github.com/pkg6/oamc/sls/data"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	"strings"
	"sync"
)

var definitionListeners = map[string]IListener{
	"host_info":   &HostInfo{},
	"cpu_percent": &CPUPercent{},
	"loadavg":     &LoadAvg{},
	"memory":      &VirtualMemory{},
	"disk":        &Disk{},
	"process":     &Process{},
}

var (
	cpuInfo          []cpu.InfoStat
	hostInfo         *host.InfoStat
	cpuPercent       []float64
	diskPartitions   []disk.PartitionStat
	loadAvg          *load.AvgStat
	memVirtualMemory *mem.VirtualMemoryStat
	processProcess   []*process.Process
)

type Listener struct {
	Config *config.Config
	ns     []string
	mls    map[string]IListener
	lock   *sync.Mutex
}

func New(config *config.Config) *Listener {
	l := &Listener{
		mls:    map[string]IListener{},
		lock:   &sync.Mutex{},
		Config: config,
	}
	for _, name := range l.Config.Listener {
		if iListener, ok := definitionListeners[name]; ok {
			l.Add(name, iListener)
		}
	}
	return l
}

func (l *Listener) Add(name string, listener IListener) {
	l.lock.Lock()
	defer l.lock.Unlock()
	_ = listener.SetGlobal()
	l.ns = append(l.ns, name)
	l.mls[name] = listener
}

func (l *Listener) Strings() string {
	return strings.Join(l.Values(), "\n")
}

func (l *Listener) Values() []string {
	l.lock.Lock()
	defer l.lock.Unlock()
	var ss []string
	for _, n := range l.ns {
		l := l.mls[n]
		context := l.String()
		if err := l.Err(); err != nil {
			continue
		}
		ss = append(ss, data.TitleContextLine(l.Name(), 80, "-", "\n"+context+"\n"))
	}
	return ss
}

type IListener interface {
	SetGlobal() error
	Err() error
	String() string
	Name() string
	Data() *data.Data
}
