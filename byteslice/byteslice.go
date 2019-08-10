package main

import "fmt"

type ByteSlice []byte

func (p *ByteSlice) Write (data []byte) (n int, err error) {
	slice := *p

	l := len(slice)
	if l + len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data)*2))
		// The copy function is predeclared and wroks for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	copy(slice[l:], data)

	*p = slice
	return len(data), nil
}

func main () {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Println(b)
}