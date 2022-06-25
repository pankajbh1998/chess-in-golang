package models

type Piece interface {
	GetColor()Color
	GetName ()PieceName
	GetValidMoves(b Board, curr *Position)map[Position]bool
	IsSetNewPiece(newPos *Position)bool
}


