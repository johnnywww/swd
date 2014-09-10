package model

type PageHandle interface {
	GetPageInfo(pageNo int, perPageRecordSize int, recordCount int) PageInfo
}

type PageInfo struct {
	PageNo            int
	RecordCount       int
	PerPageRecordSize int
	TotalPage         int
}

func NewPageHandle() PageHandle {
	return &PageInfo{}
}

func (pageInfo *PageInfo) GetPageInfo(pageNo int, perPageRecordSize int, recordCount int) PageInfo {
	pageInfo.PageNo = pageNo
	pageInfo.PerPageRecordSize = perPageRecordSize
	pageInfo.RecordCount = recordCount
	pageInfo.calcPageInfo()
	return *pageInfo
}

func (pageInfo *PageInfo) calcPageInfo() {
	if pageInfo.RecordCount < 1 {
		pageInfo.TotalPage = 0
		pageInfo.PageNo = -1
		return
	}
	if pageInfo.PerPageRecordSize < 1 {
		if pageInfo.RecordCount > 0 {
			pageInfo.TotalPage = 1
		} else {
			pageInfo.TotalPage = 0
		}
		pageInfo.PageNo = 0
		return
	}
	if (pageInfo.RecordCount % pageInfo.PerPageRecordSize) > 0 {
		pageInfo.TotalPage = pageInfo.RecordCount/pageInfo.PerPageRecordSize + 1
	} else {
		pageInfo.TotalPage = pageInfo.RecordCount / pageInfo.PerPageRecordSize
	}
	if pageInfo.PageNo > pageInfo.TotalPage-1 {
		pageInfo.PageNo = pageInfo.TotalPage - 1
	}
}
