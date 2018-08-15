package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config is an object to match the JSON
type Config struct {
	ServerName string
	Code       string
	Database   struct {
		Name     string
		Host     string
		Username string
		Password string
	}
}

func main() {
	// in Go, it asks us to create a black object of a type first,
	// then we can apply the JSON string to that object
	conf := Config{}
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// create a matching data struct, then unmarshal data into json.Unmarshal func
	// takes byte array JSON data, and pointer to the onkect that should have to data applied to
	err = json.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server Name: %s (%s) \n", conf.ServerName, conf.Code)

	db := conf.Database
	fmt.Printf(
		"Information- %s/%s/%s/%s\n",
		db.Name,
		db.Host,
		db.Username,
		db.Password,
	)
}
