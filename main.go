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

func doRequest(url string, command string) {
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

	var cats []string
	var chuck_message ChuckMessage
	if command == "random" || command == "category" {
		json.Unmarshal(body, &chuck_message)
		printColorMsg(chuck_message.Value)
	} else {
		json.Unmarshal(body, &cats)
		msg := fmt.Sprintf("Chuck's cats: \n %s\n", strings.Join(cats, ",\n "))
		printColorMsg(msg)
	}
}

func getURL(command string) string {
	var url string = ""

	switch command {
	case "random":
		url = "https://api.chucknorris.io/jokes/random"
	case "list-categories":
		url = "https://api.chucknorris.io/jokes/categories"
	case "category":
		if len(os.Args) < 3 {
			printColorMsg("Chuck Norris says you're missing a category.\nRun chuck-norris-says category [search_category]")
		} else {
			url = fmt.Sprintf("https://api.chucknorris.io/jokes/random?category=%s", os.Args[2])
		}
	default:
		msg := fmt.Sprintf("Chuck Norris doesn't know this command: %s", command)
		printColorMsg(msg)
	}

	return url
}

func printColorMsg(message string) {
	r, g, b := 255, 215, 0 // gold color
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, message)
}

func main() {
	// Check if command-line arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: chuck-norris-says [random|category|list-categories|help] [search_category]")
		return
	}

	command := os.Args[1]
	if command == "help" {
		printColorMsg("Chuck Norris needs no help, pussy.")
	}

	url := getURL(command)
	if url != "" {
		doRequest(url, command)
	}
}
