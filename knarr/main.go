package main

import "github.com/Zadigo/knarr/internal/logic"

func main() {
	tableLayer := logic.NewTableLayer("table-id")
	tableLayer.Layer.GetTableId()
}
