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
	z._instructions[0x40] = func() { // LD B,B
		z._instructions.LoadRnR(&z._r.B, &z._r.B)
	}

	z._instructions[0x41] = func() { // LD B,C
		z._instructions.LoadRnR(&z._r.B, &z._r.C)
	}

	z._instructions[0x42] = func() { // LD B,D
		z._instructions.LoadRnR(&z._r.B, &z._r.D)
	}

	z._instructions[0x43] = func() { // LD B,E
		z._instructions.LoadRnR(&z._r.B, &z._r.E)
	}

	z._instructions[0x44] = func() { // LD B,H
		z._instructions.LoadRnR(&z._r.B, &z._r.H)
	}

	z._instructions[0x45] = func() { // LD B,L
		z._instructions.LoadRnR(&z._r.B, &z._r.L)
	}

	z._instructions[0x46] = func() { // LD B,(HL)
		var err error
		z._r.B, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x47] = func() { // LD B, A
		z._instructions.LoadRnR(&z._r.B, &z._r.A)
	}

	z._instructions[0x48] = func() { // LD C,B
		z._instructions.LoadRnR(&z._r.C, &z._r.B)
	}

	z._instructions[0x49] = func() { // LD C,C
		z._instructions.LoadRnR(&z._r.C, &z._r.C)
	}

	z._instructions[0x4A] = func() { // LD C,D
		z._instructions.LoadRnR(&z._r.C, &z._r.D)
	}

	z._instructions[0x4B] = func() { // LD C,E
		z._instructions.LoadRnR(&z._r.C, &z._r.E)
	}

	z._instructions[0x4C] = func() { // LD C,H
		z._instructions.LoadRnR(&z._r.C, &z._r.H)
	}

	z._instructions[0x4D] = func() { // LD C,L
		z._instructions.LoadRnR(&z._r.C, &z._r.L)
	}

	z._instructions[0x4E] = func() { // LD C,(HL)
		var err error
		z._r.C, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x4F] = func() { // LD C,A
		z._instructions.LoadRnR(&z._r.C, &z._r.A)
	}

	// 0x50
	z._instructions[0x50] = func() { // LD D,B
		z._instructions.LoadRnR(&z._r.D, &z._r.B)
	}

	z._instructions[0x51] = func() { // LD D,C
		z._instructions.LoadRnR(&z._r.D, &z._r.C)
	}

	z._instructions[0x52] = func() { // LD D,D
		z._instructions.LoadRnR(&z._r.D, &z._r.D)
	}

	z._instructions[0x53] = func() { // LD D,E
		z._instructions.LoadRnR(&z._r.D, &z._r.E)
	}

	z._instructions[0x54] = func() { // LD D,H
		z._instructions.LoadRnR(&z._r.D, &z._r.H)
	}

	z._instructions[0x55] = func() { // LD D,L
		z._instructions.LoadRnR(&z._r.D, &z._r.L)
	}

	z._instructions[0x56] = func() { // LD D,(HL)
		var err error
		z._r.D, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x57] = func() { // LD D,A
		z._instructions.LoadRnR(&z._r.D, &z._r.A)
	}

	z._instructions[0x58] = func() { // LD E,B
		z._instructions.LoadRnR(&z._r.E, &z._r.B)
	}

	z._instructions[0x59] = func() { // LD E,C
		z._instructions.LoadRnR(&z._r.E, &z._r.C)
	}

	z._instructions[0x5A] = func() { // LD E,D
		z._instructions.LoadRnR(&z._r.E, &z._r.D)
	}

	z._instructions[0x5B] = func() { // LD E,E
		z._instructions.LoadRnR(&z._r.E, &z._r.E)
	}

	z._instructions[0x5C] = func() { // LD E,H
		z._instructions.LoadRnR(&z._r.E, &z._r.H)
	}

	z._instructions[0x5D] = func() { // LD E,L
		z._instructions.LoadRnR(&z._r.E, &z._r.L)
	}

	z._instructions[0x5E] = func() { // LD E,(HL)
		var err error
		z._r.E, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x5F] = func() { // LD E,A
		z._instructions.LoadRnR(&z._r.E, &z._r.A)
	}

	// 0x60
	z._instructions[0x60] = func() { // LD H,B
		z._instructions.LoadRnR(&z._r.H, &z._r.B)
	}

	z._instructions[0x61] = func() { // LD H,C
		z._instructions.LoadRnR(&z._r.H, &z._r.C)
	}

	z._instructions[0x62] = func() { // LD H,D
		z._instructions.LoadRnR(&z._r.H, &z._r.D)
	}

	z._instructions[0x63] = func() { // LD H,E
		z._instructions.LoadRnR(&z._r.H, &z._r.E)
	}

	z._instructions[0x64] = func() { // LD H,H
		z._instructions.LoadRnR(&z._r.H, &z._r.H)
	}

	z._instructions[0x65] = func() { // LD H,L
		z._instructions.LoadRnR(&z._r.H, &z._r.L)
	}

	z._instructions[0x66] = func() { // LD H,(HL)
		var err error
		z._r.H, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x67] = func() { // LD H,A
		z._instructions.LoadRnR(&z._r.H, &z._r.A)
	}

	z._instructions[0x68] = func() { // LD L,B
		z._instructions.LoadRnR(&z._r.L, &z._r.B)
	}

	z._instructions[0x69] = func() { // LD L,C
		z._instructions.LoadRnR(&z._r.L, &z._r.C)
	}

	z._instructions[0x6A] = func() { // LD L,D
		z._instructions.LoadRnR(&z._r.L, &z._r.D)
	}

	z._instructions[0x6B] = func() { // LD L,E
		z._instructions.LoadRnR(&z._r.L, &z._r.E)
	}

	z._instructions[0x6C] = func() { // LD L,H
		z._instructions.LoadRnR(&z._r.L, &z._r.H)
	}

	z._instructions[0x6D] = func() { // LD L,L
		z._instructions.LoadRnR(&z._r.L, &z._r.L)
	}

	z._instructions[0x6E] = func() { // LD L,(HL)
		var err error
		z._r.L, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x6F] = func() { // LD L,A
		z._instructions.LoadRnR(&z._r.L, &z._r.A)
	}

	// 0x70
	z._instructions[0x70] = func() { // LD (HL),B
		addr := CombineToAddress(z._r.H, z._r.L)
		err := mmu.WriteByte(addr, z._r.B)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x71] = func() { // LD (HL),C
		addr := CombineToAddress(z._r.H, z._r.L)
		err := mmu.WriteByte(addr, z._r.C)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x72] = func() { // LD (HL),D
		addr := CombineToAddress(z._r.H, z._r.L)
		err := mmu.WriteByte(addr, z._r.D)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x73] = func() { // LD (HL),E
		addr := CombineToAddress(z._r.H, z._r.L)
		err := mmu.WriteByte(addr, z._r.E)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x74] = func() { // LD (HL),H
		addr := CombineToAddress(z._r.H, z._r.L)
		err := mmu.WriteByte(addr, z._r.H)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x75] = func() { // LD (HL), L
		addr := CombineToAddress(z._r.H, z._r.L)
		err := mmu.WriteByte(addr, z._r.L)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x76] = func() { // HALT
		z._halt = true

		z._r.M = 1
		z._r.T = 4
	}

	z._instructions[0x77] = func() { // LD (HL),A
		addr := CombineToAddress(z._r.H, z._r.L)
		err := mmu.WriteByte(addr, z._r.A)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x78] = func() { // LD A,B
		z._instructions.LoadRnR(&z._r.A, &z._r.B)
	}

	z._instructions[0x79] = func() { // LD A,C
		z._instructions.LoadRnR(&z._r.A, &z._r.C)
	}

	z._instructions[0x7A] = func() { // LD A,D
		z._instructions.LoadRnR(&z._r.A, &z._r.D)
	}

	z._instructions[0x7B] = func() { // LD A,E
		z._instructions.LoadRnR(&z._r.A, &z._r.E)
	}

	z._instructions[0x7C] = func() { // LD A,H
		z._instructions.LoadRnR(&z._r.A, &z._r.H)
	}

	z._instructions[0x7D] = func() { // LD A,L
		z._instructions.LoadRnR(&z._r.A, &z._r.L)
	}

	z._instructions[0x7E] = func() { // LD A,(HL)
		var err error
		z._r.A, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x7F] = func() { // LD A,A
		z._instructions.LoadRnR(&z._r.A, &z._r.A)
	}

	// 0x80
	z._instructions[0x80] = func() { // ADD A,B
		z._instructions.Addr_r(&z._r.B)
	}

	z._instructions[0x81] = func() { // ADD A,C
		z._instructions.Addr_r(&z._r.C)
	}

	z._instructions[0x82] = func() { // ADD A,D
		z._instructions.Addr_r(&z._r.D)
	}

	z._instructions[0x83] = func() { // ADD A,E
		z._instructions.Addr_r(&z._r.E)
	}

	z._instructions[0x84] = func() { // ADD A,H
		z._instructions.Addr_r(&z._r.H)
	}

	z._instructions[0x85] = func() { // ADD A,L
		z._instructions.Addr_r(&z._r.L)
	}

	z._instructions[0x86] = func() { // ADD A,(HL)
		hl, err := mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Addr_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x87] = func() { // ADD A,A
		z._instructions.Addr_r(&z._r.A)
	}

	z._instructions[0x88] = func() { // ADC A,B
		z._instructions.Addc_r(&z._r.B)
	}

	z._instructions[0x89] = func() { // ADC A,C
		z._instructions.Addc_r(&z._r.C)
	}

	z._instructions[0x8A] = func() { // ADC A,D
		z._instructions.Addc_r(&z._r.D)
	}

	z._instructions[0x8B] = func() { // ADC A,E
		z._instructions.Addc_r(&z._r.E)
	}

	z._instructions[0x8C] = func() { // ADC A,H
		z._instructions.Addc_r(&z._r.H)
	}

	z._instructions[0x8D] = func() { // ADC A,L
		z._instructions.Addc_r(&z._r.L)
	}

	z._instructions[0x8E] = func() { // ADC A,(HL)
		hl, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Addc_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x8F] = func() { // ADC A,A
		z._instructions.Addc_r(&z._r.A)
	}

	// 0x90
	z._instructions[0x90] = func() { // SUB A,B
		z._instructions.Subr_r(&z._r.B)
	}

	z._instructions[0x91] = func() { // SUB A,C
		z._instructions.Subr_r(&z._r.C)
	}

	z._instructions[0x92] = func() { // SUB A,D
		z._instructions.Subr_r(&z._r.D)
	}

	z._instructions[0x93] = func() { // SUB A,E
		z._instructions.Subr_r(&z._r.E)
	}

	z._instructions[0x94] = func() { // SUB A,H
		z._instructions.Subr_r(&z._r.H)
	}

	z._instructions[0x95] = func() { // SUB A,L
		z._instructions.Subr_r(&z._r.L)
	}

	z._instructions[0x96] = func() { // SUB A,(HL)
		hl, err := mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Subr_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x97] = func() { // SUB A,A
		z._instructions.Subr_r(&z._r.A)
	}

	z._instructions[0x98] = func() { // SBC A,B
		z._instructions.Subc_r(&z._r.B)
	}

	z._instructions[0x99] = func() { // SBC A,C
		z._instructions.Subc_r(&z._r.C)
	}

	z._instructions[0x9A] = func() { // SBC A,D
		z._instructions.Subc_r(&z._r.D)
	}

	z._instructions[0x9B] = func() { // SBC A,E
		z._instructions.Subc_r(&z._r.E)
	}

	z._instructions[0x9C] = func() { // SBC A,H
		z._instructions.Subc_r(&z._r.H)
	}

	z._instructions[0x9D] = func() { // SBC A,L
		z._instructions.Subc_r(&z._r.L)
	}

	z._instructions[0x9E] = func() { // SBC A,(HL)
		hl, err := mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Subc_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0x9F] = func() { // SBC A,A
		z._instructions.Subc_r(&z._r.A)
	}

	// 0xA0
	z._instructions[0xA0] = func() { // AND B
		z._instructions.Andr_r(&z._r.B)
	}

	z._instructions[0xA1] = func() { // AND C
		z._instructions.Andr_r(&z._r.C)
	}

	z._instructions[0xA2] = func() { // AND D
		z._instructions.Andr_r(&z._r.D)
	}

	z._instructions[0xA3] = func() { // AND E
		z._instructions.Andr_r(&z._r.E)
	}

	z._instructions[0xA4] = func() { // AND H
		z._instructions.Andr_r(&z._r.H)
	}

	z._instructions[0xA5] = func() { // AND L
		z._instructions.Andr_r(&z._r.L)
	}

	z._instructions[0xA6] = func() { // AND (HL)
		hl, err := mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Andr_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0xA7] = func() { // AND A
		z._instructions.Andr_r(&z._r.A)
	}

	z._instructions[0xA8] = func() { // XOR B
		z._instructions.Xorr_r(&z._r.B)
	}

	z._instructions[0xA9] = func() { // XOR C
		z._instructions.Xorr_r(&z._r.C)
	}

	z._instructions[0xAA] = func() { // XOR D
		z._instructions.Xorr_r(&z._r.D)
	}

	z._instructions[0xAB] = func() { // XOR E
		z._instructions.Xorr_r(&z._r.E)
	}

	z._instructions[0xAC] = func() { // XOR H
		z._instructions.Xorr_r(&z._r.H)
	}

	z._instructions[0xAD] = func() { // XOR L
		z._instructions.Xorr_r(&z._r.L)
	}

	z._instructions[0xAE] = func() { // XOR (HL)
		hl, err := mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Xorr_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0xAF] = func() { // XOR A
		z._instructions.Xorr_r(&z._r.A)
	}

	// 0xB0
	z._instructions[0xB0] = func() { // OR B
		z._instructions.Orr_r(&z._r.B)
	}

	z._instructions[0xB1] = func() { // OR C
		z._instructions.Orr_r(&z._r.C)
	}

	z._instructions[0xB2] = func() { // OR D
		z._instructions.Orr_r(&z._r.D)
	}

	z._instructions[0xB3] = func() { // OR E
		z._instructions.Orr_r(&z._r.E)
	}

	z._instructions[0xB4] = func() { // OR H
		z._instructions.Orr_r(&z._r.H)
	}

	z._instructions[0xB5] = func() { // OR L
		z._instructions.Orr_r(&z._r.L)
	}

	z._instructions[0xB6] = func() { // OR (HL)
		hl, err = mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Orr_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0xB7] = func() { // OR A
		z._instructions.Orr_r(&z._r.A)
	}

	z._instructions[0xB8] = func() { // CP B
		z._instructions.Cpr_r(&z._r.B)
	}

	z._instructions[0xB9] = func() { // CP C
		z._instructions.Cpr_r(&z._r.C)
	}

	z._instructions[0xBA] = func() { // CP D
		z._instructions.Cpr_r(&z._r.D)
	}

	z._instructions[0xBB] = func() { // CP E
		z._instructions.Cpr_r(&z._r.E)
	}

	z._instructions[0xBC] = func() { // CP H
		z._instructions.Cpr_r(&z._r.H)
	}

	z._instructions[0xBD] = func() { // CP L
		z._instructions.Cpr_r(&z._r.L)
	}

	z._instructions[0xBE] = func() { // CP (HL)
		hl, err := mmu.ReadByte(CombineToAddress(z._r.H, z._r.L))
		z._instructions.Cpr_r(&hl)

		z._r.M = 2
		z._r.T = 8

		LogErrors(err)
	}

	z._instructions[0xBF] = func() { // CP A
		z._instructions.Cpr_r(&z._r.A)
	}

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
// Check if a carry needs to be set and set it
func (ins *Instructions) ZeroF(i, as byte) {
	z80._r.F = 0
	if i == 0 {
		z80._r.F |= 0x80
	}

	if as != 0 {
		z80._r.F |= 0x40
	} else {
		z80._r.F |= 0
	}
}

