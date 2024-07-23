package lang

const (
	ZHCN = "zh-CN"
	EN   = "en"
)

var (
	zh = &Language{
		CPUPercent: "CPU使用率",

		LoadAvg:       "工作负载",
		LoadAvgLoad1:  "1分钟",
		LoadAvgLoad5:  "5分钟",
		LoadAvgLoad15: "15分钟",

		MemVirtualMemory:      "内存",
		MemVirtualMemoryTotal: "总内存",
		MemVirtualMemoryFree:  "空闲内存",
		MemVirtualMemoryUsed:  "已使用内存",

		HostInfo:              "主机信息",
		HostInfoHostname:      "主机名",
		HostInfoKernelVersion: "内核版本",
		HostInfoKernelUptime:  "运行时间",
		HostInfoKernelProcs:   "正在运行进程数",

		DiskPartitionsUsage:      "磁盘信息",
		DiskPartitionsUsagePath:  "路径",
		DiskPartitionsUsageTotal: "大小",
		DiskPartitionsUsageUsed:  "已使用",
		DiskPartitionsUsageFree:  "可用",

		Processes: "进程",
	}
	en = &Language{}
)

func Get(k string) string {
	return LocalLanguage.Get(k)
}
