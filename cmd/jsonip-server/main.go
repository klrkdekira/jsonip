package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	_ "expvar"

	_ "net/http/pprof"

	"github.com/klrkdekira/jsonip"
)

func main() {
	var httpAddr string
	flag.StringVar(&httpAddr, "http", ":80", "http address")
	flag.Parse()

	http.HandleFunc("/", ipv4)
	http.HandleFunc("/ipv6", ipv6)

	err := http.ListenAndServe(httpAddr, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ipv4(w http.ResponseWriter, r *http.Request) {
	data, err := jsonip.IPv4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, data.IP)
}

func ipv6(w http.ResponseWriter, r *http.Request) {
	data, err := jsonip.IPv6()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, data.IP)
}
