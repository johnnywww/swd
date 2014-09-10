package utils

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/log"
	"time"
)

const (
	SERVER_STATE_CHECK_INTERVAL_SEC = 2
)

func GetServerState(address string) int {
	result := common.SERVER_STATE_OFFLINE
	req, err1 := http.NewRequest("GET", address, nil)
	if nil != err1 {
		log.WriteLog("errror %v", err1)
		return result
	}
	httpClient := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(SERVER_STATE_CHECK_INTERVAL_SEC * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*SERVER_STATE_CHECK_INTERVAL_SEC)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.WriteLog("errror %v", err)
		return result
	}
	// 保证I/O正常关闭
	defer resp.Body.Close()
	// 判断返回状态
	if resp.StatusCode == http.StatusOK {
		result = common.SERVER_STATE_ONLINE
	}
	return result
}

func IsEmptyStr(value string) bool {
	return len(value) < 1
}

func IsEmptyPStr(value *string) bool {
	if nil == value {
		return true
	}
	return IsEmptyStr(*value)
}

func IsAdminUser(userName string) bool {
	return common.USER_NAME_ADMIN == userName
}

func TrimRightSpace(s string) string {
	return strings.TrimRight(string(s), "\r\n\t ")
}

func SliceUniqueInt64(s []int64) []int64 {
	size := len(s)
	if size == 0 {
		return []int64{}
	}

	m := make(map[int64]bool)
	for i := 0; i < size; i++ {
		m[s[i]] = true
	}

	realLen := len(m)
	ret := make([]int64, realLen)

	idx := 0
	for key := range m {
		ret[idx] = key
		idx++
	}
	return ret
}

func DisplaySizeStr(raw float64) string {
	if raw < common.KB_SIZE {
		return fmt.Sprintf("%.1fB", raw)
	}

	if raw < common.MB_SIZE {
		return fmt.Sprintf("%.1fKB", raw/common.KB_SIZE)
	}

	if raw < common.GB_SIZE {
		return fmt.Sprintf("%.1fMB", raw/common.MB_SIZE)
	}

	if raw < common.TB_SIZE {
		return fmt.Sprintf("%.1fGB", raw/common.GB_SIZE)
	}

	if raw < common.PB_SIZE {
		return fmt.Sprintf("%.1fTB", raw/common.TB_SIZE)
	}

	if raw < common.PB_SIZE*1024 {
		return fmt.Sprintf("%.1fPB", raw/(common.PB_SIZE))
	}

	return "TooLarge"
}
