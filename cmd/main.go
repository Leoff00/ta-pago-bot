package main

import (
	"fmt"

	"github.com/leoff00/ta-pago-bot/internal/bot"
	"github.com/leoff00/ta-pago-bot/internal/db"
)

const displayArt = "\033[36m" + ` 
$$$$$$$$\  $$$$$$\        $$$$$$$\   $$$$$$\   $$$$$$\   $$$$$$\  $$\ 
\__$$  __|$$  __$$\       $$  __$$\ $$  __$$\ $$  __$$\ $$  __$$\ $$ |
   $$ |   $$ /  $$ |      $$ |  $$ |$$ /  $$ |$$ /  \__|$$ /  $$ |$$ |
   $$ |   $$$$$$$$ |      $$$$$$$  |$$$$$$$$ |$$ |$$$$\ $$ |  $$ |$$ |
   $$ |   $$  __$$ |      $$  ____/ $$  __$$ |$$ |\_$$ |$$ |  $$ |\__|
   $$ |   $$ |  $$ |      $$ |      $$ |  $$ |$$ |  $$ |$$ |  $$ |    
   $$ |   $$ |  $$ |      $$ |      $$ |  $$ |\$$$$$$  | $$$$$$  |$$\ 
   \__|   \__|  \__|      \__|      \__|  \__| \______/  \______/ \__|
                                 The bot.
                                  ❚█══█❚
` + "\033[0m"

func main() {
	fmt.Println(displayArt)
	db.Setup()

	bot.Start()
}
