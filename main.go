package main

import (
    "flag"

    "github.com/project-wind/wind-comment/app"
)

func main() {
    var configPath string
    flag.StringVar(&configPath, "config", "./config.json", "The config path")
    flag.StringVar(&configPath, "c", "./config.json", "The short alias of -config")
    flag.Parse()

    app.Init(configPath)
}
