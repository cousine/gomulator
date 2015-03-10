package gboy

import (
	"log"
)

var (
	z80 Z80
	mmu MMU
)

// Logs an error to the console if any
func LogErrors(errs ...error) {
	for _, err := range errs {
		if err != nil {
			log.Println(err)
		}
	}
}

// Combines 2 bytes (registers) into an address (word)
func CombineToAddress(r1, r2 byte) Address {
	return (Address(r1) << 8) + Address(r2)
}
