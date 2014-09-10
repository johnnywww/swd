package service

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sysware.com/ivideo/model"
)

type GetCpuInfoLinux struct {
}

func (getCpuInfoLinux *GetCpuInfoLinux) GetInfo() (*model.CpuInfo, error) {
	osErr := getLinuxOSError()
	if nil != osErr {
		return nil, osErr
	}
	contents, err := ioutil.ReadFile("/proc/stat")
	if nil != err {
		return nil, err
	}

	cpuInfo := &model.CpuInfo{}

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(string(line))
		fieldName := fields[0]
		if fieldName == "cpu" {
			parseCPUFields(fields, cpuInfo)
			break
		}
	}
	return cpuInfo, nil
}

func parseCPUFields(fields []string, stat *model.CpuInfo) {
	numFields := len(fields)
	for i := 1; i < numFields; i++ {
		val, err := strconv.ParseUint(fields[i], 10, 64)
		if nil != err {
			continue
		}

		stat.Total += val
		switch i {
		case 1:
			stat.User = val
		case 2:
			stat.Nice = val
		case 3:
			stat.System = val
		case 4:
			stat.Idle = val
		case 5:
			stat.Iowait = val
		case 6:
			stat.Irq = val
		case 7:
			stat.SoftIrq = val
		case 8:
			stat.Steal = val
		case 9:
			stat.Guest = val
		}
	}
}
