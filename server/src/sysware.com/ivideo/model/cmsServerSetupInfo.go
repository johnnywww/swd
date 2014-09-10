package model

import (
	"encoding/xml"
	"errors"
	"fmt"
)

type CMSServerSetupInfo struct {
	XMLName            xml.Name `xml:"hibernate-configuration"`
	SessionFactoryInfo CMSServerSessionFactoryInfo
}

type CMSServerSessionFactoryInfo struct {
	XMLName    xml.Name                  `xml:"session-factory"`
	Properties []CMSServerDbPropertyInfo `xml:"property"`
}

type CMSServerDbPropertyInfo struct {
	Value string `xml:",chardata"`
	Name  string `xml:"name,attr"`
}

func (cMSServerSetupInfo *CMSServerSetupInfo) getProperty(propKey string) (*CMSServerDbPropertyInfo, error) {
	var prop *CMSServerDbPropertyInfo = nil
	for i := 0; i < len(cMSServerSetupInfo.SessionFactoryInfo.Properties); i = i + 1 {
		prop = &(cMSServerSetupInfo.SessionFactoryInfo.Properties[i])
		if prop.Name == propKey {
			return prop, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("没有找到属性%s", propKey))
}

func (cMSServerSetupInfo *CMSServerSetupInfo) GetProperty(propKey string) (string, error) {
	prop, err := cMSServerSetupInfo.getProperty(propKey)
	if nil != err {
		return "", err
	}
	return prop.Value, nil
}

func (cMSServerSetupInfo *CMSServerSetupInfo) SetProperty(propKey string, value string) error {
	prop, _ := cMSServerSetupInfo.getProperty(propKey)
	if nil == prop {
		prop1 := CMSServerDbPropertyInfo{value, propKey}
		cMSServerSetupInfo.SessionFactoryInfo.Properties = append(cMSServerSetupInfo.SessionFactoryInfo.Properties, prop1)
		return nil
	}
	prop.Value = value
	return nil
}
