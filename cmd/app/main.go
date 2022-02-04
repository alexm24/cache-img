package main

import "github.com/alexm24/cache-img/internal/app"

const configPath = "configs/config.yml"

func main() {
	app.Run(configPath)
}
