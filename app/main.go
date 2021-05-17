package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"io/ioutil"
	"os"
	"unsafe"
)

func example() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func safeCreateDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		   if err != nil {
			return err
		}
	}
	return nil
}

func fileTest() error {
	err := safeCreateDir("/tmp/share")
	if err != nil {
	 return err
 	}
	d1 := []byte("hello\ngo\n")
    return ioutil.WriteFile("/tmp/share/out", d1, 0777)
}

func main() {
	example()
	// Run container with "-v host/path:/tmp/container/path"
	err := fileTest()
    if err != nil {
		panic(err)
	}
}
