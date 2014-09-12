package setup

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sysware.com/ivideo/utils"
)

type propertiesFileServiceImpl struct {
}

func (propertiesFileServiceImpl *propertiesFileServiceImpl) GetInfo(fileName string) (map[string]string, error) {
	if utils.IsEmptyStr(fileName) {
		return nil, errors.New("没有文件")
	}
	if !utils.IsFileExist(fileName) {
		return nil, errors.New(fmt.Sprintf("文件%s不存在", fileName))
	}
	fi, err := os.Open(fileName)
	if nil != err {
		return nil, err
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	result := make(map[string]string)
	for {
		//buf,err := r.ReadBytes('\n')
		str, err := r.ReadString('\n')
		if nil != err {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if !utils.IsEmptyStr(str) {
			strvs := strings.Split(str, "=")
			if len(strvs) > 1 {
			}
			result[strvs[0]] = strvs[1]
		}
	}
	return result, nil
}

func (propertiesFileServiceImpl *propertiesFileServiceImpl) SaveInfo(infos map[string]string, fileName string) error {
	if utils.IsEmptyStr(fileName) {
		return errors.New("没有文件")
	}
	fo, err := os.Create(fileName)
	if nil != err {
		return err
	}
	defer fo.Close()
	w := bufio.NewWriter(fo)
	for key, v := range infos {
		str := key + "=" + v + "\n"
		_, err = w.WriteString(str)
		if nil != err {
			return err
		}
	}
	err = w.Flush()
	return err
}
