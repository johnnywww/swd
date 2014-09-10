package service

import (
	"errors"
	"strconv"
	"strings"
	"sysware.com/ivideo/utils"
)

type GetSystemRunTimeInfoLinux struct {
}

func (getSystemRunTimeInfoLinux *GetSystemRunTimeInfoLinux) GetInfo() (*RunTimeInfo, error) {
	osErr := getLinuxOSError()
	if nil != osErr {
		return nil, osErr
	}
	var content string
	content, err := utils.ReadFileToStringNoLn("/proc/uptime")
	if nil != err {
		return nil, err
	}

	fields := strings.Fields(content)
	if len(fields) < 2 {
		return nil, errors.New("/proc/uptime parse error")
	}

	secStr := fields[0]
	var secF float64
	secF, err = strconv.ParseFloat(secStr, 64)
	if nil != err {
		return nil, err
	}
	runTimeInfo := &RunTimeInfo{}
	minTotal := secF / 60.0
	hourTotal := minTotal / 60.0

	runTimeInfo.Day = int64(hourTotal / 24.0)
	runTimeInfo.Hour = int64(hourTotal) - runTimeInfo.Day*24
	runTimeInfo.Min = int64(minTotal) - (runTimeInfo.Day * 60 * 24) - (runTimeInfo.Hour * 60)

	return runTimeInfo, nil
}
