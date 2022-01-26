package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	TULIP_FILENAME         = "tulip.jpeg"
	IMAGES_DIRECTORY       = "images/"
	HOSTS_SEPARATOR        = ";"
	ALCHEMIST_FILENAME     = "the-alchemist.png"
	ATOMIC_HABITS_FILENAME = "atomic-habits.png"
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

func tulip(w http.ResponseWriter, req *http.Request) {
	if !hostHeaderCheck(req) {
		w.WriteHeader(http.StatusForbidden)
		s := fmt.Sprintf("Host \"%s\" is not permitted to access this resource", req.Host)
		w.Write([]byte(s))
		return
	}

	fileBytes, err := ioutil.ReadFile(IMAGES_DIRECTORY + TULIP_FILENAME)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func theAlchemist(w http.ResponseWriter, req *http.Request) {
	if !hostHeaderCheck(req) {
		w.WriteHeader(http.StatusForbidden)
		s := fmt.Sprintf("Host \"%s\" is not permitted to access this resource", req.Host)
		w.Write([]byte(s))
		return
	}

	fileBytes, err := ioutil.ReadFile(IMAGES_DIRECTORY + ALCHEMIST_FILENAME)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func atomicHabits(w http.ResponseWriter, req *http.Request) {
	if !hostHeaderCheck(req) {
		w.WriteHeader(http.StatusForbidden)
		s := fmt.Sprintf("Host \"%s\" is not permitted to access this resource", req.Host)
		w.Write([]byte(s))
		return
	}

	fileBytes, err := ioutil.ReadFile(IMAGES_DIRECTORY + ATOMIC_HABITS_FILENAME)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

func main() {
	fmt.Println("HTTP Server Started and Running on PORT:", PORT)
	http.HandleFunc("/tulip", tulip)
	http.HandleFunc("/alchemist", theAlchemist)
	http.HandleFunc("/atomichabits", atomicHabits)
	fileServer := http.FileServer(http.Dir("./images"))
	http.Handle("/", fileServer)

	http.ListenAndServe(":"+PORT, nil)

}
