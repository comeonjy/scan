package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/icmp"
)

func main() {
	netaddr, err := net.ResolveIPAddr("ip4", "0.0.0.0")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenIP("ip4:icmp", netaddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, addr, _ := conn.ReadFrom(buf)
		msg, _ := icmp.ParseMessage(1, buf[0:n])
		marshal, err := msg.Marshal(nil)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(marshal))
		fmt.Println(n, addr, msg.Type, msg.Code, msg.Checksum)
	}
}
