package main

import gout "github.com/masnun/gout/library"

func main() {

	// Storage for the Server data
	var server_data gout.Server = gout.Server{}
	// A two way channel for monitoring server data
	var server_channel chan gout.Server = make(chan gout.Server)

	// Launch the processes
	go MonitorServer(server_channel)
	go StartWebServer(&server_data)

	// Infinite loop: handle incoming data
	for {
		UpdateServerData(&server_data, server_channel)
		//PrintPlayerList(server_data)
	}

}

func UpdateServerData(data *gout.Server, server_channel chan gout.Server) {
	*data = <-server_channel
}
