package main

import (
	"fmt"
	"log"
	"os"

	"github.com/loveuse/learn-go-with-tests/time"
)

const dbFileName = "game.db.json"

func main() {

	store, err := poker_time.FileSystemPlayersStoreFromFile(dbFileName)
	alerter := poker_time.BlindAlerterFunc(poker_time.StdOutAlerter)

	if err != nil {
		log.Fatal(err)
	}

	game := poker_time.NewTexasHoldem(store, alerter)
	cli := poker_time.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()

}
