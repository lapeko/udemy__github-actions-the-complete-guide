package main

import (
	"github.com/lapeko/udemy__github-actions-the-complete-guide/internal/api"
	"log"
)

func main() {
	a := api.New()
	a.Init()
	log.Fatal(a.Start())
}