// Loads a register at src in register at dst
func (ins *Instructions) LoadRnR(dst, src *byte) {
	*dst = *src

	z80._r.M = 1
	z80._r.T = 4
}

// Adds register at src to register A
func (ins *Instructions) Addr_r(src *byte) {
	sum := Address(z80._r.A) + Address(*src)

	z80._r.A = byte(sum)
	ins.ZeroF(z80._r.A, 0)

	if sum > 0xF {
		z80._r.F |= 0x10
	}

	z80._r.M = 1
	z80._r.T = 4
}

// Adds register at src to register A with carry
func (ins *Instructions) Addc_r(src *byte) {
	sum := Address(z80._r.A) + Address(*src)
	if (z80._r.F & 0x10) != 0 {
		sum += 1
	}

	z80._r.A = byte(sum)
	ins.ZeroF(z80._r.A, 0)

	if sum > 0xF {
		z80._r.F |= 0x10
	}

	z._r.M = 1
	z._r.T = 4
}

// Subtracts register at src from register A
func (ins *Instructions) Subr_r(src *byte) {
	sub := int16(z80._r.A) - int16(*src)

	z80._r.A = byte(sub)
	ins.ZeroF(z80._r.A, 1)

	if sub < 0 {
		z80._r.F |= 0x10
	}

	z80._r.M = 1
	z80._r.T = 4
}

