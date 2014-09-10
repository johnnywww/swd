package main

import (
	"runtime"
	"sysware.com/ivideo/client/http"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.StartHttp()
}
