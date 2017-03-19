package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/tecnoscimmie/tsredditbot/support"
)

var baseURL = ""
var conf support.ConfigFile
var err error
var redditUsername string
var redditPassword string
var redditClientID string
var redditClientSecret string

func main() {
	// parse ALL the parameters!
	parametersParser()

	// check for presence and correctness of the configuration file
	conf, err = support.CheckConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	// set the base URL for api calls
	baseURL = "https://api.telegram.org/bot" + conf.BotToken + "/"

	// set the webhook as the configuration file says
	_, err = http.PostForm(baseURL+"/setWebhook", url.Values{"url": {conf.URL + ":" + conf.Port + "/" + conf.Endpoint}})
	if err != nil {
		log.Fatal(err)
	}

	// start the bot!
	fmt.Println("--> Starting tsreddit bot")
	err = support.PrintBotInformations(conf.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/"+conf.Endpoint, endpointHandler)
	log.Fatal(http.ListenAndServeTLS(":"+conf.Port, conf.CertPath, conf.KeyPath, nil))
}

// handle the webhook data
func endpointHandler(w http.ResponseWriter, r *http.Request) {
	data := support.LoadJSONToTelegramObject(r.Body)
	echoText := data.Message.Text
	secureSendMessage(data, echoText)

}

// simple function to send a message back to its chat, and check for security
func secureSendMessage(tObj support.TelegramObject, text string) {

	recipient := tObj.Message.From.Username
	_ = recipient

	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(tObj.Message.Chat.ID))
	params.Set("text", "Not authorized.")

	/*for _, username := range conf.AuthorizedUsers {
		if username == recipient {
			params.Del("text")
			params.Set("text", text)
			break
		}
	}*/

	_, err := http.PostForm(baseURL+"sendMessage", params)
	if err != nil {
		log.Fatal(err)
	}
}

func parametersParser() {
	// redefine flag.Usage(), because a little bit of branding is always good
	flag.Usage = func() {
		fmt.Printf("tsreddit: TecnoScimmie's reddit posting bot.\n\nUsage:\n")
		flag.PrintDefaults()
	}

	// define and parse all the parameters available
	flag.StringVar(&redditUsername, "username", "", "reddit bot username")
	flag.StringVar(&redditPassword, "password", "", "reddit bot password")
	flag.StringVar(&redditClientID, "clientid", "", "reddit bot client ID")
	flag.StringVar(&redditClientSecret, "secret", "", "reddit bot secret")
	flag.Parse()

	// do we have all the parameters needed to run?
	if redditClientID == "" || redditClientSecret == "" || redditUsername == "" || redditPassword == "" {
		flag.Usage()
		os.Exit(1)
	}
}
