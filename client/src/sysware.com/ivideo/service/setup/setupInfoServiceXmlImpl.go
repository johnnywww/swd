package setup

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/utils"
)

type saveXMLFileContentBeforeService interface {
	saveContentBefore(setupInfo interface{}, outFile *os.File) error
}

type setupInfoServiceXmlImpl struct {
	saveXMLContentBefore saveXMLFileContentBeforeService
}

func (setupInfoServiceXmlImpl *setupInfoServiceXmlImpl) GetSetupInfo(setupInfo interface{}, fileName string) error {
	file, err := os.Open(fileName) // For read access.
	if nil != err {
		log.WriteLog("error: %v", err)
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if nil != err {
		log.WriteLog("error: %v", err)
		return err
	}
	if nil == setupInfo {
		log.WriteLog("没有配置对象")
		return err
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(data))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset != common.XML_ENCODE_GB2312 && charset != "GB2312" {
			return input, errors.New("xml: encoding gb2312 is not declared")
		}
		return input, nil
	}

	err = decoder.Decode(setupInfo)
	if nil != err {
		log.WriteLog("error: %v", err)
		return err
	}
	return nil
}

func (setupInfoServiceXmlImpl *setupInfoServiceXmlImpl) SaveSetupInfo(setupInfo interface{}, fileName string, encoding string) error {
	fout, err := os.Create(fileName)
	if nil != err {
		log.WriteLog("error: %v", err)
		return err
	}
	defer fout.Close()
	if nil == setupInfo {
		log.WriteLog("没有配置对象")
		return err
	}
	output, err := xml.MarshalIndent(setupInfo, "  ", "    ")
	if nil != err {
		log.WriteLog("error: %v\n", err)
		return err
	}
	head := `<?xml version="1.0" `
	if !utils.IsEmptyStr(encoding) {
		head = head + "encoding=\"" + encoding + "\""
	}
	head = head + "?>\n"
	fout.Write([]byte(head))
	if nil != setupInfoServiceXmlImpl.saveXMLContentBefore {
		err = setupInfoServiceXmlImpl.saveXMLContentBefore.saveContentBefore(setupInfo, fout)
		if nil != err {
			return err
		}
	}
	fout.Write(output)
	return nil
}
