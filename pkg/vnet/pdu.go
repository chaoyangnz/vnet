package vnet

type ProtocolDatagramUnit interface {
	Name() string
	Payload() ProtocolDatagramUnit
	Marshal() []byte
	Print()
}

const WORDS = 4