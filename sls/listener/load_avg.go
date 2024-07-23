package listener

import (
	"github.com/pkg6/oamc/sls/data"
	"github.com/pkg6/oamc/sls/lang"
	"github.com/shirou/gopsutil/v4/load"
)

type LoadAvg struct {
	err error
}

func (l LoadAvg) Name() string {
	return lang.Get(lang.LoadAvg)
}
func (l *LoadAvg) SetGlobal() error {
	loadAvg, l.err = load.Avg()
	return l.err
}
func (l *LoadAvg) Err() error {
	return l.err
}
func (l LoadAvg) String() string {
	return l.Data().String()
}

func (l LoadAvg) Data() *data.Data {
	dt := &data.Data{}
	dt.Format = lang.Get(lang.LoadAvgLoad1) + ":%f " +
		lang.Get(lang.LoadAvgLoad5) + ":%f " +
		lang.Get(lang.LoadAvgLoad15) + ":%f "
	dt.Data = append(dt.Data,
		loadAvg.Load1,
		loadAvg.Load5,
		loadAvg.Load15,
	)
	return dt
}
