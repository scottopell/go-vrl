package main

//#include <stdio.h>
//#include <stdlib.h>
//#include <string.h>
//#include <vrl.h>
//#cgo LDFLAGS: -L${SRCDIR}/target/release -Wl,-rpath,${SRCDIR}/target/release -lvrl_bridge -lm -ldl
import "C"
import "unsafe"

type Runtime = RustPointer

func NewRuntime() *Runtime {
	runtime := Runtime{}
	runtime.own(C.new_runtime())
	return &runtime
}

func (r *Runtime) resolve(program *Program, input string) (string, error) {
	cs := C.CString(input)
	defer C.free(unsafe.Pointer(cs))
	result := C.runtime_resolve(r.p, program.p, cs)

	if result.error != nil {
		defer C.free(unsafe.Pointer(result.error))
		return "", VrlError{str: C.GoString(result.error)}
	}

	return C.GoString((*C.char)(result.value)), nil
}

func (r *Runtime) clear() {
	C.runtime_clear(r.p)
}

func (r *Runtime) isEmpty() bool {
	return C.runtime_is_empty(r.p) != 0
}
