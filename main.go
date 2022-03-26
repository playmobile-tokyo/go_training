package main

import "fmt"

type TestObj struct {
	Id   int
	Name string
}

func main() {
	str := display("aaa")
	println(str)
}

func display(msg string) string {
	return fmt.Sprintf("[%v]", msg)
}
