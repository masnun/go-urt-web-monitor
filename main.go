package main

import (
	"flag"
	gout "github.com/masnun/gout/library"
)

const (
	HOST = "5.135.165.34"
	PORT = "27001"
)

func main() {

	// Storage for the Server data
	var server_data gout.Server = gout.Server{}
	// A two way channel for monitoring server data
	var server_channel chan gout.Server = make(chan gout.Server)


	port := flag.String("port", "8765", "Target port on server")
	flag.Parse()
	

	// Launch the processes
	go MonitorServer(server_channel)
	go StartWebServer(*port, &server_data)

	// Infinite loop: track the channel and update data
	for {
		TrackChannel(&server_data, server_channel)
	}

}

func TrackChannel(data *gout.Server, server_channel chan gout.Server) {
	*data = <-server_channel
}
