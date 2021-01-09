package vnet

import (
	"encoding/binary"
	"fmt"
)

type IPV4Datagram struct {
	version uint8
	// 20 ~60 bytes, original 4 bit is in words
	headerLength uint8
	//1000 - minimize delay
	//0100 - maximize throughput
	//0010 - maximize reliability
	//0001 - minimize monetary cost
	typeOfService uint8
	// 20 ~ 65,536 bytes
	length uint16
	timeToLive uint8
	protocol uint8
	sourceAddr [4]uint8
	destAddr [4]uint8
	data []uint8
}

func NewIPV4Datagram(bytes []byte) *IPV4Datagram {
	headerLength := (bytes[0] & 0x0F) * 4
	length := binary.BigEndian.Uint16(bytes[2:4])
	data := bytes[headerLength:length]
	return &IPV4Datagram{
		version: bytes[0] >> 4,
		headerLength: headerLength,
		typeOfService: bytes[1],
		length: length,
		timeToLive: bytes[8],
		protocol: bytes[9],
		sourceAddr: [4]uint8{bytes[12], bytes[13], bytes[14], bytes[15]},
		destAddr: [4]uint8{bytes[16], bytes[17], bytes[18], bytes[19]},
		data: data,
	}
}

// https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
func (this IPV4Datagram) Protocol() string  {
	switch this.protocol {
	case 1: return "ICMP"
	case 6: return "TCP"
	case 17: return "UDP"
	default:
		return string(this.protocol)
	}
}

func (this IPV4Datagram) Print() {
	fmt.Printf("version: %v ", this.version)
	fmt.Printf("length: 0 ~ %d ~ %v ", this.headerLength, this.length)
	fmt.Printf("ip: %d.%d.%d.%d -> %d.%d.%d.%d \n\n",
		this.sourceAddr[0], this.sourceAddr[1], this.sourceAddr[2], this.sourceAddr[3],
		this.destAddr[0], this.destAddr[1], this.destAddr[2], this.destAddr[3])
	fmt.Printf("ttl: %v protocol: %v \n", this.timeToLive, this.Protocol())
	fmt.Printf("data: %v \n", this.data)

	if this.Protocol() == "ICMP" {
		icmp := UnmarshalICMPDatagram(this.data)
		icmp.Print()
	}
}

