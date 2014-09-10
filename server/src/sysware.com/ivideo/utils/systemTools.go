package utils

import (
	"bytes"
	"os/exec"
	"runtime"
)

func IsLinuxOS() bool {
	return "linux" == runtime.GOOS
}

func IsWindowsOS() bool {
	return "windows" == runtime.GOOS
}

func CmdOutputBytes(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}

func CmdOutputStr(name string, arg ...string) (string, error) {
	value1, err := CmdOutputBytes(name, arg...)
	return string(value1), err
}

func CmdOutputNoLn(name string, arg ...string) (out string, err error) {
	out, err = CmdOutputStr(name, arg...)
	if err != nil {
		return
	}
	return TrimRightSpace(string(out)), nil
}
