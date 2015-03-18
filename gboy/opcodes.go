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
	z._instructions[0x30] = func() { // JR NC,n
		i, err := mmu.ReadByte(z._r.PC)
		decrPC := false

		if i > 127 {
			i = ^i + 1
			decrPC = true
		}

		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		if (z._r.F & 0x10) == 0 {
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

	z._instructions[0x31] = func() { // LD SP,nn
		var err error
		z._r.SP, err = mmu.ReadWord(z._r.PC)
		z._r.PC += 2

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._instructions[0x32] = func() { // LDD (HL), A
		err := mmu.WriteByte(CombineToAddress(z._r.H, z._r.L), z._r.A)

		z._r.L--
		if z._r.L == 0 {
			z._r.H--
		}

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x33] = func() { // INC SP
		z._r.SP++

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x34] = func() { // INC (HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		i, err = mmu.ReadByte(addr + 1)

		err2 := mmu.WriteByte(addr, i)

		z._instructions.ZeroF(i, 0)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err, err2)
	}

	z._instructions[0x35] = func() { // DEC (HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		i, err = mmu.ReadByte(addr - 1)

		err2 := mmu.WriteByte(addr, i)

		z._instructions.ZeroF(i, 0)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err, err2)
	}

	z._instructions[0x36] = func() { // LD (HL),n
		vatpc, err := mmu.ReadByte(z._r.PC)
		err2 := mmu.WriteByte(CombineToAddress(z._r.H, z._r.L), vatpc)

		z._r.PC++

		z._r.M = 3
		z._r.T = 12

		LogErrors(err, err2)
	}

	z._instructions[0x37] = func() { // SCF
		z._r.F |= 0x10

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x38] = func() { // JR C,n
		i, err := mmu.ReadByte(z._r.PC)
		decrPC := false

		if i > 127 {
			i = ^i + 1
			decrPC = true
		}

		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		if (z._r.F & 0x10) == 0x10 {
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

	z._instructions[0x39] = func() { // ADD HL,SP
		hl := uint32(CombineToAddress(z._r.H, z._r.L))
		hl += uint32(z._r.SP)

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

	z._instructions[0x3A] = func() { // LDD A,(HL)
		var err error
		z._r.A, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.L--
		if z._r.L == 0 {
			z._r.H--
		}

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x3B] = func() { // DEC SP
		z._r.SP--

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x3C] = func() { // INC A
		z._r.A++

		z._instructions.ZeroF(z._r.A, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x3D] = func() { // DEC A
		z._r.A--

		z._instructions.ZeroF(z._r.A, 0)

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x3E] = func() { // LD A,n
		var err error
		z._r.A, err = mmu.ReadByte(z._r.PC)

		z._r.PC++

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x3F] = func() { // CCF
		var ci byte
		if (z._r.F & 0x10) == 0 {
			ci = 0x10
		} else {
			ci = 0
		}

		z._r.F = (z._r.F & 0xEF) + ci

		z._r.M = 1
		z._r.T = 4
	}

	// 0x40
	z._instructions[0x40]

	z._instructions[0x41]

	z._instructions[0x42]

	z._instructions[0x43]

	z._instructions[0x44]

	z._instructions[0x45]

	z._instructions[0x46]

	z._instructions[0x47]

	z._instructions[0x48]

	z._instructions[0x49]

	z._instructions[0x4A]

	z._instructions[0x4B]

	z._instructions[0x4C]

	z._instructions[0x4D]

	z._instructions[0x4E]

	z._instructions[0x4F]

	// 0x50
	z._instructions[0x50]

	z._instructions[0x51]

	z._instructions[0x52]

	z._instructions[0x53]

	z._instructions[0x54]

	z._instructions[0x55]

	z._instructions[0x56]

	z._instructions[0x57]

	z._instructions[0x58]

	z._instructions[0x59]

	z._instructions[0x5A]

	z._instructions[0x5B]

	z._instructions[0x5C]

	z._instructions[0x5D]

	z._instructions[0x5E]

	z._instructions[0x5F]

	// 0x60
	z._instructions[0x60]

	z._instructions[0x61]

	z._instructions[0x62]

	z._instructions[0x63]

	z._instructions[0x64]

	z._instructions[0x65]

	z._instructions[0x66]

	z._instructions[0x67]

	z._instructions[0x68]

	z._instructions[0x69]

	z._instructions[0x6A]

	z._instructions[0x6B]

	z._instructions[0x6C]

	z._instructions[0x6D]

	z._instructions[0x6E]

	z._instructions[0x6F]

	// 0x70
	z._instructions[0x70]

	z._instructions[0x71]

	z._instructions[0x72]

	z._instructions[0x73]

	z._instructions[0x74]

	z._instructions[0x75]

	z._instructions[0x76]

	z._instructions[0x77]

	z._instructions[0x78]

	z._instructions[0x79]

	z._instructions[0x7A]

	z._instructions[0x7B]

	z._instructions[0x7C]

	z._instructions[0x7D]

	z._instructions[0x7E]

	z._instructions[0x7F]

	// 0x80
	z._instructions[0x80]

	z._instructions[0x81]

	z._instructions[0x82]

	z._instructions[0x83]

	z._instructions[0x84]

	z._instructions[0x85]

	z._instructions[0x86]

	z._instructions[0x87]

	z._instructions[0x88]

	z._instructions[0x89]

	z._instructions[0x8A]

	z._instructions[0x8B]

	z._instructions[0x8C]

	z._instructions[0x8D]

	z._instructions[0x8E]

	z._instructions[0x8F]

	// 0x90
	z._instructions[0x90]

	z._instructions[0x91]

	z._instructions[0x92]

	z._instructions[0x93]

	z._instructions[0x94]

	z._instructions[0x95]

	z._instructions[0x96]

	z._instructions[0x97]

	z._instructions[0x98]

	z._instructions[0x99]

	z._instructions[0x9A]

	z._instructions[0x9B]

	z._instructions[0x9C]

	z._instructions[0x9D]

	z._instructions[0x9E]

	z._instructions[0x9F]

	// 0xA0
	z._instructions[0xA0]

	z._instructions[0xA1]

	z._instructions[0xA2]

	z._instructions[0xA3]

	z._instructions[0xA4]

	z._instructions[0xA5]

	z._instructions[0xA6]

	z._instructions[0xA7]

	z._instructions[0xA8]

	z._instructions[0xA9]

	z._instructions[0xAA]

	z._instructions[0xAB]

	z._instructions[0xAC]

	z._instructions[0xAD]

	z._instructions[0xAE]

	z._instructions[0xAF]

	// 0xB0
	z._instructions[0xB0]

	z._instructions[0xB1]

	z._instructions[0xB2]

	z._instructions[0xB3]

	z._instructions[0xB4]

	z._instructions[0xB5]

	z._instructions[0xB6]

	z._instructions[0xB7]

	z._instructions[0xB8]

	z._instructions[0xB9]

	z._instructions[0xBA]

	z._instructions[0xBB]

	z._instructions[0xBC]

	z._instructions[0xBD]

	z._instructions[0xBE]

	z._instructions[0xBF]

	// 0xC0
	z._instructions[0xC0]

	z._instructions[0xC1]

	z._instructions[0xC2]

	z._instructions[0xC3]

	z._instructions[0xC4]

	z._instructions[0xC5]

	z._instructions[0xC6]

	z._instructions[0xC7]

	z._instructions[0xC8]

	z._instructions[0xC9]

	z._instructions[0xCA]

	z._instructions[0xCB]

	z._instructions[0xCC]

	z._instructions[0xCD]

	z._instructions[0xCE]

	z._instructions[0xCF]

	// 0xD0
	z._instructions[0xD0]

	z._instructions[0xD1]

	z._instructions[0xD2]

	z._instructions[0xD3]

	z._instructions[0xD4]

	z._instructions[0xD5]

	z._instructions[0xD6]

	z._instructions[0xD7]

	z._instructions[0xD8]

	z._instructions[0xD9]

	z._instructions[0xDA]

	z._instructions[0xDB]

	z._instructions[0xDC]

	z._instructions[0xDD]

	z._instructions[0xDE]

	z._instructions[0xDF]

	// 0xE0
	z._instructions[0xE0]

	z._instructions[0xE1]

	z._instructions[0xE2]

	z._instructions[0xE3]

	z._instructions[0xE4]

	z._instructions[0xE5]

	z._instructions[0xE6]

	z._instructions[0xE7]

	z._instructions[0xE8]

	z._instructions[0xE9]

	z._instructions[0xEA]

	z._instructions[0xEB]

	z._instructions[0xEC]

	z._instructions[0xED]

	z._instructions[0xEE]

	z._instructions[0xEF]

	// 0xF0
	z._instructions[0xF0]

	z._instructions[0xF1]

	z._instructions[0xF2]

	z._instructions[0xF3]

	z._instructions[0xF4]

	z._instructions[0xF5]

	z._instructions[0xF6]

	z._instructions[0xF7]

	z._instructions[0xF8]

	z._instructions[0xF9]

	z._instructions[0xFA]

	z._instructions[0xFB]

	z._instructions[0xFC]

	z._instructions[0xFD]

	z._instructions[0xFE]

	z._instructions[0xFF]
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
