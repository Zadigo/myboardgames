package backend

import (
	"encoding/json"
	"log"
	"os"
)

type JsonFileData struct {
	levelOne   []MarvelCard
	levelTwo   []MarvelCard
	levelThree []MarvelCard
}

func ReadJsonFile(path string) JsonFileData {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data JsonFileData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		panic(err)
	}

	log.Print(data)

	return data
}
