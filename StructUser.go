package main

import "github.com/gocql/gocql"

type User struct {
	Id gocql.UUID `json:"Id"`
	Name string `json:"Name"`
}