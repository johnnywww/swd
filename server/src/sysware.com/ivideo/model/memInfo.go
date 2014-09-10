package model

type MemInfo struct {
	Buffers   uint64
	Cached    uint64
	MemTotal  uint64
	MemFree   uint64
	SwapTotal uint64
	SwapUsed  uint64
	SwapFree  uint64
}
