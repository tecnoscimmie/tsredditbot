package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
var botSession support.Session

func main() {
	// parse ALL the parameters!
	parametersParser()

	botSession, err = support.NewSession()
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
	defer r.Body.Close()

	err := data.DecodeJSON(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if debug {
		log.Printf("got message -> %+v\n", data)
	}

	if data.Message.Chat.ID == botSession.ChatID {
		log.Println("got URL from command /post on group ", botSession.GroupHandle, "posting on reddit...")

		url, err := splitPostCommand(data.Message.Text)

		if debug {
			log.Println("splitted url:", url)

		}

		if err != nil {
			log.Println(err)
			return
		}

		validURL, err := support.ValidateURL(url)
		if err != nil {
			if err = data.ReplyBackToChat("Invalid URL :("); err != nil {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}

		if debug {
			log.Println("validated url:", validURL)
		}

		redditSession, err := reddit.NewSession(redditUsername, redditPassword, redditClientID, redditClientSecret)
		if err != nil {
			log.Println(err)
			return
		}

		if err = redditSession.Post(validURL); err != nil {
			log.Println(err)
			return
		}

		if err = data.ReplyBackToChat("Posted!"); err != nil {
			log.Println(err)
			return
		}

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

func splitPostCommand(c string) (string, error) {
	// split the string for 6 characters (the "/post " command, space included)
	if strings.HasPrefix(c, "/post") {
		return strings.TrimSpace(c[6:]), nil
	}

	return "", errors.New("string is not a \"/post\" command")
}
