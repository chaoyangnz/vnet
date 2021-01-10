package vnet

import (
	"encoding/binary"
	"fmt"
)

type TCPSegment struct {
	sourcePort uint16
	destPort uint16
	seqNumber uint32
	ackNumber uint32 // if ACK set
	headerLength uint8 // 4 bits
	reserved uint8 // 3 bits
	ns	uint8 // 1 bit
	cwr uint8 // 1 bit
	ece uint8 // 1 bit
	urg uint8 // 1 bit
	ack uint8 // 1 bit
	psh uint8 // 1 bit
	rst uint8 // 1 bit
	syn uint8 // 1 bit
	fin uint8 // 1 bit
	windowSize uint16
	checksum uint16
	urgentPointer uint16 // if URG set
	options []uint8 // 20 - 40 bytes
	data []uint8
	payload ProtocolDatagramUnit
}

func NewTCPSegment(bytes []byte) *TCPSegment {
	headerLength := ((bytes[12] >> 4) & 0b1111) * WORDS
	return &TCPSegment{
		sourcePort: binary.BigEndian.Uint16(bytes[0:2]),
		destPort: binary.BigEndian.Uint16(bytes[2:4]),
		seqNumber: binary.BigEndian.Uint32(bytes[4:8]),
		ackNumber: binary.BigEndian.Uint32(bytes[8:12]),
		headerLength: headerLength,
		reserved: (bytes[12] >> 1) & 0b111,
		ns: bytes[12] & 0b1,
		cwr: (bytes[13] >> 7) & 0b1,
		ece: (bytes[13] >> 6) & 0b1,
		urg: (bytes[13] >> 5) & 0b1,
		ack: (bytes[13] >> 4) & 0b1,
		psh: (bytes[13] >> 3) & 0b1,
		rst: (bytes[13] >> 2) & 0b1,
		syn: (bytes[13] >> 1) & 0b1,
		fin: (bytes[13] >> 0) & 0b1,
		windowSize: binary.BigEndian.Uint16(bytes[14:16]),
		checksum: binary.BigEndian.Uint16(bytes[16:18]),
		urgentPointer: binary.BigEndian.Uint16(bytes[18:20]),
		options: bytes[20:headerLength],
		data: bytes[headerLength:],
	}
}

func (this *TCPSegment) Name() string {
	return "TCP"
}

func (this *TCPSegment) Payload() ProtocolDatagramUnit {
	return this.payload
}

func (this *TCPSegment) Marshal() []byte {
	return nil
}

func (this *TCPSegment) Print() {
	fmt.Printf("---- %s \n", this.Name())
	fmt.Printf("source port: %d -> dest port: %d \n", this.sourcePort, this.destPort)
	fmt.Printf("syn: %d, ack: %d, fin: %d \n", this.syn, this.ack, this.fin)
	fmt.Printf("raw payload: %s \n", string(this.data))
}

