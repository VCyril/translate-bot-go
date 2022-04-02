package main

import (
	reversoApi "github.com/BRUHItsABunny/go-reverso-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	godotenv "github.com/joho/godotenv"
	"log"
	"os"
	"translateBot/pkg/telegram"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tgBot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Error while creating BotApi: %s", err.Error())
	}
	//tgBot.Debug = true

	client := reversoApi.GetReversoClient()

	bot := telegram.NewBot(tgBot, client)
	err = bot.Start()
	if err != nil {
		log.Fatalf("Error while starting bot: %s", err.Error())
	}


	//fmt.Println(trResp.ContextResults.Results)

}