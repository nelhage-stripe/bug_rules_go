package cgo_tags

// #include <windows.h>
import "C"

import "log"

func myfunc() {
	log.Printf("This is Windows")
}
