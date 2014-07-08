package main

import (
	"flag"
	"fmt"
	gout "github.com/masnun/gout/library"
	"strconv"
	"strings"
)

func main() {

	// Storage for the Server data
	var server_data gout.Server = gout.Server{}
	// A two way channel for monitoring server data
	var server_channel chan gout.Server = make(chan gout.Server)

	webport := flag.String("port", "8765", "Target port to launch the web server")
	delay := flag.String("delay", "5", "The interval (in seconds) between server data refresh")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Please enter a server address. Example: urtbd.com:1111")

	} else {
		server := args[0]
		server_parts := strings.Split(server, ":")
		host := server_parts[0]
		port := server_parts[1]

		delay_sec, err := strconv.Atoi(*delay)
		if err != nil {
			delay_sec = 5
		}

		// Launch the processes
		go MonitorServer(host, port, delay_sec, server_channel)
		go StartWebServer(*webport, &server_data)

		// Infinite loop: track the channel and update data
		for {
			TrackChannel(&server_data, server_channel)
		}
	}
}

func TrackChannel(data *gout.Server, server_channel chan gout.Server) {
	*data = <-server_channel
}
