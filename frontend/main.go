package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	HOSTS_SEPARATOR = ";"
	HOMEPAGE        = "index.html"
)

var (
	PORT  string
	Hosts []string // list of hosts - delimiter - ";"
)

func hostHeaderCheck(req *http.Request) bool {
	host := req.Host
	// Strip off the :port if the server is running on a non-default port
	if strings.Contains(host, ":") {
		host = strings.Split(host, ":")[0]
	}
	return hostInHosts(host)
}

func hostInHosts(host string) bool {
	if len(Hosts) == 0 {
		return true
	}

	for _, h := range Hosts {
		if h == host {
			return true
		}
	}
	return false

}

func parseHostsFromEnv() {
	found, ok := os.LookupEnv("HOSTS_TO_SERVE")
	if !ok {
		return
	}
	Hosts = strings.Split(found, HOSTS_SEPARATOR)
	fmt.Println("Serving Requests only from these Hosts: ", Hosts)
}

func parseServerPortFromEnv() {
	found, ok := os.LookupEnv("SERVER_PORT")
	if ok {
		PORT = found
	} else {
		PORT = "8080"
	}

}

func init() {
	parseHostsFromEnv()
	parseServerPortFromEnv()

}

func main() {
	fmt.Println("HTTP Server Started and Running on PORT:", PORT)
	fileServer := http.FileServer(http.Dir("./webpage"))
	http.Handle("/", fileServer)

	http.ListenAndServe(":"+PORT, nil)

}
