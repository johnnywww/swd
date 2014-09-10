package service

import (
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"time"
)

type GetCpuStatisticsInfoLinux struct {
}

var cpuInfos [common.STATISTICS_LIST_CAPACITY]*model.CpuInfo

func init() {
	go updateCpuInfos()
}

func updateCpuInfos() {
	for {
		updateCpuStatisticsInfo()
		time.Sleep(common.STATISTICS_TIME_INTERVAL_SEC * time.Second)
	}
}

func updateCpuStatisticsInfo() error {
	cpuInfo, err := NewGetCpuInfo().GetInfo()
	if nil != err {
		return err
	}

	length := common.STATISTICS_LIST_CAPACITY
	for i := length - 1; i > 0; i-- {
		cpuInfos[i] = cpuInfos[i-1]
	}

	cpuInfos[0] = cpuInfo
	return nil
}

func (getCpuStatisticsInfoLinux *GetCpuStatisticsInfoLinux) GetInfo() (map[string]string, error) {
	result := map[string]string{}
	if nil == cpuInfos[1] {
		return result, nil
	}
	nCpuInfo := cpuInfos[0]
	oCpuInfo := cpuInfos[1]
	dTot := float64(nCpuInfo.Total - oCpuInfo.Total)
	invQuotient := 100.00 / dTot
	dI := float64(nCpuInfo.Idle - oCpuInfo.Idle)
	result["idle"] = fmt.Sprintf("%.1f%%", dI*invQuotient)
	result["busy"] = fmt.Sprintf("%.1f%%", (dTot-dI)*100/dTot)
	result["user"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.User-oCpuInfo.User)*invQuotient)
	result["nice"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.Nice-oCpuInfo.Nice)*invQuotient)
	result["system"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.System-oCpuInfo.System)*invQuotient)
	result["iowait"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.Iowait-oCpuInfo.Iowait)*invQuotient)
	result["irq"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.Irq-oCpuInfo.Irq)*invQuotient)
	result["softirq"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.SoftIrq-oCpuInfo.SoftIrq)*invQuotient)
	result["steal"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.Steal-oCpuInfo.Steal)*invQuotient)
	result["guest"] = fmt.Sprintf("%.1f%%", float64(nCpuInfo.Guest-oCpuInfo.Guest)*invQuotient)
	return result, nil
}
