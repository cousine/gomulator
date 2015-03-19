package gboy

func (z *Z80) InitCBMap() {
	// CB00
	z._cbprefix[0x00] = func() { // RLC B
		z._cbprefix.RLCr_r(&z._r.B)
	}

	z._cbprefix[0x01] = func() { // RLC C
		z._cbprefix.RLCr_r(&z._r.C)
	}

	z._cbprefix[0x02] = func() { // RLC D
		z._cbprefix.RLCr_r(&z._r.D)
	}

	z._cbprefix[0x03] = func() { // RLC E
		z._cbprefix.RLCr_r(&z._r.E)
	}

	z._cbprefix[0x04] = func() { // RLC H
		z._cbprefix.RLCr_r(&z._r.H)
	}

	z._cbprefix[0x05] = func() { // RLC L
		z._cbprefix.RLCr_r(&z._r.L)
	}

	z._cbprefix[0x06] = func() { // RLC (HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.RLCr_r(&hl)

		err2 := mmu.WriteByte(addr, hl)

		z._r.M = 4
		z._r.T = 16

		LogErrors(err, err2)
	}

	z._cbprefix[0x07] = func() { // RLC A
		z._cbprefix.RLCr_r(&z._r.A)
	}

	z._cbprefix[0x08] = func() { // RRC B
		z._cbprefix.RRCr_r(&z._r.B)
	}

	z._cbprefix[0x09] = func() { // RRC C
		z._cbprefix.RRCr_r(&z._r.C)
	}

	z._cbprefix[0x0A] = func() { // RRC D
		z._cbprefix.RRCr_r(&z._r.D)
	}

	z._cbprefix[0x0B] = func() { // RRC E
		z._cbprefix.RRCr_r(&z._r.E)
	}

	z._cbprefix[0x0C] = func() { // RRC H
		z._cbprefix.RRCr_r(&z._r.H)
	}

	z._cbprefix[0x0D] = func() { // RRC L
		z._cbprefix.RRCr_r(&z._r.L)
	}

	z._cbprefix[0x0E] = func() { // RRC (HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.RRCr_r(&hl)

		err2 := mmu.WriteByte(addr, hl)

		z._r.M = 4
		z._r.T = 16

		LogErrors(err, err2)
	}

	z._cbprefix[0x0F] = func() { // RRC A
		z._cbprefix.RRCr_r(&z._r.A)
	}

	// CB10
	z._cbprefix[0x10] = func() { // RL B
		z._cbprefix.RLr_r(&z._r.B)
	}

	z._cbprefix[0x11] = func() { // RL C
		z._cbprefix.RLr_r(&z._r.C)
	}

	z._cbprefix[0x12] = func() { // RL D
		z._cbprefix.RLr_r(&z._r.D)
	}

	z._cbprefix[0x13] = func() { // RL E
		z._cbprefix.RLr_r(&z._r.E)
	}

	z._cbprefix[0x14] = func() { // RL H
		z._cbprefix.RLr_r(&z._r.H)
	}

	z._cbprefix[0x15] = func() { // RL L
		z._cbprefix.RLr_r(&z._r.L)
	}

	z._cbprefix[0x16] = func() { // RL (HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.RLr_r(&hl)

		err2 := mmu.WriteByte(addr, hl)

		z._r.M = 4
		z._r.T = 16

		LogErrors(err, err2)
	}

	z._cbprefix[0x17] = func() { // RL A
		z._cbprefix.RLr_r(&z._r.A)
	}

	z._cbprefix[0x18] = func() { // RR B
		z._cbprefix.RRr_r(&z._r.B)
	}

	z._cbprefix[0x19] = func() { // RR C
		z._cbprefix.RRr_r(&z._r.C)
	}

	z._cbprefix[0x1A] = func() { // RR D
		z._cbprefix.RRr_r(&z._r.D)
	}

	z._cbprefix[0x1B] = func() { // RR E
		z._cbprefix.RRr_r(&z._r.E)
	}

	z._cbprefix[0x1C] = func() { // RR H
		z._cbprefix.RRr_r(&z._r.H)
	}

	z._cbprefix[0x1D] = func() { // RR L
		z._cbprefix.RRr_r(&z._r.L)
	}

	z._cbprefix[0x1E] = func() { // RR (HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.RRr_r(&hl)

		err2 := mmu.WriteByte(addr, hl)

		z._r.M = 4
		z._r.T = 16

		LogErrors(err, err2)
	}

	z._cbprefix[0x1F] = func() { // RR A
		z._cbprefix.RRr_r(&z._r.A)
	}

	// CB20
	z._cbprefix[0x20] = func() { // SLA B
		z._cbprefix.SLAr_r(&z._r.B)
	}

	z._cbprefix[0x21] = func() { // SLA C
		z._cbprefix.SLAr_r(&z._r.C)
	}

	z._cbprefix[0x22] = func() { // SLA D
		z._cbprefix.SLAr_r(&z._r.D)
	}

	z._cbprefix[0x23] = func() { // SLA E
		z._cbprefix.SLAr_r(&z._r.E)
	}

	z._cbprefix[0x24] = func() { // SLA H
		z._cbprefix.SLAr_r(&z._r.H)
	}

	z._cbprefix[0x25] = func() { // SLA L
		z._cbprefix.SLAr_r(&z._r.L)
	}

	z._cbprefix[0x26] = XX

	z._cbprefix[0x27] = func() { // SLA A
		z._cbprefix.SLAr_r(&z._r.A)
	}

	z._cbprefix[0x28] = func() { // SRA B
		z._cbprefix.SRAr_r(&z._r.B)
	}

	z._cbprefix[0x29] = func() { // SRA C
		z._cbprefix.SRAr_r(&z._r.C)
	}

	z._cbprefix[0x2A] = func() { // SRA D
		z._cbprefix.SRAr_r(&z._r.D)
	}

	z._cbprefix[0x2B] = func() { // SRA E
		z._cbprefix.SRAr_r(&z._r.E)
	}

	z._cbprefix[0x2C] = func() { // SRA H
		z._cbprefix.SRAr_r(&z._r.H)
	}

	z._cbprefix[0x2D] = func() { // SRA L
		z._cbprefix.SRAr_r(&z._r.L)
	}

	z._cbprefix[0x2E] = XX

	z._cbprefix[0x2F] = func() { // SRA A
		z._cbprefix.SRAr_r(&z._r.A)
	}

	// CB30
	z._cbprefix[0x30] = func() { // SWAP B
		z._cbprefix.SWAPr_r(&z._r.B)
	}

	z._cbprefix[0x31] = func() { // SWAP C
		z._cbprefix.SWAPr_r(&z._r.C)
	}

	z._cbprefix[0x32] = func() { // SWAP D
		z._cbprefix.SWAPr_r(&z._r.D)
	}

	z._cbprefix[0x33] = func() { // SWAP E
		z._cbprefix.SWAPr_r(&z._r.E)
	}

	z._cbprefix[0x34] = func() { // SWAP H
		z._cbprefix.SWAPr_r(&z._r.H)
	}

	z._cbprefix[0x35] = func() { // SWAP L
		z._cbprefix.SWAPr_r(&z._r.L)
	}

	z._cbprefix[0x36] = XX

	z._cbprefix[0x37] = func() { // SWAP A
		z._cbprefix.SWAPr_r(&z._r.B)
	}

	z._cbprefix[0x38] = func() { // SRL B
		z._cbprefix.SRLr_r(&z._r.B)
	}

	z._cbprefix[0x39] = func() { // SRL C
		z._cbprefix.SRLr_r(&z._r.C)
	}

	z._cbprefix[0x3A] = func() { // SRL D
		z._cbprefix.SRLr_r(&z._r.D)
	}

	z._cbprefix[0x3B] = func() { // SRL E
		z._cbprefix.SRLr_r(&z._r.E)
	}

	z._cbprefix[0x3C] = func() { // SRL H
		z._cbprefix.SRLr_r(&z._r.H)
	}

	z._cbprefix[0x3D] = func() { // SRL L
		z._cbprefix.SRLr_r(&z._r.L)
	}

	z._cbprefix[0x3E] = XX

	z._cbprefix[0x3F] = func() { // SRL A
		z._cbprefix.SRLr_r(&z._r.B)
	}

	// CB40
	z._cbprefix[0x40] = func() { //BIT 0,B
		z._cbprefix.BITxr(0, &z._r.B)
	}

	z._cbprefix[0x41] = func() { //BIT 0,C
		z._cbprefix.BITxr(0, &z._r.C)
	}

	z._cbprefix[0x42] = func() { //BIT 0,D
		z._cbprefix.BITxr(0, &z._r.D)
	}

	z._cbprefix[0x43] = func() { //BIT 0,E
		z._cbprefix.BITxr(0, &z._r.E)
	}

	z._cbprefix[0x44] = func() { //BIT 0,H
		z._cbprefix.BITxr(0, &z._r.H)
	}

	z._cbprefix[0x45] = func() { //BIT 0,L
		z._cbprefix.BITxr(0, &z._r.L)
	}

	z._cbprefix[0x46] = func() { //BIT 0,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(0, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x47] = func() { //BIT 0,A
		z._cbprefix.BITxr(0, &z._r.A)
	}

	z._cbprefix[0x48] = func() { //BIT 1,B
		z._cbprefix.BITxr(1, &z._r.B)
	}

	z._cbprefix[0x49] = func() { //BIT 1,C
		z._cbprefix.BITxr(1, &z._r.C)
	}

	z._cbprefix[0x4A] = func() { //BIT 1,D
		z._cbprefix.BITxr(1, &z._r.D)
	}

	z._cbprefix[0x4B] = func() { //BIT 1,E
		z._cbprefix.BITxr(1, &z._r.E)
	}

	z._cbprefix[0x4C] = func() { //BIT 1,H
		z._cbprefix.BITxr(1, &z._r.H)
	}

	z._cbprefix[0x4D] = func() { //BIT 1,L
		z._cbprefix.BITxr(1, &z._r.L)
	}

	z._cbprefix[0x4E] = func() { //BIT 1,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(1, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x4F] = func() { //BIT 1,A
		z._cbprefix.BITxr(1, &z._r.A)
	}

	// CB50
	z._cbprefix[0x50] = func() { //BIT 2,B
		z._cbprefix.BITxr(2, &z._r.B)
	}

	z._cbprefix[0x51] = func() { //BIT 2,C
		z._cbprefix.BITxr(2, &z._r.C)
	}

	z._cbprefix[0x52] = func() { //BIT 2,D
		z._cbprefix.BITxr(2, &z._r.D)
	}

	z._cbprefix[0x53] = func() { //BIT 2,E
		z._cbprefix.BITxr(2, &z._r.E)
	}

	z._cbprefix[0x54] = func() { //BIT 2,H
		z._cbprefix.BITxr(2, &z._r.H)
	}

	z._cbprefix[0x55] = func() { //BIT 2,L
		z._cbprefix.BITxr(2, &z._r.L)
	}

	z._cbprefix[0x56] = func() { //BIT 2,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(2, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x57] = func() { //BIT 2,A
		z._cbprefix.BITxr(2, &z._r.A)
	}

	z._cbprefix[0x58] = func() { //BIT 3,B
		z._cbprefix.BITxr(3, &z._r.B)
	}

	z._cbprefix[0x59] = func() { //BIT 3,C
		z._cbprefix.BITxr(3, &z._r.C)
	}

	z._cbprefix[0x5A] = func() { //BIT 3,D
		z._cbprefix.BITxr(3, &z._r.D)
	}

	z._cbprefix[0x5B] = func() { //BIT 3,E
		z._cbprefix.BITxr(3, &z._r.E)
	}

	z._cbprefix[0x5C] = func() { //BIT 3,H
		z._cbprefix.BITxr(3, &z._r.H)
	}

	z._cbprefix[0x5D] = func() { //BIT 3,L
		z._cbprefix.BITxr(3, &z._r.L)
	}

	z._cbprefix[0x5E] = func() { //BIT 3,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(3, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x5F] = func() { //BIT 3,A
		z._cbprefix.BITxr(3, &z._r.A)
	}

	// CB60
	z._cbprefix[0x60] = func() { //BIT 4,B
		z._cbprefix.BITxr(4, &z._r.B)
	}

	z._cbprefix[0x61] = func() { //BIT 4,C
		z._cbprefix.BITxr(4, &z._r.C)
	}

	z._cbprefix[0x62] = func() { //BIT 4,D
		z._cbprefix.BITxr(4, &z._r.D)
	}

	z._cbprefix[0x63] = func() { //BIT 4,E
		z._cbprefix.BITxr(4, &z._r.E)
	}

	z._cbprefix[0x64] = func() { //BIT 4,H
		z._cbprefix.BITxr(4, &z._r.H)
	}

	z._cbprefix[0x65] = func() { //BIT 4,L
		z._cbprefix.BITxr(4, &z._r.L)
	}

	z._cbprefix[0x66] = func() { //BIT 4,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(4, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x67] = func() { //BIT 4,A
		z._cbprefix.BITxr(4, &z._r.A)
	}

	z._cbprefix[0x68] = func() { //BIT 5,B
		z._cbprefix.BITxr(5, &z._r.B)
	}

	z._cbprefix[0x69] = func() { //BIT 5,C
		z._cbprefix.BITxr(5, &z._r.C)
	}

	z._cbprefix[0x6A] = func() { //BIT 5,D
		z._cbprefix.BITxr(5, &z._r.D)
	}

	z._cbprefix[0x6B] = func() { //BIT 5,E
		z._cbprefix.BITxr(5, &z._r.E)
	}

	z._cbprefix[0x6C] = func() { //BIT 5,H
		z._cbprefix.BITxr(5, &z._r.H)
	}

	z._cbprefix[0x6D] = func() { //BIT 5,L
		z._cbprefix.BITxr(5, &z._r.L)
	}

	z._cbprefix[0x6E] = func() { //BIT 5,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(5, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x6F] = func() { //BIT 5,A
		z._cbprefix.BITxr(5, &z._r.A)
	}

	// CB70
	z._cbprefix[0x70] = func() { //BIT 6,B
		z._cbprefix.BITxr(6, &z._r.B)
	}

	z._cbprefix[0x71] = func() { //BIT 6,C
		z._cbprefix.BITxr(6, &z._r.C)
	}

	z._cbprefix[0x72] = func() { //BIT 6,D
		z._cbprefix.BITxr(6, &z._r.D)
	}

	z._cbprefix[0x73] = func() { //BIT 6,E
		z._cbprefix.BITxr(6, &z._r.E)
	}

	z._cbprefix[0x74] = func() { //BIT 6,H
		z._cbprefix.BITxr(6, &z._r.H)
	}

	z._cbprefix[0x75] = func() { //BIT 6,L
		z._cbprefix.BITxr(6, &z._r.L)
	}

	z._cbprefix[0x76] = func() { //BIT 6,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(6, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x77] = func() { //BIT 6,A
		z._cbprefix.BITxr(6, &z._r.A)
	}

	z._cbprefix[0x78] = func() { //BIT 7,B
		z._cbprefix.BITxr(7, &z._r.B)
	}

	z._cbprefix[0x79] = func() { //BIT 7,C
		z._cbprefix.BITxr(7, &z._r.C)
	}

	z._cbprefix[0x7A] = func() { //BIT 7,D
		z._cbprefix.BITxr(7, &z._r.D)
	}

	z._cbprefix[0x7B] = func() { //BIT 7,E
		z._cbprefix.BITxr(7, &z._r.E)
	}

	z._cbprefix[0x7C] = func() { //BIT 7,H
		z._cbprefix.BITxr(7, &z._r.H)
	}

	z._cbprefix[0x7D] = func() { //BIT 7,L
		z._cbprefix.BITxr(7, &z._r.L)
	}

	z._cbprefix[0x7E] = func() { //BIT 7,(HL)
		addr := CombineToAddress(z._r.H, z._r.L)
		hl, err := mmu.ReadByte(addr)
		z._cbprefix.BITxr(7, &hl)

		z._r.M = 3
		z._r.T = 12

		LogErrors(err)
	}

	z._cbprefix[0x7F] = func() { //BIT 7,A
		z._cbprefix.BITxr(7, &z._r.A)
	}

	// CB80 > CBFF
	for i := 0x80; i <= 0xFF; i++ {
		z._cbprefix[i] = XX
	}
}

