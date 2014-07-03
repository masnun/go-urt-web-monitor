package main

import (
	"fmt"
	gout "github.com/masnun/gout/library"
	"os"
	"text/tabwriter"
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
	host := "5.135.165.34"
	port := "27001"

	var response string = gout.GetServerResponse(host, port)
	return gout.ParseResponse(response)

}

func PrintPlayerList(server gout.Server) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "name\tping\tpoints")
	for _, player := range server.Players {
		fmt.Fprintln(w, player.Name+"\t"+player.Ping+"\t"+player.Points)
	}
	fmt.Fprintln(w)
	w.Flush()
}
