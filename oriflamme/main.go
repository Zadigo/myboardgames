package main

import (
	"fmt"

	"github.com/Zadigo/oriflamme/internal/models"
)

func main() {
	a := models.Player{
		Username: "example",
	}
	fmt.Printf("%v", a)
}
