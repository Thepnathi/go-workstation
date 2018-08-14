package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ServerName string `json:"Name of Server"`
	Database   struct {
		Name     string
		Host     string
		Username string
		Password string
	}
}

func main() {
	conf := Config{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server Name: %s (%s)", conf.ServerName)

	db := conf.Database
	fmt.Printf(
		"Information- %s-%s-%s-%s",
		db.Name,
		db.Host,
		db.Username,
		db.Password,
	)
}
