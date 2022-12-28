package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const database = "database"

type Database struct {
	File string
}

func NewDatabase() *Database {
	return &Database{File: database}
}

func (d *Database) Close() error {
	return nil
}

func (d *Database) Get(key string) (string, error) {
	file, err := os.OpenFile(d.File, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid line: %s", line)
		}
		if parts[0] == key {
			return parts[1], nil
		}
	}
	return "", nil
}

func (d *Database) Set(key, value string) error {
	file, err := os.OpenFile(d.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = file.WriteString(key + "," + value + "\n")
	return err
}

func main() {
	db := NewDatabase()
	defer db.Close()

	err := db.Set("foo", "bar")
	if err != nil {
		panic(err)
	}
	db.Set("foo1", "bar")

	value, err := db.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
