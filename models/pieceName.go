package models

type PieceName string

const(
	KingName = PieceName("K")
	QueenName = PieceName("Q")
	BishopName = PieceName("B")
	RookName = PieceName("R")
	PawnName = PieceName("P")
	KnightName = PieceName("N")
)
