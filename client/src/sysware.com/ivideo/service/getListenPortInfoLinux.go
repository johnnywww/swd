package service

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"strings"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/utils"
)

type GetListenPortInfoLinux struct {
}

func (getListenPortInfoLinux *GetListenPortInfoLinux) GetInfo() ([]int64, error) {
	osErr := getLinuxOSError()
	if nil != osErr {
		return []int64{}, osErr
	}
	bs, err := utils.CmdOutputBytes("ss", "-n", "-l")
	if nil != err {
		return []int64{}, err
	}

	reader := bufio.NewReader(bytes.NewBuffer(bs))

	// ignore the first line
	var line []byte
	line, _, err = reader.ReadLine()
	if err == io.EOF || nil != err {
		return []int64{}, err
	}

	ret := []int64{}

	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF || nil != err {
			break
		}

		arr := strings.Fields(string(line))
		arrlen := len(arr)

		if arrlen != 4 && arrlen != 5 {
			log.WriteLog("output of [ss -n -l] format error")
			continue
		}

		ci := 2
		if arrlen == 5 {
			ci = 3
		}

		location := strings.LastIndex(arr[ci], ":")
		port := arr[ci][location+1:]

		if p, e := strconv.ParseInt(port, 10, 64); e != nil {
			log.WriteLog("parse port to int64 fail: %s", e)
			continue
		} else {
			ret = append(ret, p)
		}

	}

	ret = utils.SliceUniqueInt64(ret)

	return ret, nil
}
