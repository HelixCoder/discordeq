package menu

import (
	"fmt"
	"github.com/xackery/eqemuconfig"
	"os"
	"strings"
)

func ShowMenu() (err error) {
	var option string
	var isConfigLoaded bool
	fmt.Println("\n===DiscordEQ Plugin===")

	config, err := eqemuconfig.LoadConfig()
	status := "Good"
	if err != nil {
		status = fmt.Sprintf("Bad (%s)", err.Error())
	} else {
		isConfigLoaded = true
		status = fmt.Sprintf("Good (%s)", config.Longame)
	}
	fmt.Printf("1) Reload eqemu_config.xml (Status: %s)\n", status)
	if isConfigLoaded {
		isEverythingGood := true

		status = "Good"
		if config.Discord.Username == "" || config.Discord.Password == "" {
			status = "Not configured"
			isEverythingGood = false
		}
		if config.Discord.ServerID == "" && config.Discord.ChannelID == "" {
			status = "Bad"
			isEverythingGood = false
		}
		fmt.Printf("2) Configure Discord settings inside eqemu_config.xml (Status: %s)\n", status)

		status = "Good"

		if !isChatLoggingEnabled(&config) {
			isEverythingGood = false
			status = "Bad"
		}
		fmt.Printf("3) Enable Chat Logging (Status: %s)\n", status)
		fmt.Printf("4) Quest File For Discord Chat (Status: %s)\n", status)
		if isEverythingGood {
			fmt.Println("5) Start DiscordEQ")
		}
	}
	fmt.Println("Q) Quit")

	fmt.Scan(&option)
	fmt.Println("You chose option:", option)
	option = strings.ToLower(option)
	if option == "q" || option == "exit" || option == "quit" {
		fmt.Println("Quitting")
		os.Exit(0)
	} else if option == "2" {
		err = fmt.Errorf("Configuring discord")
		for err != nil {
			err = menuDiscord(&config)
		}
		err = fmt.Errorf("Configuring discord")

	} else if option == "3" {
		err = fmt.Errorf("Configuring SQL")
		for err != nil {
			err = menuChatLog(&config)
		}
		err = fmt.Errorf("Configuring SQL")

	} else {
		fmt.Println("Invalid option")
		err = fmt.Errorf("No option chosen")
	}
	return
}
