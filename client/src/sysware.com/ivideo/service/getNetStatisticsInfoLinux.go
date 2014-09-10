package service

import (
	"fmt"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
	"time"
)

var netInfos = make(map[string][2]*model.NetCardInfo)

func init() {
	go updateNetInfos()
}

func updateNetInfos() {
	for {
		updateNetStatisticsInfo()
		time.Sleep(common.STATISTICS_TIME_INTERVAL_SEC * time.Second)
	}
}

func updateNetStatisticsInfo() error {
	netCardInfos, err := NewGetNetInfo().GetInfo()
	if nil != err {
		return err
	}

	for i := 0; i < len(netCardInfos); i++ {
		iface := netCardInfos[i].Iface
		netInfos[iface] = [2]*model.NetCardInfo{&netCardInfos[i], netInfos[iface][0]}
	}
	return nil
}

type GetNetStatisticsInfoLinux struct {
}

func (getNetStatisticsInfoLinux *GetNetStatisticsInfoLinux) GetInfo() ([][]string, error) {
	ret := [][]string{}
	for iFace, netcards := range netInfos {
		var netRecvBytes int64 = 0
		if netcards[1] != nil {
			netRecvBytes = netcards[0].InBytes - netcards[1].InBytes
		}
		var netTransBytes int64 = 0
		if netcards[1] != nil {
			netTransBytes = netcards[0].OutBytes - netcards[1].OutBytes
		}
		var netTotalBytes int64 = 0
		if netcards[1] != nil {
			netTotalBytes = netcards[0].TotalBytes - netcards[1].TotalBytes
		}
		var netTotalDroppedBytes int64 = 0
		if netcards[1] != nil {
			netTotalDroppedBytes = netcards[0].TotalDropped - netcards[1].TotalDropped
		}

		item := []string{
			iFace,
			fmt.Sprintf("%s/s", utils.DisplaySizeStr(float64(netRecvBytes))),
			fmt.Sprintf("%s/s", utils.DisplaySizeStr(float64(netTransBytes))),
			fmt.Sprintf("%s/s", utils.DisplaySizeStr(float64(netTotalBytes))),
			fmt.Sprintf("%d/s", netTotalDroppedBytes),
		}
		ret = append(ret, item)
	}
	return ret, nil
}
