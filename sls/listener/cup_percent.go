package listener

import (
	"github.com/pkg6/oamc/sls/data"
	"github.com/pkg6/oamc/sls/lang"
	"github.com/shirou/gopsutil/v4/cpu"
	"time"
)

type CPUPercent struct {
	err error
}

func (c *CPUPercent) SetGlobal() error {
	cpuPercent, c.err = cpu.Percent(time.Second, false)
	return c.err
}

func (c *CPUPercent) Err() error {
	return c.err
}

func (c CPUPercent) Name() string {
	return lang.Get(lang.CPUPercent)
}

func (c CPUPercent) String() string {
	return c.Data().String()
}

func (c CPUPercent) Data() *data.Data {
	dt := &data.Data{}
	dt.Format = "%f"
	for _, v := range cpuPercent {
		dt.Data = append(dt.Data, v)
	}
	return dt
}
