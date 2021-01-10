package vnet

import (
	"encoding/binary"
	"fmt"
)

type ICMPType uint8

const (
	ICMP_ECHO_REQUEST ICMPType = 8
	ICMP_ECHO_REPLY ICMPType = 0
)

func (this ICMPType) String() string {
	switch this {
	case ICMP_ECHO_REQUEST: return "echo_reply"
	case ICMP_ECHO_REPLY: return "echo_request"
	default:
		return string(this)
	}
}

type ICMPDatagram struct {
	type0 ICMPType
	code uint8
	checksum uint16
	identifier uint16
	sequence uint16
	data []uint8
	payload ProtocolDatagramUnit
}

func EmptyICMPDatagram() *ICMPDatagram {
	return &ICMPDatagram{}
}

func NewICMPDatagram(bytes []byte) *ICMPDatagram {
	return &ICMPDatagram{
		type0:      ICMPType(bytes[0]),
		code:       bytes[1],
		checksum:   binary.BigEndian.Uint16(bytes[2:4]),
		identifier: binary.BigEndian.Uint16(bytes[4:6]),
		sequence:   binary.BigEndian.Uint16(bytes[6:8]),
		data:    bytes[8:],
	}
}

func (this *ICMPDatagram) Name() string {
	return "ICMP"
}

func (this *ICMPDatagram) Marshal() []byte {
	return nil
}

func (this *ICMPDatagram) Payload() ProtocolDatagramUnit {
	return this.payload
}

func (this *ICMPDatagram) Print()  {
	fmt.Printf("ICMP %v, data (%d): %v \n", this.type0.String(), len(this.data), forceASCII(this.data))
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