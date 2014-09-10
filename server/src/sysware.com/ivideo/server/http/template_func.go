package http

import (
	"fmt"
	"html/template"
	"strings"
	"sysware.com/ivideo/common"
)

func getPageStartIndex(pageNo int) int {
	return pageNo*10 + 1
}

func getPageEndIndex(pageNo int) int {
	return (pageNo + 1) * 10
}

func getPageHref(title string, pageNo int, href string) string {
	if pageNo < 0 {
		return fmt.Sprintf("<a href=\"#\">%s</a>", title)
	} else {
		return fmt.Sprintf("<a href=\"%s?pageNo=%d\">%s</a>", href, pageNo, title)
	}
}

func getPageLiPageHref(liClass string, title string, pageNo int, href string) string {
	if pageNo < 0 {
		return fmt.Sprintf("<li class=\"%s disabled \">%s</li>", liClass, getPageHref(title, pageNo, href))
	} else {
		return fmt.Sprintf("<li class=\"%s\">%s</li>", liClass, getPageHref(title, pageNo, href))
	}
}

func getPageIndexHtml(pageNo int, totalPage int, href string) template.HTML {
	pageHtml := []string{}
	if pageNo < 1 {
		pageHtml = append(pageHtml, getPageLiPageHref("prev disabled", common.PAGE_TITLE_PREV, -1, "#"))
	} else {
		pageHtml = append(pageHtml, getPageLiPageHref("prev", common.PAGE_TITLE_PREV, pageNo-1, href))
	}
	pageHtml = append(pageHtml, getPageLiPageHref("active", fmt.Sprintf("%d", pageNo+1), pageNo, href))
	endPageNo := totalPage - 1
	if (pageNo + 5) < (totalPage - 1) {
		endPageNo = pageNo + 5
	}
	for i := pageNo + 1; i <= endPageNo; i++ {
		pageHtml = append(pageHtml, getPageLiPageHref("", fmt.Sprintf("%d", i), i, href))
	}
	if endPageNo < totalPage-1 {
		pageHtml = append(pageHtml, getPageLiPageHref("next", common.PAGE_TITLE_NEXT, endPageNo, href))
	} else {
		pageHtml = append(pageHtml, getPageLiPageHref("next", common.PAGE_TITLE_NEXT, -1, "#"))
	}
	return template.HTML(strings.Join(pageHtml, "\n"))
}

func getServerType(serverType int) string {
	switch serverType {
	case common.SERVER_TYPE_SIP:
		return "SIP服务器"
	case common.SERVER_TYPE_CMS:
		return "中心管理服务器"
	case common.SERVER_TYPE_MTS:
		return "转发服务器"
	case common.SERVER_TYPE_APS:
		return "报警服务器"
	default:
		return ""
	}
}
