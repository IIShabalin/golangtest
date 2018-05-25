package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

var serverPort = os.Getenv("TEST_PORT")
var processDirectory, _ = os.Getwd()

type Network struct {
	Addresses []net.Addr `json:addresses`
	Port      string     `json:port`
}

type Process struct {
	Directory string
}

type State struct {
	Process Process `json:process`
	Network Network `json:network`
}

func buildState() State {
	addrs, _ := net.InterfaceAddrs()
	return State{
		Process: Process{
			Directory: processDirectory,
		},
		Network: Network{
			Addresses: addrs,
			Port:      serverPort,
		},
	}
}

func stateHandler(w http.ResponseWriter, r *http.Request) {
	state := buildState()
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	err := enc.Encode(state)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Marshalling error"))
	}
}

func main() {
	http.HandleFunc("/state", stateHandler)

	fmt.Println("Serving port:", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
}
