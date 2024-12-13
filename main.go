package main

import (
	"log"
	"os"

	"github.com/8ea7b571/MoliBlog/config"
	"github.com/8ea7b571/MoliBlog/internal/mApp"
)

func init() {
	_, err := os.Stat(mApp.SRC)
	if os.IsNotExist(err) {
		err = os.MkdirAll(mApp.SRC, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = os.Stat(mApp.DST)
	if os.IsNotExist(err) {
		err = os.MkdirAll(mApp.DST, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	cfg := config.MConfig{
		Host: "0.0.0.0",
		Port: 8080,
		Root: "D:/Projects/Go/MoliBlog",
	}

	mapp := mApp.NewMApp(&cfg)
	mapp.Run()
}
