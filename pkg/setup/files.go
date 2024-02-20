package setup

import (
	"log"
	"os"
)

func Pwd() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Default().Fatalln("Error getting the current working directory ->", err.Error())
	}
	log.Default().Println("Current working directory ->", cwd)
}
