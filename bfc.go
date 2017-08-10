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
	stack := []int{}

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
			pos := c.EmitWithArgs(JmpIfZero, 0)
			stack = append(stack, pos)
		case ']':
			open_brac := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			close_brac := c.EmitWithArgs(JmpIfNZero, open_brac)
			c.instructions[open_brac].Arg = close_brac
		}

		c.position++
	}

	return c.instructions
}
