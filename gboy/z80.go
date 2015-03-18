package gboy

// Misc Types
type Address uint16

// Flags
const FZero = 0x80   // ZERO
const FOp = 0x40     // Operation
const FHCarry = 0x20 // Half-Carry
const FCarry = 0x10  // Carry

type Z80 struct {
	_clock        Clock
	_r            Registers
	_instructions Instructions
}

// Clock
type Clock struct {
	M, T int
}

// Registers
type Registers struct {
	A, B, C, D, E, H, L, F byte
	PC, SP                 Address
	M, T                   int
}

// Reset the CPU by clearing the registers and reseting the PC and SP
func (z *Z80) Reset() {
	z._r.A = 0
	z._r.B = 0
	z._r.C = 0
	z._r.D = 0
	z._r.E = 0
	z._r.H = 0
	z._r.L = 0
	z._r.F = 0

	z._r.M = 0
	z._r.T = 0

	// Reset execution
	z._r.PC = 0
	z._r.SP = 0

	// Reset the clock
	z._clock.M = 0
	z._clock.T = 0
}

func (z *Z80) Dispatch() {
	z.InitInstructions()
	go z.Execute()
}

func (z *Z80) Execute() {
	for {
		op, err := mmu.ReadByte(z._r.PC) // Fetch the instruction
		z._r.PC++                        // Increment the PC
		z._instructions[Address(op)]()
		z._r.PC &= 0xFFFF // Mask the PC to 16 bits

		// Update the CPU time
		z._clock.M += z._r.M
		z._clock.T += z._r.T

		LogErrors(err)
	}
}
