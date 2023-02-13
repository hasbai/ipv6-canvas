package server

import (
	"errors"
	"fmt"
	"github.com/hasbai/ipv6-canvas/lib"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv6"
	"log"
	"net"
	"time"
)

const ProtocolIPv6ICMP = 58

func serveICMP() {
	conn, err := icmp.ListenPacket("ip6:ipv6-icmp", "::")
	if err != nil {
		log.Fatalf("Error listening: %s", err)
	}
	conn6 := conn.IPv6PacketConn()
	if conn6 == nil {
		log.Fatalf("Error getting IPv6 packet conn: %s", err)
	}
	log.Println("ICMP listening on", conn6.LocalAddr())

	for {
		err = handleICMP(conn6)
		if err != nil {
			log.Fatalf("Error serving icmp: %s", err)
		}
	}
}

func handleICMP(conn *ipv6.PacketConn) error {
	buf := make([]byte, 64)
	n, _, addr, err := conn.ReadFrom(buf)
	if err != nil {
		return err
	}

	start := time.Now()

	icmpMsg, err := icmp.ParseMessage(ProtocolIPv6ICMP, buf[:n])
	if err != nil {
		return err
	}
	if icmpMsg.Type != ipv6.ICMPTypeEchoRequest { // only accept ping
		return nil
	}

	ipAddr, ok := addr.(*net.IPAddr)
	if !ok {
		return errors.New(fmt.Sprintf("addr %s is not an IPAddr", addr))
	}

	pixel := lib.IP2Pixel(ipAddr.IP)
	modifyImage(pixel)

	log.Printf("%s %dμs", addr, time.Since(start).Microseconds())

	return nil
}
