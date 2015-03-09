package gboy

// Misc Types
type Address uint16
type Word uint16

// Flags
const FZero = 0x80   // ZERO
const FOp = 0x40     // Operation
const FHCarry = 0x20 // Half-Carry
const FCarry = 0x10  // Carry

type Z80 struct {
	_m, _t int
	_r     Registers
}

// Registers
type Registers struct {
	A, B, C, D, E, H, L, F byte
	PC, SP                 uint16
	M, T                   int
}
