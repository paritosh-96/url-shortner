package main

import (
	"urlShortner/config"
	_ "urlShortner/config"
	"urlShortner/service"
)

func main() {
	service.Start(config.SetupParams.Port)
}
