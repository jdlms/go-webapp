package main

import (
	"errors"
	"fmt"
)

func main() {



}
func Connect() error {
return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("error creating users: %w", err)
	}
	return nil
}