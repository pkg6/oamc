package listener

import (
	"fmt"
	"github.com/pkg6/oamc/sls/data"
	"github.com/pkg6/oamc/sls/lang"
	"github.com/shirou/gopsutil/v4/disk"
	"strings"
)

type Disk struct {
	err error
}

func (d Disk) Name() string {
	return lang.Get(lang.DiskPartitionsUsage)
}
func (d *Disk) SetGlobal() error {
	diskPartitions, d.err = disk.Partitions(true)
	return d.err
}
func (d *Disk) Err() error {
	return d.err
}
func (d Disk) String() string {
	return d.Data().String()
}

func (d Disk) Data() *data.Data {
	dt := &data.Data{}
	format := lang.Get(lang.DiskPartitionsUsagePath) + ":%s " +
		lang.Get(lang.DiskPartitionsUsageTotal) + ":%s " +
		lang.Get(lang.DiskPartitionsUsageUsed) + ":%s " +
		lang.Get(lang.DiskPartitionsUsageFree) + ":%s"
	dt.Format = "%s"
	var paths []string
	for _, partition := range diskPartitions {
		usageStat, _ := disk.Usage(partition.Mountpoint)
		paths = append(paths, fmt.Sprintf(format,
			usageStat.Path,
			data.FormatBytes(float64(usageStat.Total), "GB"),
			data.FormatBytes(float64(usageStat.Used), "GB"),
			data.FormatBytes(float64(usageStat.Free), "GB"),
		))
	}
	dt.Data = append(dt.Data,
		strings.Join(paths, "\n"),
	)
	return dt
}
