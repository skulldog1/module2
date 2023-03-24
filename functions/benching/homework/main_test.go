package main

import (
	"testing"
)

func BenchmarkMapUserProducts(b *testing.B) {
	users := genUsers()
	products := genProducts()
	for n := 0; n < b.N; n++ {
		MapUserProducts(users, products)
	}
}

func BenchmarkMapUserProducts2(b *testing.B) {
	users := genUsers()
	products := genProducts()
	for n := 0; n < b.N; n++ {
		MapUserProducts2(users, products)
	}
}
