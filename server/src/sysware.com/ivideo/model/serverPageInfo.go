package model

type ServerPageInfo struct {
	Servers []ServerInfo
	Page    PageInfo
}

type ServerPageHandle interface {
	GetServerPageInfo(pageNo int, perPageRecordSize int) ServerPageInfo
}

func NewServerPageHandle() ServerPageHandle {
	return &ServerPageInfo{}
}

func (serverPageInfo *ServerPageInfo) getSliceServers(pageNo int, perPageRecordSize int, totalServers []ServerInfo) []ServerInfo {
	rcount := len(totalServers)
	pageInfo := NewPageHandle().GetPageInfo(pageNo, perPageRecordSize, rcount)
	startIndex := pageInfo.PageNo * pageInfo.PerPageRecordSize
	endIndex := 0
	if (pageInfo.PageNo+1)*perPageRecordSize > len(totalServers) {
		endIndex = rcount
	} else {
		endIndex = (pageInfo.PageNo + 1) * perPageRecordSize
	}
	if len(totalServers) < 1 {
		return totalServers
	} else {
		return totalServers[startIndex:endIndex]
	}
}

func (serverPageInfo *ServerPageInfo) GetServerPageInfo(pageNo int, perPageRecordSize int) ServerPageInfo {
	if nil == serverPageInfo {
		serverPageInfo = &ServerPageInfo{}
	}
	setupInfo := &SetupInfo{}
	if !NewSetupHandle().GetSetupInfo(setupInfo) {
		return *serverPageInfo
	}
	serverPageInfo.Servers = serverPageInfo.getSliceServers(pageNo, perPageRecordSize, setupInfo.Servers)
	serverPageInfo.Page = NewPageHandle().GetPageInfo(pageNo, perPageRecordSize, len(setupInfo.Servers))
	return *serverPageInfo
}
