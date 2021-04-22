package ipaddr
import (
		"strconv"
		"strings"
		"fmt"
)

type IP struct {
	octects [4]uint8
	addr uint32
	str string
}

func NewIP(ip string) (*IP, error) {
	newIP := &IP{str: ip}
	octects := strings.Split(ip, ".")
	for i, oct := range octects {
		val, err := strconv.ParseUint(oct, 10, 8); if err != nil {
			return nil, err
		}
		newIP.octects[i] = uint8(val)
	}

	var addr uint32 = 0
	var shiftby uint8 = 24
	for _, oct := range newIP.octects {
		val := uint32(oct) << shiftby
		addr += uint32(val)
		shiftby -= 8
	}
	newIP.addr = addr
	return newIP, nil
}

func NewIPfromAddr(addr uint32) *IP {
	var octects [4]uint8
	var andValue uint32 = 0xFF000000
	var shiftby uint8 = 24
	for i := 0; i < 4; i++ {
		var val uint32 = addr & andValue
		octects[i] = uint8(val >> shiftby)
		andValue >>= 8
		shiftby -= 8
	}
	newIP := &IP{addr: addr, octects: octects}
	for _, oct := range octects {
		newIP.str += fmt.Sprintf("%d.", oct)
	}
	newIP.str = newIP.str[:len(newIP.str)-1]
	return newIP
}

func (ip *IP)Next() *IP {
	return NewIPfromAddr(ip.addr + 1)
}

func (ip *IP)Equals(otherip *IP) bool {
	if ip.addr == otherip.addr {
		return true
	}
	return false
}

func (ip *IP)String() string {
	return ip.str
}


/*
func (end *IP) Minus(start *IP) int {
	shiftby, total := 0,0
	for i:=3; i >= 0; i-- {
		val := end.octects[i] - start.octects[i]
		total += (val << shiftby)
		shiftby += 8
	}
	return total
}
*/

