//golangcitest:config_path testdata/gofmt.yml
package testdata

/*
 #include <stdio.h>
 #include <stdlib.h>

 void myprint(char* s) {
 	printf("%d\n", s);
 }
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func _() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}

func GofmtNotSimplified() {
	var x []string
	fmt.Print(x[1:len(x)]) // want "File is not properly formatted"
}
