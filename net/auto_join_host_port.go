//自动获取本地地址，并组装host和port
//代码源自Fabric源码https://github.com/hyperledger/fabric/blob/release/core/peer/peer.go
package main

import (
	"fmt"
	"net"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	fmt.Println("net.InterfaceAddrs", addrs)

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	_, port, _ := net.SplitHostPort("0.0.0.0:7052")
	fmt.Println("net.SplitHostPort 0.0.0.0:7052, port", port)

	address := net.JoinHostPort(GetLocalIP(), port)
	fmt.Println("net.JoinHostPort(GetLocalIP(), port)", address)
}
