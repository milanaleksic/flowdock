package main

import (
	"log"
	"os"

	"github.com/milanaleksic/flowdock"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Please send exactly one argument: your personal API key")
	}
	client := flowdock.NewClient(os.Args[1])
	if mentions, err := client.GetMyMentions(1); err != nil {
		log.Fatalf("Mentions could not have been fetched: %+v", err)
	} else {
		for _, m := range mentions {
			log.Printf("Mention: %+v\n", m)
		}
	}
	if privateMessages, err := client.GetMyUnreadPrivateMessages(); err != nil {
		log.Fatalf("Private message could not have been fetched: %+v", err)
	} else {
		for _, m := range privateMessages {
			log.Printf("Private message: %+v\n", m)
		}
	}
}
