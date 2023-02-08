package client

import (
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv6"
)

const ProtocolIPv6ICMP = 58

var pingData []byte

func dial() *icmp.PacketConn {
	conn, err := icmp.ListenPacket("udp6", "::")
	if err != nil {
		log.Fatalf("listen packet err: %v", err)
	}
	return conn
}

func Ping(addr string) {
	conn := dial()
	defer conn.Close()
	err := ping(net.ParseIP(addr), conn)
	if err != nil {
		log.Fatalf("error ping %s: %v", addr, err)
	}
}

func ping(ip net.IP, conn *icmp.PacketConn) error {
	_, err := conn.WriteTo(pingData, &net.UDPAddr{IP: ip, Zone: "en0"})
	if err != nil {
		return err
	}

	timeStart := time.Now()
	conn.SetReadDeadline(time.Now().Add(time.Second * 2)) // timeout 2 s
	rb := make([]byte, 1024)
	n, peer, err := conn.ReadFrom(rb)
	if err != nil {
		return err
	}

	rm, err := icmp.ParseMessage(ProtocolIPv6ICMP, rb[:n])
	if err != nil {
		return err
	}

	data, err := rm.Body.Marshal(ProtocolIPv6ICMP)
	if err != nil {
		return err
	}

	log.Printf("got %d bytes from %v after %v: %s", n, peer, time.Since(timeStart), data)
	return nil
}

func init() {
	switch runtime.GOOS {
	case "darwin", "ios":
	case "linux":
		out, err := exec.Command("sysctl", "net.ipv4.ping_group_range").Output()
		if err != nil {
			log.Fatalf("sysctl err: %v", err)
		}
		if string(out) == "0" {
			log.Println("you may need to adjust the net.ipv4.ping_group_range kernel state")
		}
	default:
		log.Println("not supported on", runtime.GOOS)
		return
	}

	msg := icmp.Message{
		Type: ipv6.ICMPTypeEchoRequest, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	var err error
	pingData, err = msg.Marshal(nil)
	if err != nil {
		log.Fatal("marshal ping data error ", err)
	}
}
