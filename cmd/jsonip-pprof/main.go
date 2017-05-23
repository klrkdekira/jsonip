package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/klrkdekira/jsonip"
)

func main() {
	file, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	var isIPv6 bool
	flag.BoolVar(&isIPv6, "ipv6", false, "Show IPv6")
	flag.Parse()

	var data *jsonip.JSONIP
	if isIPv6 {
		data, err = jsonip.IPv6()
	} else {
		data, err = jsonip.IPv4()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(data.IP)
}
