package main

import (
	"fmt"

	"github.com/Zadigo/knarr/internal/logic"
)

func main() {
	tableLayer := logic.NewTableLayer("table-id")
	player := tableLayer.Layer.GetPlayer("")
	fmt.Print(player.Uuid)
}
