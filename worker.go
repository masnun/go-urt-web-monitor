package main

import (
	"fmt"
	gout "github.com/masnun/gout/library"
	"time"
	"strconv"
)

func MonitorServer(host string, port string, delay int, channel chan gout.Server) {
	fmt.Println("Monitoring: " + host + ":" + port + " with " + strconv.Itoa(delay) + " secs refresh interval")

	var repeatTimer time.Duration = time.Duration(delay) * time.Second
	for _ = range time.Tick(repeatTimer) {
		//fmt.Println(x)
		var server gout.Server = RefreshServerDetails(host, port)
		channel <- server
	}

}

func RefreshServerDetails(host string, port string) gout.Server {
	var response string = gout.GetServerResponse(host, port)
	return gout.ParseResponse(response)

}
