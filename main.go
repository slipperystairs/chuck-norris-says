package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ChuckMessage struct {
	Value string `json:"value"`
}

func getColorMsg(message string) {
	r, g, b := 255, 215, 0 // gold color
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, message)
}

func main() {
	// Check if command-line arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: chuck-norris-says [random|category|list-categories] [search_category]")
		return
	}

	command := os.Args[1]
	if command == "help" {
		getColorMsg("Chuck Norris needs no help, pussy")
		return
	} else if command != "random" && command != "category"  && command != "list-categories" {
		msg := fmt.Sprintf("Chuck Norris doesn't know this command: %s", command)
		getColorMsg(msg)
		return
	} else if command == "list-categories" {
		categories := []string{
			"animal",
			"career",
			"celebrity",
			"dev",
			"explicit",
			"fashion",
			"food",
			"history",
			"money",
			"movie",
			"music",
			"political",
			"religion",
			"science",
			"sport",
			"travel",
		}
		msg := fmt.Sprintf("Chuck's cats: \n %s\n", strings.Join(categories, ",\n "))
		getColorMsg(msg)
		return
	}

	var category string = ""
	if command == "category" && len(os.Args) < 3 {
		getColorMsg("Chuck Norris says you're missing a category.\nRun chuck-norris-says category [search_category]")
	} else if len(os.Args) > 2 {
		category = os.Args[2]
	}

	var url string = "https://api.chucknorris.io/jokes/random"
	if command == "category" {
		url = fmt.Sprintf(url + "?category=%s", category)
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error: ", err)
		return
	}

	var chuck_message ChuckMessage
	json.Unmarshal(body, &chuck_message)
	getColorMsg(chuck_message.Value)
}
