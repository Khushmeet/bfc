package main

import "io"

type BFMachine struct {
	code   string
	ip     int
	memory [30000]int
	dp     int
	input  io.Reader
	output io.Writer
	buffer []byte
}

func NewBFMachine(code string, in io.Reader, out io.Writer) *BFMachine {
	return &BFMachine{
		code: code,
		input: in,
		output: out,
		buffer: make([]byte, 1),
	}
}

func (bfm *BFMachine) readChar() {
	n, err := bfm.input.Read(bfm.buffer)

	if err!=nil {
		panic(err)
	}

	if n!=1 {
		panic("Wrong number of bytes read!")
	}

	bfm.memory[bfm.dp] = int(bfm.buffer[0])
}

func (bfm *BFMachine) putChar() {
	bfm.buffer[0] = byte(bfm.memory[bfm.dp])

	n, err := bfm.output.Write(bfm.buffer)

	if err!=nil {
		panic(err)
	}

	if n!=1 {
		panic("Wrong number of bytes written!")
	}
}

func (bfm *BFMachine) Execute() {
	for bfm.ip < len(bfm.code) {
		instruction := bfm.code[bfm.ip]

		switch instruction {
		case '+': bfm.memory[bfm.dp]++
		case '-': bfm.memory[bfm.dp]--
		case '>': bfm.dp++
		case '<': bfm.dp--
		case '.': bfm.putChar()
		case ',': bfm.readChar()
		case '[':
			if bfm.memory[bfm.dp] == 0 {
				depth := 1
				for depth != 0 {
					bfm.ip++
					switch bfm.code[bfm.ip] {
					case '[':
						depth++
					case ']':
						depth--
					}
				}
			}
		case ']':
			if bfm.memory[bfm.dp] != 0 {
				depth := 1
				for depth != 0 {
					bfm.ip--
					switch bfm.code[bfm.ip] {
					case '[':
						depth--
					case ']':
						depth++
					}
				}
			}
		}
		bfm.ip++
	}
}


