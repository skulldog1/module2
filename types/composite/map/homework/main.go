package main

import "fmt"

type Project struct {
	Name  string
	Stars int
}

func main() {
	projects := []Project{
		{
			Name:  "https://github.com/docker/compose",
			Stars: 27600,
		},
		{
			Name:  "https://github.com/GoesToEleven/GolangTraining",
			Stars: 8800,
		},
		{
			Name:  "https://github.com/docker-library/golang",
			Stars: 1300,
		},
		{
			Name:  "https://github.com/golangci/golangci-lint",
			Stars: 12000,
		},
		{
			Name:  "https://github.com/redis/go-redis",
			Stars: 16600,
		},
		{
			Name:  "https://github.com/polaris1119/golangweekly",
			Stars: 1900,
		},
		{
			Name:  "https://github.com/gocolly/colly",
			Stars: 19000,
		},
		{
			Name:  "https://github.com/gothinkster/golang-gin-realworld-example-app",
			Stars: 2200,
		},
		{
			Name:  "https://github.com/hprose/hprose-golang",
			Stars: 1200,
		},
		{
			Name:  "https://github.com/graarh/golang-socketio",
			Stars: 413,
		},
		{
			Name:  "https://github.com/hyper0x/go_command_tutorial",
			Stars: 3400,
		},
		{
			Name:  "https://github.com/codebangkok/golang",
			Stars: 104,
		},
		{
			Name:  "https://github.com/mailru/easyjson",
			Stars: 4100,
		},
	}

	// в цикле запишите в map
	karta := map[string]Project{}
	for key, val := range projects {
		karta[val.Name] = projects[key]
	}
	// в цикле пройдитесь по мапе и выведите значения в консоль
	for i, k := range karta {
		fmt.Println(i, k)
	}
}
