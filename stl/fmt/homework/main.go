package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func generateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}

func main() {
	p := Person{Name: "Andy", Age: 18}

	// вывод значений структуры
	fmt.Println("simple struct:", p)

	// вывод названий полей и их значений
	fmt.Printf("detailed struct: %+v\n", p)

	// вывод названий полей и их значений в виде инициализации
	fmt.Printf("Golang struct: %#v\n", p)

	money := 10.00

	fmt.Println(generateSelfStory(p.Name, p.Age, money))

}
