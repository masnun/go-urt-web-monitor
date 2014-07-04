## Urban Terror Server Monitor

Now faster, lighter and more efficient 



### What is it? 

It is a web based application that monitors a given Urban Terror server and displays the players list. It is updated in a pre-specified time interval. 

The application can also display desktop notifications (if supported by your web browser) when new players join the server. 

Never miss another game with this web application :) 

You can checkout a demo app: <a href="http://masnun.net:8080/">http://masnun.net:8080/</a> 

### Why another version? 

The very first version was written in PHP and it had a lot of issues handling the simultaneous connections and querying the servers. Later another version was created in Python to avoid those problems. It was doing good however we believed that there was scope of improvements. So we revamped the application for better performance and less resource utilization. 

What are the new features in this version? 

* The server is now queried at fixed interval, one background worker does this and serves all requests from memory. Previously the server was queried on every request. That was an unnecessary resource usage. 

* We are using Golang and it's goroutines. Go is insanely fast, goroutines are lightweight threads. High performance gain with low resource consumption. 

* We are bringing in websockets as the default mode of client-server communication. Less http requests, better performance. 

### How can I set it up? 

If you want to run it locally or set up on a server, here's help. The project is written in the "Go" programming language. Please first get Go installed on your system. You can find setup instructions on Google. 

You can easily install the app using the `go get` command. It's simple: 

```bash
	go get github.com/masnun/go-urt-web-monitor
```

Thanks to Go's wonderful tooling system, the above command will download the project and it's dependencies to appropriate directory under your `$GOPATH`. 

Now we need to build the app: 

```bash
	cd $GOPATH/src/github.com/masnun/go-urt-web-monitor
```

That should create a binary in the same directory. Make it execuatble and run it: 

```bash
	chmod a+x go-urt-web-monitor
	./go-urt-web-monitor
```

When the webserver has started, please visit: <a href="http://localhost:8080/">http://localhost:8080/</a> to view the application. 


### How can I contribute? 

Please fork and submit pull requests. If you found any bugs, please create an issue on Github. 


