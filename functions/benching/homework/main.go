package main

import (
	"fmt"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	ID       int64
	Name     string `fake:"{firstname}"`
	Products []Product
}

type Product struct {
	UserID int64
	Name   string `fake:"{sentence:3}"`
}

func main() {
	users := genUsers()
	fmt.Println(users)
	products := genProducts()
	fmt.Println(products)
	users = MapUserProducts2(users, products)
	fmt.Println(users)
}

func MapUserProducts(users []User, products []Product) []User {
	// Проинициализируйте карту продуктов по айди пользователей
	for i, user := range users {
		for _, product := range products { // избавьтесь от цикла в цикле
			if product.UserID == user.ID {
				users[i].Products = append(users[i].Products, product)
			}
		}
	}
	return users
}

func MapUserProducts2(users []User, products []Product) []User {
	productMap := make(map[int64][]Product)
	for _, product := range products {
		productMap[product.UserID] = append(productMap[product.UserID], product)
	}
	for i, user := range users {
		if userProducts, ok := productMap[user.ID]; ok {
			users[i].Products = userProducts
		}
	}
	return users
}

func genProducts() []Product {
	products := make([]Product, 1000)
	for i, product := range products {
		_ = gofakeit.Struct(&product)
		product.UserID = int64(rand.Intn(100) + 1)
		products[i] = product
	}

	return products
}

func genUsers() []User {
	users := make([]User, 100)
	for i, user := range users {
		_ = gofakeit.Struct(&user)
		user.ID = int64(i)
		users[i] = user
	}

	return users
}
