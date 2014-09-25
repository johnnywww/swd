package service

import (
	"errors"
	"os"
	"syscall"
	"sysware.com/ivideo/log"
)

type executeLinuxProcForkChildImpl struct {
}

func (executeLinuxProcForkChildImpl *executeLinuxProcForkChildImpl) Exec(cmdLine string, dir string) error {
	executeLinuxProc := executeLinuxProcStartProcessImpl{}
	pid, _, errno := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if errno != 0 {
		return syscall.Errno(errno)
	}

	if pid > 0 {
		puid := int(pid)
		processInfo, _ := os.FindProcess(puid)
		if nil != processInfo {
			_, err := processInfo.Wait()
			if nil != err {
				log.WriteLog("Error: find process %s", err.Error())
			}
		}
		return nil
	} else if pid < 0 {
		return errors.New("fork error")
	}

	// pid1, _, errno1 := syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	// if errno1 != 0 {
	// 	os.Exit(0)
	// 	return errors.New("fork error")
	// }
	// if pid1 > 0 {
	// 	os.Exit(0)
	// 	return nil
	// } else if pid1 < 0 {
	// 	os.Exit(0)
	// 	return errors.New("fork error")
	// }
	_ = syscall.Umask(0)
	_, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.WriteLog("Error: syscall.Setsid errno: %d", s_errno)
	}
	os.Chdir("/")
	result := executeLinuxProc.Exec(cmdLine, dir)
	os.Exit(0)
	return result
}
