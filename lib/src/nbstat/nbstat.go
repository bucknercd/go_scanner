package nbstat

import (
	"base"
	"strings"
	"fmt"
)

func Query(host string) bool {
	// connect to 137
	timeout := 100
	sock, _ := base.NewSocket(host, 139, "udp", timeout)
	if sock != nil {
		nbstatQuery := []byte{0xac,0xdc,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x20,0x43,0x4b}
		nbstatQuery = append(nbstatQuery, []byte(strings.Repeat("A", 30))...)
		nbstatQuery = append(nbstatQuery, []byte{0x00,0x21,0x00,0x01}...)
		sock.Send(nbstatQuery)
		resp, err := sock.Recv()
		if err != nil {
			fmt.Printf("Error during NBSTAT QUERY: %v\n", err)
			return false
		}

		if len(resp) < 100 {
			fmt.Printf("NBSTAT response too small.\n")
			return false
		}
		return true
	}
	return false
}