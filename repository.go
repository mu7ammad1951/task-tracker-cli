package main

import (
	"io"
	"log"
	"os"
)

const fileName = "./tasks.json"

func LoadData() []byte {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("failed to open data file")
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("failed to read data")
	}

	return data
}

func SaveData(data []byte) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Fatal("failed to save file")
	}
}