// Subtracts register at src from register A with carry
func (ins *Instructions) Subc_r(src *byte) {
	sub := int16(z80._r.A) - int16(*src)
	if (z80._r.F & 0x10) != 0 {
		sub -= 1
	}

	z80._r.A = byte(sub)
	ins.ZeroF(z80._r.A, 1)

	if sub < 0 {
		z80._r.F |= 0x10
	}

	z80._r.M = 1
	z80._r.T = 4
}

// Logical Operations
// AND
func (ins *Instructions) Andr_r(src *byte) {
	z80._r.A &= *src
	ins.ZeroF(z80._r.A, 0)

	z80._r.M = 1
	z80._r.T = 4
}

// XOR
func (inc *Instructions) Xorr_r(src *byte) {
	z80._r.A ^= *src
	ins.ZeroF(z80._r.A, 0)

	z80._r.M = 1
	z80._r.T = 4
}

// OR
func (inc *Instructions) Orr_r(src *byte) {
	z80._r.A |= *src
	ins.ZeroF(z80._r.A, 0)

	z80._r.M = 1
	z80._r.T = 4
}

// Compare register at src to register A
func (ins *Instructions) Cpr_r(src *byte) {
	i := int16(z80._r.A)
	i -= int16(*src)

	ins.ZeroF(i, 1)
	if i < 0 {
		z80._r.F |= 0x10
	}

	z80._r.M = 1
	z80._r.T = 4
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
