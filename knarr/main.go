package main

import "github.com/Zadigo/knarr/internal/logic"

func main() {
	tableLayer := logic.NewTableLayer("table-id")
	player := tableLayer.Layer.GetPlayer("")
}
