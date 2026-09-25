package shared

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	for _, address := range addrs {
		ipnet, ok := address.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipStr := ipnet.IP.String()

				if strings.Contains(ipStr, "192.168.") {
					return ipStr
				}
			}
		}
	}

	return "No IP address found"
}
