package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.bug.st/serial"
)

const (
	INTERVAL_SEARCH_FOR_PORT = 200 * time.Millisecond
	INTERVAL_BUSY_PORT_RETRY = 500 * time.Millisecond

	URL = "file:///path/to/any/file/or/url"
)

func main() {
	ticker := time.NewTicker(INTERVAL_SEARCH_FOR_PORT)

	for range ticker.C {
		ports, err := serial.GetPortsList()
		if err != nil {
			os.Exit(1)
		}

		port, ok := getPort(ports)
		if !ok {
			continue
		}

		sc := bufio.NewScanner(port)
		for sc.Scan() {
			cmd := Parse(sc.Text())
			if err := cmd(); err != nil {
				log.Printf("failed to run command: %+v\n", err)
			}
		}
	}
}

func isCommand(text string) bool {
	// for now it'll only send OPEN
	return text == "OPEN"
}

func getPort(ports []string) (serial.Port, bool) {
	portName := find(ports)
	if portName == "" {
		return nil, false
	}

	return returnWhenReady(portName)
}

func find(ports []string) string {
	for _, s := range ports {
		if strings.Contains(s, "tty.usbmodem") || strings.Contains(s, "Seeeduino") {
			return s
		}
	}
	return ""
}

func returnWhenReady(port string) (serial.Port, bool) {
	var pErr *serial.PortError

	// gonna recurisve-call the func but goto is safe here
retry:
	p, err := serial.Open(port, &serial.Mode{})
	switch {
	case err == nil:
		return p, true
	case errors.As(err, &pErr) && pErr.Code() == serial.PortBusy:
		time.Sleep(INTERVAL_BUSY_PORT_RETRY)
		goto retry
	default:
		// cannot really recover from any other error than PortBusy, so let it panic
		panic(fmt.Errorf("error opening port: %w", err))
	}
}
