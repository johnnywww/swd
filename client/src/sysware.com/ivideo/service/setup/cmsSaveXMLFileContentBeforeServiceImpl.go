package setup

import (
	"os"
)

type cmsSaveXMLFileContentBeforeServiceImpl struct {
	setupInfoServiceXmlImpl
}

func NewCMSSetupInfoService() SetupInfoService {
	return &setupInfoServiceXmlImpl{saveXMLContentBefore: &cmsSaveXMLFileContentBeforeServiceImpl{}}
}

func (cmsSaveXMLFileContentBeforeServiceImpl *cmsSaveXMLFileContentBeforeServiceImpl) saveContentBefore(setupInfo interface{}, outFile *os.File) error {
	head := `<!DOCTYPE hibernate-configuration PUBLIC "-//Hibernate/Hibernate Configuration DTD 3.0//EN"
                                         "http://www.hibernate.org/dtd/hibernate-configuration-3.0.dtd">`
	outFile.Write([]byte(head + "\n"))
	return nil
}
