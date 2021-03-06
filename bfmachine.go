package main

import "io"

type BFMachine struct {
	code []*Instruction
	ip   int

	memory [30000]int
	dp     int

	input  io.Reader
	output io.Writer
	buffer []byte
}

func NewBFMachine(code []*Instruction, in io.Reader, out io.Writer) *BFMachine {
	return &BFMachine{
		code:   code,
		input:  in,
		output: out,
		buffer: make([]byte, 1),
	}
}

func (bfm *BFMachine) readChar() {
	n, err := bfm.input.Read(bfm.buffer)

	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("Wrong number of bytes read!")
	}

	bfm.memory[bfm.dp] = int(bfm.buffer[0])
}

func (bfm *BFMachine) putChar() {
	bfm.buffer[0] = byte(bfm.memory[bfm.dp])

	n, err := bfm.output.Write(bfm.buffer)

	if err != nil {
		panic(err)
	}

	if n != 1 {
		panic("Wrong number of bytes written!")
	}
}

func (bfm *BFMachine) Execute() {
	for bfm.ip < len(bfm.code) {
		inst := bfm.code[bfm.ip]

		switch inst.Type {
		case Plus:
			bfm.memory[bfm.dp] += inst.Arg
		case Minus:
			bfm.memory[bfm.dp] -= inst.Arg
		case Right:
			bfm.dp += inst.Arg
		case Left:
			bfm.dp -= inst.Arg
		case PutChar:
			for i := 0; i < inst.Arg; i++ {
				bfm.putChar()
			}
		case ReadChar:
			for i := 0; i < inst.Arg; i++ {
				bfm.readChar()
			}
		case JmpIfZero:
			if bfm.memory[bfm.dp] == 0 {
				bfm.ip = inst.Arg
				continue
			}
		case JmpIfNZero:
			if bfm.memory[bfm.dp] != 0 {
				bfm.ip = inst.Arg
				continue
			}
		}
		bfm.ip++
	}
}
