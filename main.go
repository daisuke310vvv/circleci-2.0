package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func main() {
	fmt.Println("hello world", uuid.NewV4().String())
}
