package model

import (
	"testing"
)

func Test_GetPageInfo(t *testing.T) {
	pageInfo := PageInfo{}
	pageInfo.GetPageInfo(3, 3, 10)
	if 3 != pageInfo.PageNo {
		t.Error("Page is Error")
	}
}
