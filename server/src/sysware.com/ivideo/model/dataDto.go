package model

import (
	"encoding/json"
)

type DataDto struct {
	Succ bool
	Msg  string
	Data interface{}
}

func getErrorDto(message string) DataDto {
	return DataDto{Succ: false, Msg: message}
}

func getDataDto(d interface{}) DataDto {
	return DataDto{Succ: true, Msg: "", Data: d}
}

func GetErrorDtoJson(message string) string {
	dto := getErrorDto(message)
	bs, err := json.Marshal(dto)
	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}

func GetErrorObjDtoJson(err error) string {
	return GetErrorDtoJson(err.Error())
}

func GetDataDtoJson(d interface{}) string {
	dto := getDataDto(d)
	bs, err := json.Marshal(dto)

	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}
