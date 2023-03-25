package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Cache struct {
	data map[string]*User
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]*User, 100)}
}

func (c *Cache) Set(key string, u *User) {
	c.data[key] = u
}

func (c *Cache) Get(key string) *User {
	return c.data[key]
}

type User struct {
	ID       int
	Nickname string
	Email    string
}

func main() {
	var users []*User
	emails := []string{"robpike@gmail.com", "davecheney@gmail.com", "bradfitzpatrick@email.ru", "eliben@gmail.com", "quasilyte@mail.ru"}
	for i := range emails {
		nickname := strings.Split(emails[i], "@")[0]
		users = append(users, &User{
			ID:       i + 1,
			Nickname: nickname,
			Email:    emails[i],
		})
	}
	cache := NewCache()
	for i := range users {
		key := strings.Join([]string{users[i].Nickname, strconv.Itoa(users[i].ID)}, ":")
		cache.Set(key, users[i])
	}
	keys := []string{"robpike:1", "davecheney:2", "bradfitzpatrick:3", "eliben:4", "quasilyte:5"}
	for i := range keys {
		fmt.Println(cache.Get(keys[i]))
	}
}
