package gboy

import (
	"fmt"
)

type Instructions [0xFF]func()

/* Loads up the instructions map for the GameBoys's Z80
/ Yup all 256 of them were typed T_T
/
/ Probably this can be cleaned up and abstracted more since
/ alot of the functionality is shared between instructions. */
func (z *Z80) InitInstructions() {
	// 0x00
	z._instructions[0x00] = func() { // NOP
		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x01] = func() { // LD BC,nn
		var err, err2 error

		z._r.C, err = mmu.ReadByte(z._r.PC)
		z._r.B, err2 = mmu.ReadByte(z._r.PC + 1)
		z._r.PC += 2

		z._r.M = 3
		z._r.T = 12

		LogErrors(err, err2)
	}

	z._instructions[0x02] = func() { // LD (BC),A
		pc, err := mmu.ReadByte(z._r.PC)
		addr := CombineToAddress(z._r.B, z._r.L)
		mmu.WriteByte(addr, pc)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x03] = func() { // INC BC
		z._r.C++
		if z._r.C == 0 {
			z._r.B++
		}

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x04] = func() { // INC B
		z._r.B++
		z._instructions.ZeroF(z._r.B, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x05] = func() { // DEC B
		z._r.B--
		z._instructions.ZeroF(z._r.B, 0)

		z._r.M = 1
		z._r.M = 4
	}

	z._instructions[0x06] = func() { // LD B,n
		var err error
		z._r.B, err = mmu.ReadByte(z._r.PC)
		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x07] = func() { // RLC A
		var ci, co byte

		if (z._r.A & 0x80) != 0 {
			ci = 1
			co = 0x10
		} else {
			ci = 0
			co = 0
		}

		z._r.A = (z._r.A << 1) + ci
		z._r.F = (z._r.F & 0xEF) + co

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x08] = func() { // LD (nn),SP
	}

	z._instructions[0x09] = func() { // ADD HL,BC
		hl := uint32(CombineToAddress(z._r.H, z._r.L))
		hl += uint32(CombineToAddress(z._r.B, z._r.C))

		if hl > 0xFFFF {
			z._r.F |= 0x10
		} else {
			z._r.F &= 0xEF
		}

		z._r.H = byte(hl >> 8)
		z._r.L = byte(hl)

		z._r.M = 3
		z._r.T = 12
	}

	z._instructions[0x0A] = func() { // LD A,(BC)
		var err error
		addr := Address(z._r.B + z._r.C)
		z._r.A, err = mmu.ReadByte(addr)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x0B] = func() { // DEC BC
		z._r.C--
		if z._r.C == 0xFF {
			z._r.B--
		}

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x0C] = func() { // INC C
		z._r.C++
		z._instructions.ZeroF(z._r.C, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x0D] = func() { // DEC C
		z._r.C--
		z._instructions.ZeroF(z._r.C, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x0E] = func() { // LD C,n
		var err error
		z._r.C, err = mmu.ReadByte(z._r.PC)
		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x0F] = func() { // RRC A
		var ci, co byte

		if (z._r.A & 1) != 0 {
			ci = 0x80
			co = 0x10
		} else {
			ci = 0
			co = 0
		}

		z._r.A = (z._r.A >> 1) + ci
		z._r.F = (z._r.F & 0xEF) + co

		z._r.M = 1
		z._r.T = 4
	}

	// 0x10
	z._instructions[0x10] = func() { // STOP
		i, err := mmu.ReadByte(z._r.PC)
		decrPC := false

		if i > 127 {
			i = ^i + 1
			decrPC = true
		}

		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		if z._r.B != 0 {
			if decrPC {
				z._r.PC -= Address(i)
			} else {
				z._r.PC += Address(i)
			}
			z._r.M++
			z._r.T += 4
		}

		LogErrors(err)
	}

	z._instructions[0x11] = func() { // LD DE,nn
		var err, err2 error
		z._r.E, err = mmu.ReadByte(z._r.PC)
		z._r.D, err2 = mmu.ReadByte(z._r.PC + 1)

		z._r.PC += 2

		z._r.M = 3
		z._r.T = 12

		LogErrors(err, err2)
	}

	z._instructions[0x12] = func() { // LD (DE),A
		addr := Address((z._r.D << 8) + z._r.E)
		mmu.WriteByte(addr, z._r.A)

		z._r.M = 2
		z._r.T = 8
	}

	z._instructions[0x13] = func() { // INC DE
		z._r.E++
		if z._r.E == 0 {
			z._r.D++
		}

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x14] = func() { // INC D
		z._r.D++
		z._instructions.ZeroF(z._r.D, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x15] = func() { // DEC D
		z._r.D--
		z._instructions.ZeroF(z._r.D, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x16] = func() { // LD D,n
		var err error
		z._r.D, err = mmu.ReadByte(Address(z._r.PC))

		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x17] = func() { // RL A
		var ci, co byte

		if (z._r.F & 0x10) != 0 {
			ci = 1
		} else {
			ci = 0
		}

		if (z._r.A & 0x80) != 0 {
			co = 0x10
		} else {
			co = 0
		}

		z._r.A = (z._r.A << 1) + ci
		z._r.F = (z._r.F & 0xEF) + co

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x18] = func() { // JR n
		i, err := mmu.ReadByte(z._r.PC)
		decrPC := false

		if i > 127 {
			i = ^i + 1
			decrPC = true
		}

		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		if decrPC {
			z._r.PC -= Address(i)
		} else {
			z._r.PC += Address(i)
		}

		z._r.M++
		z._r.T += 4

		LogErrors(err)
	}

	z._instructions[0x19] = func() { // ADD HL,DE
		hl := uint32(CombineToAddress(z._r.H, z._r.L))
		hl += uint32(CombineToAddress(z._r.D, z._r.E))

		if hl > 0xFFFF {
			z._r.F |= 0x10
		} else {
			z._r.F &= 0xEF
		}

		z._r.H = byte(hl >> 8)

		z._r.M = 3
		z._r.T = 12
	}

	z._instructions[0x1A] = func() { // LD A,(DE)
		var err error
		addr := CombineToAddress(z._r.D, z._r.E)
		z._r.A, err = mmu.ReadByte(addr)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x1B] = func() { // DEC DE
		z._r.E--
		if z._r.E == 0xFF {
			z._r.B--
		}

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x1C] = func() { // INC E
		z._r.E++
		z._instructions.ZeroF(z._r.E, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x1D] = func() { // DEC E
		z._r.E--
		z._instructions.ZeroF(z._r.E, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x1E] = func() { // LD E,n
		var err error
		z._r.E, err = mmu.ReadByte(z._r.PC)
		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x1F] = func() { // RR A
		var ci, co byte

		if (z._r.F & 0x10) != 0 {
			ci = 0x80
		} else {
			ci = 0
		}

		if (z._r.A & 1) != 0 {
			co = 0x10
		} else {
			co = 0
		}

		z._r.A = (z._r.A >> 1) + ci
		z._r.F = (z._r.F & 0xEF) + co

		z._r.M = 1
		z._r.T = 4
	}

	// 0x20
	z._instructions[0x20] = func() { // JR NZ,n
		i, err := mmu.ReadByte(z._r.PC)
		decrPC := false

		if i > 127 {
			i = ^i + 1
			decrPC = true
		}

		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		if (z._r.F & 0x80) == 0 {
			if decrPC {
				z._r.PC -= Address(i)
			} else {
				z._r.PC += Address(i)
			}

			z._r.M++
			z._r.T += 4
		}

		LogErrors(err)
	}

	z._instructions[0x21] = func() { // LD HL,nn
		var err, err2 error
		z._r.L, err = mmu.ReadByte(z._r.PC)
		z._r.H, err2 = mmu.ReadByte(z._r.PC + 1)
		z._r.PC += 2

		z._r.M = 3
		z._r.T = 12

		LogErrors(err, err2)
	}

	z._instructions[0x22] = func() { // LDI (HL),A
		err := mmu.WriteByte(CombineToAddress(z._r.H, z._r.L), z._r.A)

		z._r.L++
		if z._r.L == 0 {
			z._r.H++
		}

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x23] = func() { // INC HL
		z._r.L++
		if z._r.L == 0 {
			z._r.H++
		}

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x24] = func() { // INC H
		z._r.H++
		z._instructions.ZeroF(z._r.H, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x25] = func() { // DEC H
		z._r.H--
		z._instructions.ZeroF(z._r.H, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x26] = func() { // LD H,n
		var err error
		z._r.H, err = mmu.ReadByte(z._r.PC)
		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x27] = XX // DAA

	z._instructions[0x28] = func() { // JR Z, n
		i, err := mmu.ReadByte(z._r.PC)
		decrPC := false

		if i > 127 {
			i = ^i + 1
			decrPC = true
		}

		z._r.PC++
		z._r.M = 2
		z._r.T = 8

		if (z._r.F & 0x80) == 0x80 {
			if decrPC {
				z._r.PC -= Address(i)
			} else {
				z._r.PC += Address(i)
			}

			z._r.M++
			z._r.T += 4
		}

		LogErrors(err)
	}

	z._instructions[0x29] = func() { // ADD HL,HL
		hl := uint32(CombineToAddress(z._r.H, z._r.L))
		hl += uint32(CombineToAddress(z._r.H, z._r.L))

		if hl > 0xFFFF {
			z._r.F |= 0x10
		} else {
			z._r.F &= 0xEF
		}

		z._r.H = byte((hl >> 8))
		z._r.L = byte(hl)

		z._r.M = 3
		z._r.T = 12
	}

	z._instructions[0x2A] = func() { // LDI A,(HL)
		var err error
		z._r.A, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.L++
		if z._r.L == 0 {
			z._r.H++
		}

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x2B] = func() { // DEC HL
		z._r.L--
		if z._r.L == 0xFF {
			z._r.H--
		}

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x2C] = func() { // INC L
		z._r.L++
		z._instructions.ZeroF(z._r.L, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x2D] = func() { // DEC L
		z._r.L--
		z._instructions.ZeroF(z._r.L, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x2E] = func() { // LD L,n
		var err error
		z._r.L, err = mmu.ReadByte(z._r.PC)
		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x2F] = func() { // CPL
		z._r.A = ^z._r.A
		z._instructions.ZeroF(z._r.A, 1)

		z._r.M = 1
		z._r.T = 4
	}

	// 0x30
	// 0x40
	// 0x50
	// 0x60
	// 0x70
	// 0x80
	// 0x90
	// 0xA0
	// 0xB0
	// 0xC0
	// 0xD0
	// 0xE0
	// 0xF0
}

// Helper methods
func (ins *Instructions) ZeroF(i, as byte) {
	z80._r.F = 0
	if i == 0 {
		z80._r.F |= 128
	}

	if as != 0 {
		z80._r.F |= 0x40
	} else {
		z80._r.F |= 0
	}
}

// Unimplemented instruction error!
var XX = func() {
	addr := z80._r.PC - 1
	LogErrors(InstructionNotImplementedErr(addr))
}

// Errors
type InstructionNotImplementedErr Address

func (e InstructionNotImplementedErr) Error() string {
	return fmt.Sprintf("Unimplemented instruction at %#010x, stopping", Address(e))
}
