package main

// #include <stdio.h>
// void callC(){
//   printf("Calling C code from Golang!\n");
// }
import "C"

import "fmt"

func main() {
	fmt.Println("Go is about to call C!")
	C.callC()
	fmt.Println("C function already called successfully!")
}
