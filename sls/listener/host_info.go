package listener

import (
	"github.com/pkg6/oamc/sls/data"
	"github.com/pkg6/oamc/sls/lang"
	"github.com/shirou/gopsutil/v4/host"
)

type HostInfo struct {
	err error
}

func (h *HostInfo) SetGlobal() error {
	hostInfo, h.err = host.Info()
	return h.err
}
func (h *HostInfo) Err() error {
	return h.err
}
func (h HostInfo) Name() string {
	return lang.Get(lang.HostInfo)
}

func (h HostInfo) String() string {
	return h.Data().String()
}

func (h HostInfo) Data() *data.Data {
	dt := &data.Data{}
	dt.Format = lang.Get(lang.HostInfoHostname) + ":%s " +
		lang.Get(lang.HostInfoKernelVersion) + ":%s " +
		lang.Get(lang.HostInfoKernelUptime) + ":%dDay" +
		lang.Get(lang.HostInfoKernelProcs) + ":%d"
	dt.Data = append(dt.Data,
		hostInfo.Hostname,
		hostInfo.KernelVersion,
		data.FormatSeconds(int(hostInfo.Uptime), "day"),
		hostInfo.Procs,
	)
	return dt
}
