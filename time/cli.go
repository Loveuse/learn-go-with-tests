package poker_time

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	// PlayerPrompt is the text asking the user for the number of players
	PlayerPrompt = "Please enter the number of players: "

	// BadPlayerInputErrMsg is the text telling the user they did bad things
	BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

	// BadWinnerInputMsg is the text telling the user they declared the winner wrong
	BadWinnerInputMsg = "invalid winner input, expect format of 'PlayerName wins'"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// PlayPoker starts the game
func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(c.readLine())

	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}

	c.game.Start(numberOfPlayers)

	winnerInput := c.readLine()
	winner, err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(c.out, BadWinnerInputMsg)
		return
	}

	c.game.Finish(winner)
}
func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}