// Helper methods
// Rotate register r and shift left 1 bit (Carry)
func (ins *Instructions) RLCr_r(r *byte) {
	var ci, co byte

	if (*r & 0x80) != 0 {
		ci = 1
		co = 0x10
	} else {
		ci = 0
		co = 0
	}

	*r = (*r << 1) + ci
	ins.ZeroF(*r, 0)

	z80._r.F = (z80._r.F & 0xEF) + co

	z80._r.M = 2
	z80._r.T = 8
}

// Rotate register r and shift 1 bit right (Carry)
func (ins *Instructions) RRCr_r(r *byte) {
	var ci, co byte

	if (*r & 0x1) != 0 {
		ci = 0x80
		co = 0x10
	} else {
		ci = 0
		co = 0
	}

	*r = (*r >> 1) + ci
	ins.ZeroF(*r, 0)

	z80._r.F = (z80._r.F & 0xEF) + co

	z80._r.M = 2
	z80._r.T = 8
}

// Rotate register r and shift to the left 1 bit
func (ins *Instructions) RLr_r(r *byte) {
	var ci, co byte

	if (*r & 0x10) != 0 {
		ci = 1
		co = 0x10
	} else {
		ci = 0
		co = 0
	}

	*r = (*r << 1) + ci
	ins.ZeroF(*r, 0)

	z80._r.F = (z80._r.F & 0xEF) + co

	z80._r.M = 2
	z80._r.T = 8
}

