package icmp

import (
	"net"
	//"fmt"
	"encoding/binary"
	"time"
)

func Request(host string) bool {
	timeout := 10
	conn, err := net.Dial("ip4:1", host)
	if err != nil {
		//fmt.Printf("ICMP: %v\n", err)
		return false
	}
	var msg [8]byte
	msg[0] = 0x08  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum, fix later
	msg[3] = 0  // checksum, fix later
	msg[4] = 0  // identifier[0]
	msg[5] = 0x02 //identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 0x01 // sequence[1]
	
	message := msg[0:]
	message = append(message, []byte("123456789ABCDEFG")...)

	check := checkSum(message)
	message[2] = byte(check >> 8)
	message[3] = byte(check & 255)

	_, err = conn.Write(message)
	if err != nil {
		//fmt.Println(err)
		return false
	}

	conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	_, err = conn.Read(message)
	if err != nil {
		//fmt.Println(err)
		return false
	}
	return true
}

func checkSum(msg []byte) uint16 {
	var sum int = 0
	for i:=0; i < len(msg); i+=2 {
		sum += int(binary.BigEndian.Uint16(msg[i:i+2]))
	}

	for i:=0; i < 1; i++ {
		sum += (sum >> 16)
	}
	var answer uint16 = uint16(^sum)
	return answer
}