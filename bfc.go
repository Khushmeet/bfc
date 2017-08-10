package main

type Compiler struct {
	code     string
	code_len int
	position int

	instructions []*Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		code_len:     len(code),
		instructions: []*Instruction{},
	}
}

func (c *Compiler) EmitWithArgs(insType byte, count int) int {
	ins := &Instruction{
		Type: insType,
		Arg:  count,
	}

	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}

func (c *Compiler) CompileFoldableInstruction(char byte, insType byte) {
	count := 1

	for c.position < c.code_len-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.EmitWithArgs(insType, count)
}

func (c *Compiler) Compile() []*Instruction {
	for c.position < c.code_len {
		curr := c.code[c.position]

		switch curr {
		case '+':
			c.CompileFoldableInstruction('+', Plus)
		case '-':
			c.CompileFoldableInstruction('-', Minus)
		case '<':
			c.CompileFoldableInstruction('<', Left)
		case '>':
			c.CompileFoldableInstruction('>', Right)
		case '.':
			c.CompileFoldableInstruction('.', PutChar)
		case ',':
			c.CompileFoldableInstruction(',', ReadChar)
		case '[':
			c.CompileFoldableInstruction('[', JmpIfZero)
		case ']':
			c.CompileFoldableInstruction(']', JmpIfNZero)
		}

		c.position++
	}

	return c.instructions
}
