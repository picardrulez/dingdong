package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var VERSION = "v2.0"

func main() {
	versionFlag := flag.Bool("v", false, "display version")
	humanFlag := flag.Bool("h", false, "human readable")
	flag.Parse()

	if *versionFlag {
		fmt.Println("dingdong " + VERSION)
		return
	}

	if len(flag.Arg(0)) < 2 {
		fmt.Println("error:  an argument must be given.  Please consult the dingdong manpage for more information.")
		os.Exit(2)
	}

	host, port := getArgs(flag.Arg(0))

	if checkAvailability(host, port) {
		if *humanFlag {
			fmt.Println("Listening")
			os.Exit(0)
			return
		} else {
			fmt.Println("1")
			os.Exit(0)
			return
		}
	} else {
		if *humanFlag {
			fmt.Println("Not Listening")
			os.Exit(1)
			return
		} else {
			fmt.Println("0")
			os.Exit(1)
			return
		}
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
