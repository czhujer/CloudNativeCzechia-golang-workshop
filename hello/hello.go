package main

import "fmt"
import "hello/print"

func main() {
	fmt.Println("Hello world!")
}

func init() {
	print.Print("INIT over custom print package")
}
