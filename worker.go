package main

import (
	gout "github.com/masnun/gout/library"
	"time"
)

func MonitorServer(channel chan gout.Server) {
	var repeatTimer time.Duration = 5 * time.Second
	for _ = range time.Tick(repeatTimer) {
		//fmt.Println(x)
		var server gout.Server = RefreshServerDetails()
		channel <- server
	}

}

func RefreshServerDetails() gout.Server {
	var response string = gout.GetServerResponse(HOST, PORT)
	return gout.ParseResponse(response)

}
