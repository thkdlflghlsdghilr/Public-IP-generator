package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	stdout := flag.Bool("o", true, "Outputs generated IPs to stdout with newlines")
	write := flag.String("w", "", "Outputs generated IPs to file separated with newlines")
	timeBool := flag.Bool("t", true, "Prints time it took to generate IPs once finished")
	flag.Parse()
	start := time.Now()
	f, err := os.OpenFile(*write, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Gross nested for loops but idk a better way to do this in golang
	for i := 0; i < 255; i++ {
		for j := 0; j < 255; j++ {
			for k := 0; k < 255; k++ {
				for l := 0; l < 255; l++ {
					// using gross list of ifs to check, it works and it's fast so ima leave it
					// might just move ijkl's value forward instead of returning false to imporve performance
					if is_public_ip(i, j, k, l) {
						if *stdout {
							fmt.Printf("%d.%d.%d.%d\n", i, j, k, l)
						}
						if *write != "" {
							if _, err := f.WriteString(fmt.Sprintf("%d.%d.%d.%d\n", i, j, k, l)); err != nil {
								panic(err)
							}
						}
					}
				}
			}
		}
	}
	if *timeBool {
		fmt.Printf("IP list genaration took %s\n", time.Since(start))
	}
}

func is_public_ip(i int, j int, k int, l int) bool {
	// Non-public IP ranges and reason for why they are not public (source: https://en.wikipedia.org/wiki/Reserved_IP_addresses)):
	// 0.0.0.0 - 0.255.255.255 Current network
	// 10.0.0.0 - 10.255.255.255 Used for local communications within a private network
	// 100.64.0.0 - 100.127.255.255 Shared address space
	// 127.0.0.0 - 127.255.255.255 Used for loopback addresses to the local host
	// 169.254.0.0 - 169.254.255.255 Used for link-local addresses
	// 172.16.0.0 - 172.31.255.255 Used for local communications within a private network
	// 192.0.0.0 - 192.0.0.255 IETF Protocol Assignments
	// 192.0.2.0 - 192.0.2.255 Assigned as TEST-NET-1, documentation and examples
	// 192.88.99.0 - 192.88.99.255 Reserved
	// 192.168.0.0 - 192.168.255.255 Used for local communications within a private network
	// 198.18.0.0 - 198.19.255.255 Used for benchmark testing of inter-network communications between two separate subnets
	// 198.51.100.0 - 198.51.100.255 Assigned as TEST-NET-2, documentation and examples
	// 203.0.113.0 - 203.0.113.255 Assigned as TEST-NET-3, documentation and examples
	// 224.0.0.0 - 239.255.255.255 In use for IP multicast
	// 233.252.0.0 - 233.252.0.255 Assigned as MCAST-TEST-NET, documentation and examples
	// 240.0.0.0 - 255.255.255.255 Reserved for future use + Reserved for the "limited broadcast" destination address

	// This list is not the same order as the above list. My OCD is killing me but I'm too lazy to fix it XD
	// also very icky list of if statements, there has got to be a better way to do this
	if i == 10 && check_range(0, 255, j) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 172 && check_range(16, 31, j) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 192 && j == 168 && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 169 && j == 254 && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 127 && check_range(0, 255, j) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 0 && check_range(0, 255, j) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 192 && j == 0 && k == 0 && check_range(0, 255, l) {
		return false
	} else if i == 192 && j == 0 && k == 2 && check_range(0, 255, l) {
		return false
	} else if i == 192 && j == 88 && k == 99 && check_range(0, 255, l) {
		return false
	} else if i == 198 && (j == 18 || j == 19) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 198 && j == 51 && k == 100 && check_range(0, 255, l) {
		return false
	} else if i == 100 && check_range(64, 127, j) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 203 && j == 0 && k == 113 && check_range(0, 255, l) {
		return false
	} else if check_range(224, 239, l) && check_range(0, 255, j) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else if i == 233 && j == 252 && k == 0 && check_range(0, 255, l) {
		return false
	} else if check_range(240, 255, i) && check_range(0, 255, j) && check_range(0, 255, k) && check_range(0, 255, l) {
		return false
	} else {
		return true
	}
}

// This is inclusive!!!
func check_range(min int, max int, value int) bool {
	if value >= min && value <= max {
		return true
	} else {
		return false
	}
}
