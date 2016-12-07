package main

/*
#cgo CFLAGS: -I../Dylib
#cgo LDFLAGS: -L. -lkeyboard
#include <keyboard.h>
*/
import "C"

import (
	"fmt"
)

func main() {
	C.InitKeyboard()

	fmt.Printf("\nEnter: ")

	for {
		r := C.GetCharacter()

		fmt.Printf("%c", r)

		if r == 113 {
			break
		}
	}

	C.CloseKeyboard()

}
