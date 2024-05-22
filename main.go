package main

import (
	"quakes/internal/services"
)

func main() {
	data := services.ByYear(2024)
	println(data)

}
