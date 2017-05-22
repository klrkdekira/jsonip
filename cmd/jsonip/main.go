package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/klrkdekira/jsonip"
)

func main() {
	var isIPv6 bool
	flag.BoolVar(&isIPv6, "ipv6", false, "Show IPv6")
	flag.Parse()

	var data *jsonip.JSONIP
	var err error
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
