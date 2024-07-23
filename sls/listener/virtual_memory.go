package listener

import (
	"github.com/pkg6/oamc/sls/data"
	"github.com/pkg6/oamc/sls/lang"
	"github.com/shirou/gopsutil/v4/mem"
)

type VirtualMemory struct {
	err error
}

func (v VirtualMemory) Name() string {
	return lang.Get(lang.MemVirtualMemory)
}
func (v *VirtualMemory) SetGlobal() error {
	memVirtualMemory, v.err = mem.VirtualMemory()
	return v.err
}
func (v *VirtualMemory) Err() error {
	return v.err
}
func (v VirtualMemory) String() string {
	return v.Data().String()
}

func (v VirtualMemory) Data() *data.Data {
	dt := &data.Data{}
	dt.Format = lang.Get(lang.MemVirtualMemoryTotal) + ":%s " +
		lang.Get(lang.MemVirtualMemoryFree) + ":%s " +
		lang.Get(lang.MemVirtualMemoryUsed) + ":%s "
	dt.Data = append(dt.Data,
		data.FormatBytes(float64(memVirtualMemory.Total), "MB"),
		data.FormatBytes(float64(memVirtualMemory.Free), "MB"),
		data.FormatBytes(float64(memVirtualMemory.Used), "MB"),
	)
	return dt
}
