package main

import "fmt"

type ClientImple struct {
}

func (my ClientImple) HttpClient(msg string) (string, error) {
	fmt.Printf("ClientImple ClientImple is : %s , %s", msg, my)
	return fmt.Sprintf("}%v{", msg), nil
}

func (my ClientImple) ResponseGen(msg string) (string, error) {
	fmt.Printf("ClientImple ResponseGen is : %s , %s", msg, my)
	return fmt.Sprintf(">%v<", msg), nil
}
