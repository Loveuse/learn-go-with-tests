package main

import (
	"fmt"
	"log"
	"os"

	"github.com/loveuse/learn-go-with-tests/cmdprojectstruct"
)

const dbFileName = "game.db.json"

func main() {

	store, err := poker.FileSystemPlayersStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("could not set up the store: %v", err)
	}

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	CLI := poker.NewCLI(store, os.Stdin)
	CLI.PlayPoker()

}
