package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdulrahim-m/team-manager/internal/config"
	"github.com/abdulrahim-m/team-manager/internal/discord"
	"github.com/abdulrahim-m/team-manager/internal/telegram"
	"github.com/lpernett/godotenv"
)

func main() {
	godotenv.Load()
	config.LoadConfig()

	teleToDiscChannel := make(chan string, 10)
	// discToTele := make(chan string, 10)

	go discord.StartDiscordBot(teleToDiscChannel)
	go telegram.StartTelegramBot(teleToDiscChannel)

	printLogo()
	log.Println("Server started on :53053...")
	err := http.ListenAndServe(":53053", nil)
	log.Fatal(err)

	select {}
}

func printLogo() {
	fmt.Println(`
   /\\\\\\\\\\\\\\\\\     /\\\\                 /\\\\
   \/\\\\////////////\\\\ \/\\\\\\\           /\\\\\\\
    \/\\\\          \/\\\\ \/\\\\//\\\      /\\\///\\\\
     \/\\\\\\\\\\\\\\\\///  \/\\\\\///\\\ /\\\//  \/\\\\
      \/\\\\///////////      \/\\\\  \///\\\//     \/\\\\
       \/\\\\                 \/\\\\    \///        \/\\\\
        \/\\\\                 \/\\\\                \/\\\\
         \/\\\\                 \/\\\\                \/\\\\
          \/\\\\                 \/\\\\                \/\\\\
           \/\\\\                 \/\\\\                \/\\\\
            \////                  \////                 \////
		`)
}