// Rotate register r and shift to the right 1 bit
func (ins *Instructions) RRr_r(r *byte) {
	var ci, co byte

	if (*r & 0x1) != 0 {
		ci = 0x80
		co = 0x10
	} else {
		ci = 0
		co = 0
	}

	*r = (*r >> 1) + ci
	ins.ZeroF(*r, 0)

	z80._r.F = (z80._r.F & 0xEF) + co

	z80._r.M = 2
	z80._r.T = 8
}

// Arithmatic shift left
func (ins *Instructions) SLAr_r(r *byte) {
	var co byte

	if (*r & 0x80) != 0 {
		co = 0x10
	} else {
		co = 0
	}

	*r = *r << 1
	ins.ZeroF(*r, 0)

	z80._r.F = (z80._r.F & 0xEF) + co

	z80._r.M = 2
	z80._r.T = 8
}

// Arithmatic shift right
func (ins *Instructions) SRAr_r(r *byte) {
	var co byte

	ci := *r & 0x80
	if (*r & 0x1) != 0 {
		co = 0x10
	} else {
		co = 0
	}

	*r = (*r >> 1) + ci
	ins.ZeroF(*r, 0)

	z80._r.F = (z80._r.F & 0xEF) + co

	z80._r.M = 2
	z80._r.T = 8
}

// Swap register r with (hl)
func (ins *Instructions) SWAPr_r(r *byte) {
	var err, err2 error

	tr := *r
	*r, err = mmu.ReadByte(CombineToAddress(z80._r.H, z80._r.L))
	err2 = mmu.WriteByte(CombineToAddress(z80._r.H, z80._r.L), tr)

	z80._r.M = 4
	z80._r.L = 16

	LogErrors(err, err2)
}

// Shift register r right 1 bit
func (ins *Instructions) SRLr_r(r *byte) {
	var co byte

	if (*r & 1) != 0 {
		co = 0x10
	} else {
		co = 0
	}

	*r = *r >> 1
	ins.ZeroF(*r, 0)

	z80._r.F = (z80._r.F & 0xEF) + co

	z80._r.M = 2
	z80._r.T = 8
}

// Bit Manipulation
func (ins *Instructions) BITxr(x byte, r *byte) {
	operand := x << 1

	if operand == 0 {
		operand = 1
	}

	ins.ZeroF((*r & operand), 0)

	z80._r.M = 2
	z80._r.T = 8
}
