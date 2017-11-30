package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	host, port := getArgs(os.Args[1])
	if checkAvailability(host, port) {
		fmt.Println("Listening")
		os.Exit(1)
		return
	} else {
		fmt.Println("Not Listening")
		return
	}
}

func checkAvailability(host string, port string) bool {
	_, err := net.DialTimeout("tcp", host+":"+port, 5*time.Second)

	if err != nil {
		return false
	} else {
		return true
	}
}

func getArgs(userarg string) (string, string) {
	if strings.ContainsAny(userarg, ":") {
		hostarray := strings.Split(userarg, ":")
		host := hostarray[0]
		port := hostarray[1]
		return host, port
	} else {
		host := "localhost"
		port := userarg
		return host, port
	}
}
