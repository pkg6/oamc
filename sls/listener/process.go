package listener

import (
	"fmt"
	"github.com/pkg6/oamc/sls/data"
	"github.com/pkg6/oamc/sls/lang"
	"github.com/shirou/gopsutil/v4/process"
	"sort"
	"strings"
)

var (
	ProcessStatus = []string{
		process.Running,
	}
)

type Process struct {
	err error
}

func (p *Process) SetGlobal() error {
	processProcess, p.err = process.Processes()
	return p.err
}

func (p *Process) Err() error {
	return p.err
}

func (p Process) String() string {
	return p.Data().String()
}

func (p Process) Name() string {
	return lang.Get(lang.Processes)
}

// Data
//rss：已分配到物理内存的常驻集大小。这是进程当前正在使用的实际物理内存量。
//vms：虚拟内存大小。这是进程当前使用的虚拟内存总量。
//hwm：高水位标记。这是进程曾经使用过的最大物理内存量，但现在已经释放。
//data：数据段大小。这是进程数据段的大小，包括初始化的数据、未初始化的数据和堆。
//stack：堆栈大小。这是进程堆栈的大小，包括函数调用和本地变量。
//locked：锁定大小。这是被进程锁定的物理内存大小，通常用于防止页面置换。
//swap：交换区使用量。这是进程当前使用的交换空间的大小，用于存放暂时不需要的内存数据。
func (p Process) Data() *data.Data {
	dt := &data.Data{}
	dt.Format = "%s"

	var sortMemory []int
	var mapMemory = map[int]string{}
	for _, process := range processProcess {
		b, _ := process.IsRunning()
		name, _ := process.Name()
		cwd, _ := process.Cwd()
		memory, err := process.MemoryInfo()
		if err != nil {
			continue
		}
		if b && (memory.RSS > 0 || memory.VMS > 0 || memory.HWM > 0 || memory.Data > 0 || memory.Stack > 0 || memory.Locked > 0 || memory.Swap > 0) {
			s := fmt.Sprintf("%d|%s|%s|%s|%s|%s|%s|%s|%s|%s",
				process.Pid,
				name,
				cwd,
				data.FormatBytes(float64(memory.RSS), "MB"),
				data.FormatBytes(float64(memory.VMS), "MB"),
				data.FormatBytes(float64(memory.HWM), "MB"),
				data.FormatBytes(float64(memory.Data), "MB"),
				data.FormatBytes(float64(memory.Stack), "MB"),
				data.FormatBytes(float64(memory.Locked), "MB"),
				data.FormatBytes(float64(memory.Swap), "MB"),
			)
			mRSS := int(memory.RSS)
			if _, ok := mapMemory[mRSS]; ok {
				continue
			}
			sortMemory = append(sortMemory, mRSS)
			mapMemory[mRSS] = s
		}
	}
	sort.Slice(sortMemory, func(i, j int) bool {
		return sortMemory[i] > sortMemory[j]
	})
	var ps []string
	for _, memory := range sortMemory {
		if s, ok := mapMemory[memory]; ok {
			ps = append(ps, s)
		}
	}
	dt.Data = []any{strings.Join(ps, "\n")}
	return dt
}
