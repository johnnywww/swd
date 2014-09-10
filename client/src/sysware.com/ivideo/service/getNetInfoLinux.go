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

type GetNetInfoLinux struct {
}

func (getNetInfoLinux *GetNetInfoLinux) GetInfo() ([]model.NetCardInfo, error) {
	contents, err := ioutil.ReadFile("/proc/net/dev")
	if nil != err {
		return nil, err
	}

	result := []model.NetCardInfo{}

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		raw_line := string(lineBytes)
		idx := strings.Index(raw_line, ":")
		if idx < 0 {
			continue
		}

		netCardInfo := model.NetCardInfo{}

		line := raw_line[idx+1 : len(raw_line)]
		eth := raw_line[0:idx]
		netCardInfo.Iface = strings.Trim(eth, " ")

		fields := strings.Fields(line)

		netCardInfo.InBytes, _ = strconv.ParseInt(fields[0], 10, 64)
		netCardInfo.InPackages, _ = strconv.ParseInt(fields[1], 10, 64)
		netCardInfo.InErrors, _ = strconv.ParseInt(fields[2], 10, 64)
		netCardInfo.InDropped, _ = strconv.ParseInt(fields[3], 10, 64)
		netCardInfo.InFifoErrs, _ = strconv.ParseInt(fields[4], 10, 64)
		netCardInfo.InFrameErrs, _ = strconv.ParseInt(fields[5], 10, 64)
		netCardInfo.InCompressed, _ = strconv.ParseInt(fields[6], 10, 64)
		netCardInfo.InMulticast, _ = strconv.ParseInt(fields[7], 10, 64)

		netCardInfo.OutBytes, _ = strconv.ParseInt(fields[8], 10, 64)
		netCardInfo.OutPackages, _ = strconv.ParseInt(fields[9], 10, 64)
		netCardInfo.OutErrors, _ = strconv.ParseInt(fields[10], 10, 64)
		netCardInfo.OutDropped, _ = strconv.ParseInt(fields[11], 10, 64)
		netCardInfo.OutFifoErrs, _ = strconv.ParseInt(fields[12], 10, 64)
		netCardInfo.OutCollisions, _ = strconv.ParseInt(fields[13], 10, 64)
		netCardInfo.OutCarrierErrs, _ = strconv.ParseInt(fields[14], 10, 64)
		netCardInfo.OutCompressed, _ = strconv.ParseInt(fields[15], 10, 64)

		netCardInfo.TotalBytes = netCardInfo.InBytes + netCardInfo.OutBytes
		netCardInfo.TotalPackages = netCardInfo.InPackages + netCardInfo.OutPackages
		netCardInfo.TotalErrors = netCardInfo.InErrors + netCardInfo.OutErrors
		netCardInfo.TotalDropped = netCardInfo.InDropped + netCardInfo.OutDropped

		result = append(result, netCardInfo)
	}

	return result, nil
}
