package arp

import (
	"net"
	"fmt"
)


/*
ip4:1 => bare ipv4
*/

func Request(host string) string {
	conn, err := net.Dial("ip4:1", host)
	if err != nil {
		fmt.Printf("ARP: %v\n", err)
	}

	var msg [46]byte
	msg[0] = 0    // HW type
	msg[1] = 0x01 // HW type: ethernet
	msg[2] = 0x08 // Proto type
	msg[3] = 0    // Proto type : ipv4
	msg[4] = 0x06 // HW size
	msg[5] = 0x04 // Proto size
	msg[6] = 0    // opcode
	msg[7] = 0x01 // opcode : request
	//message := msg[0:]

	//message = append(message, []byte{})

	conn.Write([]byte("HOLA"))
	conn.Close()
	return ""
}