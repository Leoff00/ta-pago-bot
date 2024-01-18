package main

import (
	"log"

	"github.com/leoff00/ta-pago-bot/database"
	"github.com/leoff00/ta-pago-bot/pkg/bot"
)

func main() {
	log.Default().Println("TA PAGO! The bot.")
	
	db, err := database.Setup()
	if err != nil {
		log.Fatal(err.Error())
	}

	bot.Start(db)
}
