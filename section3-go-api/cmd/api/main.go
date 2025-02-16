package main

import "github.com/lapeko/udemy__github-actions-the-complete-guide/section3-go-api/internal/api"

func main() {
	a := api.New()
	a.Init()
	a.Start()
}
