package poker_time

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
