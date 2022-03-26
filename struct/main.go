package main

import "fmt"

type TestObj struct {
	Id   int
	Name string
}

func main() {
	str := display("aaa")
	println(str)
	repository := TestRepository{
		client: InfraStructure{},
	}
	str = repository.GetResponse(str)
	println(str)

	str = display("bbb")
	repository = TestRepository{
		client: ClientImple{},
	}
	str = repository.GetResponse(str)
	println(str)
}

func display(msg string) string {
	return fmt.Sprintf("[%v]", msg)
}

type TestRepository struct {
	client Client
}

type Client interface {
	HttpClient(string) (string, error)
	ResponseGen(string) (string, error)
}

func (rep TestRepository) GetResponse(msg string) string {
	str, err := rep.client.HttpClient(msg)
	if err != nil {
		panic(err)
	}
	str, err = rep.client.ResponseGen(str)
	if err != nil {
		panic(err)
	}
	return str
}

type InfraStructure struct {
}

func (my InfraStructure) HttpClient(msg string) (string, error) {
	fmt.Printf("ClientImpl is : %s , %s", msg, my)
	return fmt.Sprintf("{%v}", msg), nil
}

func (my InfraStructure) ResponseGen(msg string) (string, error) {
	fmt.Printf("ResponseGen is : %s , %s", msg, my)
	return fmt.Sprintf("<%v>", msg), nil
}
