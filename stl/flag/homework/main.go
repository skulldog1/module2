package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type Config struct {
	AppName    string `json:"app_name"`
	Production bool   `json:"production"`
}

func main() {
	confFile := flag.String("conf", "", "path to config file")
	flag.Parse()

	if *confFile == "" {
		fmt.Println("config file path is required")
		return
	}

	fileContent, err := ioutil.ReadFile(*confFile)
	if err != nil {
		fmt.Println("error reading config file:", err)
		return
	}

	conf := Config{}
	err = json.Unmarshal(fileContent, &conf)
	if err != nil {
		fmt.Println("error decoding config file:", err)
		return
	}

	fmt.Printf("Production: %v\nAppName: %s\n", conf.Production, conf.AppName)
}
