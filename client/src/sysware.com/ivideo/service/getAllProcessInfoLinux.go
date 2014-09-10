package service

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type GetAllProcessInfoLinux struct {
}

func (getAllProcessInfoLinux *GetAllProcessInfoLinux) GetInfo() ([]model.ProcessInfo, error) {
	result := []model.ProcessInfo{}
	dirs, err := utils.GetDirChildDirs("/proc")
	if nil != err {
		return result, err
	}

	// id dir is a number, it should be a pid. but don't trust it
	dirs_len := len(dirs)
	if dirs_len == 0 {
		return result, nil
	}

	var pid int
	for _, pidstr := range dirs {
		if pid, err = strconv.Atoi(pidstr); nil != err {
			err = nil
			continue
		} else {
			p := model.ProcessInfo{Id: pid}

			err = getAllProcessInfoLinux.getCmdLine(&p)
			if nil != err {
				log.WriteLog("read %s fail: %s", fmt.Sprintf("/proc/%d/cmdline", p.Id), err)
				continue
			}
			err = getAllProcessInfoLinux.getProcStatus(&p)
			if nil != err {
				log.WriteLog("read %s fail: %s", fmt.Sprintf("/proc/%d/status", p.Id), err)
				continue
			}
			err = getAllProcessInfoLinux.getCpuMemRate(&p)
			if nil != err {
				log.WriteLog("read mem: %s", err)
				continue
			}
			err = getAllProcessInfoLinux.getRealPath(&p)
			if nil != err {
				continue
			}
			result = append(result, p)
		}
	}

	return result, nil
}

func (getAllProcessInfoLinux *GetAllProcessInfoLinux) getRealPath(processInfo *model.ProcessInfo) (err error) {
	processInfo.RealPath, err = os.Readlink(fmt.Sprintf("/proc/%d/exe", processInfo.Id))
	return
}

func (getAllProcessInfoLinux *GetAllProcessInfoLinux) getCmdLine(processInfo *model.ProcessInfo) (err error) {
	cmdline_file := fmt.Sprintf("/proc/%d/cmdline", processInfo.Id)
	if !utils.IsFileExist(cmdline_file) {
		err = errors.New("文件不存在")
		return
	}
	processInfo.Cmdline, err = utils.ReadFileToStringNoLn(cmdline_file)
	return
}

func (getAllProcessInfoLinux *GetAllProcessInfoLinux) getCpuMemRate(processInfo *model.ProcessInfo) error {
	bs, err := utils.CmdOutputBytes("ps", "aux")
	if nil != err {
		return err
	}
	reader := bufio.NewReader(bytes.NewBuffer(bs))
	var line []byte
	line, _, err = reader.ReadLine()
	if err == io.EOF || nil != err {
		return errors.New("ps no content")
	}
	var pid int
	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF || nil != err {
			break
		}

		arr := strings.Fields(string(line))
		arrlen := len(arr)

		if arrlen < 7 {
			log.WriteLog("output of [ps aux] format error")
			continue
		}

		pid, err = strconv.Atoi(arr[1])
		if nil != err {
			log.WriteLog("err %v", err)
			continue
		}
		if pid == processInfo.Id {
			processInfo.CPURate, err = strconv.ParseFloat(arr[2], 32)
			if nil != err {
				return err
			}

			processInfo.MemRate, err = strconv.ParseFloat(arr[3], 32)
			if nil != err {
				return err
			}
		}

	}
	return nil
}

func (getAllProcessInfoLinux *GetAllProcessInfoLinux) getProcStatus(processInfo *model.ProcessInfo) (err error) {
	status_file := fmt.Sprintf("/proc/%d/status", processInfo.Id)
	if !utils.IsFileExist(status_file) {
		err = errors.New("文件不存在")
		return
	}
	var content []byte
	content, err = ioutil.ReadFile(status_file)
	if nil != err {
		return
	}
	want := map[string]bool{
		"Name:":   true,
		"State:":  true,
		"VmRSS:":  true,
		"VmSize:": true,
	}
	reader := bufio.NewReader(bytes.NewBuffer(content))
	for {
		bs, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		line := string(bs)
		fields := strings.Fields(string(line))
		_, ok := want[fields[0]]
		if !ok {
			continue
		}
		switch fields[0] {
		case "Name:":
			processInfo.Name = strings.TrimSpace(fields[1])
		case "State:":
			processInfo.State = strings.TrimSpace(fields[1])
		case "VmRSS:":
			processInfo.MemRealSize, err = strconv.ParseInt(fields[1], 10, 32)
			if nil != err {
				return err
			}
			processInfo.MemRealSize = processInfo.MemRealSize * common.KB_SIZE
		case "VmSize:":
			processInfo.MemVmSize, err = strconv.ParseInt(fields[1], 10, 32)
			if nil != err {
				return err
			}
			processInfo.MemVmSize = processInfo.MemVmSize * common.KB_SIZE
		}
	}

	return
}
