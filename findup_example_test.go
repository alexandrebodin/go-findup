package findup_test

import (
	"fmt"
	"log"

	"github.com/alexandrebodin/go-findup"
)

func ExampleFind() {
	path, err := findup.Find(".zshrc")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path)
}
func ExampleFindFrom() {
	path, err := findup.FindFrom("nginx.log", "/var/log/nginx")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path)
}
