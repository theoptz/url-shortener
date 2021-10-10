package main

import app2 "github.com/theoptz/url-shortener/internal/infrastructure/app"

func main() {
	app := app2.NewApp()
	app.Run()
}
