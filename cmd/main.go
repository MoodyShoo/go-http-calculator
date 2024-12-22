package main

import "github.com/MoodyShoo/go-http-calculator/internal/application"

func main() {
	app := application.New()
	app.RunServer()
}
