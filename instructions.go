package main

const (
	Plus       byte = '+'
	Minus      byte = '-'
	Right      byte = '>'
	Left       byte = '<'
	PutChar    byte = '.'
	ReadChar   byte = ','
	JmpIfZero  byte = '['
	JmpIfNZero byte = ']'
)

type Instruction struct {
	Type byte
	Arg  int
}
