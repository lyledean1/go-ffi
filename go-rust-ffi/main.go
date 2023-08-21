package main

/*
#cgo LDFLAGS: -L. ./ffi_rs/target/release/libffi_rs.dylib -ldl -ldl -lpthread

#include <stdlib.h>

void print_value(const char *value);
*/
import "C"
import "unsafe"

func main() {
	str := C.CString("Hello from Go!")
	defer C.free(unsafe.Pointer(str))

	C.print_value(str)
}
