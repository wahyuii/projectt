package main

import (
	"fmt"
	ikhsan "projectt/app"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	a := &ikhsan.App{}
	a.Initialize()
	a.Run()
}
