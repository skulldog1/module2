package main

import (
	"fmt"
)

type calc struct {
	a, b   float64
	result float64
}

func NewCalc() *calc {
	return &calc{}
}

func (c *calc) SetA(a float64) *calc {
	c.a = a
	return c
}

func (c *calc) SetB(b float64) *calc {
	c.b = b

	return c
}

func (c *calc) Do(operation func(a, b float64) float64) *calc {
	c.result = operation(c.a, c.b)

	return c
}

func (c calc) Result() float64 {
	return c.result
}

func multiply(a, b float64) float64 { // реализуйте по примеру divide, sum, average
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}

func sum(a, b float64) float64 {
	return a + b
}
func average(a, b float64) float64 {
	return (a + b) / 2
}

func main() {
	calc := NewCalc()
	res := calc.SetA(10).SetB(34).Do(multiply).Result()
	fmt.Println(res)
	res2 := calc.Result()
	fmt.Println(res2)
	if res != res2 {
		panic("object statement is not persist")
	}
	res3 := calc.SetA(10).SetB(34).Do(divide).Result()
	fmt.Println(res3)
	res4 := calc.SetA(10).SetB(34).Do(sum).Result()
	fmt.Println(res4)
	res5 := calc.SetA(10).SetB(34).Do(average).Result()
	fmt.Println(res5)
}
