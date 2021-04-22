package hostalive
import (
	"tcpscan"
	"fmt"
	"sync"
	"base"
	//"nbstat" // nbstat.Query is hanging. need to add a timeout
	//"arp"
	"icmp"
)

func Scan(ip string, m *sync.Mutex) {
	var tcpPorts []int
	if alive(ip) {
		tcpPorts = tcpscan.Scan(ip)
		m.Lock()
		fmt.Printf("\n%s\nAlive tcp ports: %v\n\n", ip, tcpPorts)
		m.Unlock()
	}
	//arp.Scan(ip)
	//icmp.Scan(ip)
}

func alive(host string) bool {
	timeout := 500
	for _, port := range []int{22, 25, 80, 443, 445} {
		sock, err := base.NewSocket(host, port, "tcp", timeout)
		if sock != nil && err == nil {
			//fmt.Printf("ALIVE: PORT %d detected -------------\n", port)
			sock.Close()
			return true
		}
	}
	if icmp.Request(host) {
		return true
	}
	//arp.Request(host)
	return false
}