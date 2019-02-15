package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var VERSION = "v1.1"

func main() {
	versionFlag := flag.Bool("v", false, "display version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("dingdong " + VERSION)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("error:  an argument must be given.  Please consult the dingdong manpage for more information.")
		os.Exit(1)
	}
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
