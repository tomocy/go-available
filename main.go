package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checkAndShowIfDomainNameIsAvailable(scanner.Text())
		time.Sleep(1 * time.Second)
	}
}

func checkAndShowIfDomainNameIsAvailable(domainName string) {
	if checkIfDomainNameIsAvailable(domainName) {
		fmt.Println(domainName, ": ○")
		return
	}
	fmt.Println(domainName, ": ×")
}

func checkIfDomainNameIsAvailable(domainName string) bool {
	const whoisURL string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisURL+":43")
	if err != nil {
		return false
	}
	defer conn.Close()

	conn.Write([]byte(domainName + "\r\n"))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		lowerText := strings.ToLower(scanner.Text())
		if !strings.Contains(lowerText, "no match") {
			return false
		}

		break
	}

	return true
}
