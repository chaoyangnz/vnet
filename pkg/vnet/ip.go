package vnet

import (
	"encoding/binary"
	"fmt"
)

type IPProtocol uint8
// https://en.wikipedia.org/wiki/List_of_IP_protocol_numbers
const (
	IP_ICMP IPProtocol = 1
	IP_TCP  IPProtocol = 6
	IP_UDP  IPProtocol = 17
)

func (this IPProtocol) String() string  {
	switch this {
	case IP_ICMP: return "ICMP"
	case IP_TCP: return "TCP"
	case IP_UDP: return "UDP"
	default:
		fmt.Printf("Unsupported payload protocol: %d", this)
		return "UNKNOWN"
	}
}

type IPv4Datagram struct {
	version uint8
	// 20 ~60 bytes, original 4 bit is in words
	headerLength uint8
	//1000 - minimize delay
	//0100 - maximize throughput
	//0010 - maximize reliability
	//0001 - minimize monetary cost
	serviceType uint8
	// 20 ~ 65,536 bytes
	length     uint16
	identification uint16
	flags uint8 // 3 bits
	fragmentOffset uint16 // 13 bits
	ttl        uint8
	protocol   IPProtocol
	sourceAddr [4]uint8
	destAddr   [4]uint8
	options    []uint8
	data       []uint8
	payload    ProtocolDatagramUnit
}

func EmptyIPv4Datagram() *IPv4Datagram {
	return &IPv4Datagram{}
}

func NewIPv4Datagram(bytes []byte) *IPv4Datagram {
	headerLength := (bytes[0] & 0x0F) * WORDS
	length := binary.BigEndian.Uint16(bytes[2:4])
	protocol := IPProtocol(bytes[9])
	data := bytes[headerLength:length]
	var payload ProtocolDatagramUnit
	if len(data) != 0 {
		switch protocol {
		case IP_ICMP:
			payload = NewICMPDatagram(data)
			break
		case IP_TCP:
			payload = NewTCPSegment(data)
			break
		case IP_UDP:
		default:
			fmt.Printf("Unsupported payload protocol: %d", protocol)
		}
	}

	return &IPv4Datagram{
		version:      bytes[0] >> 4,
		headerLength: headerLength,
		serviceType:  bytes[1],
		length:       length,
		identification: binary.BigEndian.Uint16(bytes[4:6]),
		flags: bytes[6] >> 5,
		fragmentOffset: binary.BigEndian.Uint16(bytes[6:8]) & 0x1FFF,
		ttl:          bytes[8],
		protocol:     protocol,
		sourceAddr:   [4]uint8{bytes[12], bytes[13], bytes[14], bytes[15]},
		destAddr:     [4]uint8{bytes[16], bytes[17], bytes[18], bytes[19]},
		options: bytes[20:headerLength],
		data:         data,
		payload:      payload,
	}
}

func (this *IPv4Datagram) Marshal() []byte {
	return nil
}

func (this *IPv4Datagram) Name() string {
	return "IPV4"
}

func (this IPv4Datagram) Payload() ProtocolDatagramUnit {
	return this.payload
}

func (this IPv4Datagram) Print() {
	fmt.Printf("==== IP version: %v ", this.version)
	fmt.Printf("length: 0 ~ %d ~ %v ", this.headerLength, this.length)
	fmt.Printf("ip: %d.%d.%d.%d -> %d.%d.%d.%d \n",
		this.sourceAddr[0], this.sourceAddr[1], this.sourceAddr[2], this.sourceAddr[3],
		this.destAddr[0], this.destAddr[1], this.destAddr[2], this.destAddr[3])
	fmt.Printf("ttl: %v protocol: %v \n", this.ttl, this.protocol.String())
	fmt.Printf("data: %v \n", this.data)

	if this.payload != nil {
		this.payload.Print()
	}
	fmt.Println()
}

