package service

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type GetMemInfoLinux struct {
}

func (getMemInfoLinux *GetMemInfoLinux) GetInfo() (*model.MemInfo, error) {
	osErr := getLinuxOSError()
	if nil != osErr {
		return nil, osErr
	}
	want := map[string]bool{
		"Buffers:":   true,
		"Cached:":    true,
		"MemTotal:":  true,
		"MemFree:":   true,
		"SwapTotal:": true,
		"SwapFree:":  true}

	contents, err := ioutil.ReadFile("/proc/meminfo")
	if nil != err {
		return nil, err
	}

	memInfo := &model.MemInfo{}

	reader := bufio.NewReader(bytes.NewBuffer(contents))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(string(line))
		fieldName := fields[0]

		_, ok := want[fieldName]
		if ok && len(fields) == 3 {
			val, numerr := strconv.ParseUint(fields[1], 10, 64)
			if nil != numerr {
				continue
			}
			switch fieldName {
			case "Buffers:":
				memInfo.Buffers = val * common.KB_SIZE
			case "Cached:":
				memInfo.Cached = val * common.KB_SIZE
			case "MemTotal:":
				memInfo.MemTotal = val * common.KB_SIZE
			case "MemFree:":
				memInfo.MemFree = val * common.KB_SIZE
			case "SwapTotal:":
				memInfo.SwapTotal = val * common.KB_SIZE
			case "SwapFree:":
				memInfo.SwapFree = val * common.KB_SIZE
			}
		}
	}
	memInfo.SwapUsed = memInfo.SwapTotal - memInfo.SwapFree

	return memInfo, nil
}
