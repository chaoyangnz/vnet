package vnet

import (
	"encoding/binary"
	"fmt"
)

type ICMPDatagram struct {
	type0 uint8
	code uint8
	checksum uint16
	identifier uint16
	sequence uint16
	data []uint8
}

func UnmarshalICMPDatagram(bytes []byte) *ICMPDatagram {
	return &ICMPDatagram{
		type0:       bytes[0],
		code:       bytes[1],
		checksum:   binary.BigEndian.Uint16(bytes[2:4]),
		identifier: binary.BigEndian.Uint16(bytes[4:6]),
		sequence:   binary.BigEndian.Uint16(bytes[6:8]),
		data:    bytes[8:],
	}
}

func (this *ICMPDatagram) Type() string  {
	switch this.type0 {
	case 0: return "echo_reply"
	case 8: return "echo_request"
	default:
		return string(this.type0)
	}
}

func (this *ICMPDatagram) Print()  {
	fmt.Printf("ICMP %v, data (%d): %v \n", this.Type(), len(this.data), forceASCII(this.data))
}

//func MarshalICMPDatagram(*ICMPDatagram) []byte {
//	return &ICMPDatagram{
//		Type:       bytes[0],
//		Code:       bytes[1],
//		Checksum:   binary.BigEndian.Uint16(bytes[2:4]),
//		Identifier: binary.BigEndian.Uint16(bytes[4:6]),
//		Sequence:   binary.BigEndian.Uint16(bytes[6:8]),
//		Payload:    bytes[8:],
//	}
//}

func forceASCII(s []byte) string {
	rs := make([]rune, 0, len(s))
	for _, r := range s {
		//if r <= 127 {
		rs = append(rs, rune(r))
		//}
	}
	return string(rs)
}