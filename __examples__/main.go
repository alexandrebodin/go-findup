package main

import (
	"fmt"

	"github.com/alexandrebodin/go-findup"
)

func main() {
	path, _ := findup.Find(".zshrc")
	fmt.Println(path)

	path, _ = findup.FindFrom("hosts", "/etc")
	fmt.Println(path)
}
