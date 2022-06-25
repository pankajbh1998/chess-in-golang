package models

const (
	Active = GameState("Active")
	Paused = GameState("Paused")
	BlackWon = GameState("Black Won")
	WhiteWon = GameState("White Won")
	NewGame = GameState("New Game")
)

type GameState string
