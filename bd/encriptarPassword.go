package bd

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(pass string) (string, error) {
	fmt.Println(len(pass))
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	fmt.Println(string(bytes))
	return string(bytes), err
}
