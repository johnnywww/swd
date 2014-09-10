package model

type ProcessInfo struct {
	Id          int
	Name        string
	Cmdline     string
	RealPath    string
	State       string
	MemRealSize int64
	MemVmSize   int64
	CPURate     float64
	MemRate     float64
}

type ServerProcessInfo struct {
	ServerInfo
	ProcessInfo
}
