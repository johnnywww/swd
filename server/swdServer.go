package main

import (
	"sysware.com/ivideo/server/http"
)

func main() {
	http.StartHttp(http.NewRouteHandle())

}
