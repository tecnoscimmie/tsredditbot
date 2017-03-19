package support

import (
	"fmt"
	"net/http"
)

// PrintBotInformations pretty-prints information about the running bot
func PrintBotInformations(botKey string) error {
	baseURL := "https://api.telegram.org/bot" + botKey + "/getMe"
	data, err := http.Get(baseURL)
	botData, err := LoadJSONToUser(data.Body)

	fmt.Println("Bot handle: @" + botData.Username)
	fmt.Println("Bot name: " + botData.FirstName)

	return err
}
