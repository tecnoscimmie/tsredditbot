package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/tecnoscimmie/tsredditbot/reddit"
	"github.com/tecnoscimmie/tsredditbot/support"
)

var baseURL = ""
var conf support.ConfigFile
var err error
var redditUsername string
var redditPassword string
var redditClientID string
var redditClientSecret string
var debug bool

func main() {
	// parse ALL the parameters!
	parametersParser()

	botSession, err := support.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	// start the bot!
	fmt.Println("--> Starting tsreddit bot")
	err = botSession.PrintBotInformations()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/"+botSession.Endpoint, endpointHandler)

	log.Fatal("main http server error -> ", http.ListenAndServeTLS(":"+botSession.Port, botSession.Configuration.CertPath, botSession.Configuration.KeyPath, nil))
}

// handle the webhook data
func endpointHandler(w http.ResponseWriter, r *http.Request) {
	var data support.TelegramObject
	err := data.DecodeJSON(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if debug {
		log.Printf("got message -> %+v\n", data)
	}
	echoText := data.Message.Text
	secureSendMessage(data, echoText)

}

// simple function to send a message back to its chat, and check for security
func secureSendMessage(tObj support.TelegramObject, text string) {

	recipient := tObj.Message.From.Username
	log.Println("got message from", recipient)

	params := url.Values{}
	if _, err := url.ParseRequestURI(tObj.Message.Text); err != nil || strings.HasPrefix(tObj.Message.Text, "/") {
		params.Set("chat_id", strconv.Itoa(tObj.Message.Chat.ID))
		params.Set("text", "Not a valid URL :(")
	} else {
		params.Set("chat_id", strconv.Itoa(tObj.Message.Chat.ID))
		params.Set("text", "Posted! :D")

		s, err := reddit.NewSession(redditUsername, redditPassword, redditClientID, redditClientSecret)
		if err != nil {
			return
		}
		s.Post(tObj.Message.Text)
	}
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
	flag.BoolVar(&debug, "debug", false, "enable debug output")
	flag.Parse()

	// do we have all the parameters needed to run?
	if redditClientID == "" || redditClientSecret == "" || redditUsername == "" || redditPassword == "" {
		flag.Usage()
		os.Exit(1)
	}
}
