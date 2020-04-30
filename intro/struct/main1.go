package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Bicara() string {
	return "Hello"
}

func (p Person) Langkah() string {
	return "Berjalan"
}

func (p Person) SayHello(helo string) string {
	return p.Name + helo
}

func (p *Person) ChangeName(name string) (bool, error) {
	if name == "" {
		return false, errors.New("Error tidak bisa ganti karena kosong")
	}

	p.Name = name
	return true, nil
}

func main() {
	p := Person{Name: "Angga", Age: 12}

	ok, err := p.ChangeName("")
	if ok {
		fmt.Println("Pergantian nama berhasil")
	} else {
		fmt.Println(err)
	}

	// fmt.Println(p)
	// fmt.Println(p.Bicara())
	// fmt.Println(p.Langkah())
	// fmt.Println(p.SayHello("Selamat Pagi"))
}
