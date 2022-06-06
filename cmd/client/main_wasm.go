package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}
