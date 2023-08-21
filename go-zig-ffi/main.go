package main

/*
#cgo LDFLAGS: -L. -lmath
#include <stdint.h>

int32_t add(int32_t a, int32_t b);
*/
import "C"
import "fmt"

func main() {
	a, b := C.int32_t(5), C.int32_t(7)
	result := C.add(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, result)
}
