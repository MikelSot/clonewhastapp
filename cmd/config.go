package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	AllowedOrigins  []string `json:"allowed_origins"`
	AllowedMethods  []string `json:"allowed_methods"`
	LogFolder       string   `json:"log_folder"`
	Env             string   `json:"env"`
	PortHTTP        uint     `json:"port_http"`
	PrivateFileSign string   `json:"private_file_sign"`
	PublicFileSign  string   `json:"public_file_sign"`
	Database        Database `json:"database"`
}

func newConfiguration(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	conf := Configuration{}
	if err := json.Unmarshal(file, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}

type Database struct {
	Engine   string `json:"engine"`
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     uint   `json:"port"`
	Name     string `json:"name"`
	SSLMode  string `json:"ssl_mode"`
}

func newDatabase(engine, user, password, server, dbname, sslMode string, port uint) Database {
	return Database{
		Engine:   engine,
		User:     user,
		Password: password,
		Server:   server,
		Port:     port,
		Name:     dbname,
		SSLMode:  sslMode,
	}
}
