package iperf

import (
	"log"
	"os/exec"
)

var (
	Debug          = false
	binaryLocation = ""
)

func init() {
	if binaryLocation == "" {
		iperfPath, err := exec.LookPath("iperf3")
		if err != nil {
			log.Fatalf("error initializing iperf: %v\n", err)
		}
		binaryLocation = iperfPath
	}
}
