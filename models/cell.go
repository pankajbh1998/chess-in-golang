package models

type Cell struct{
	pos *Position
	piece *Piece
	color Color
}

func NewCell(pos *Position, piece *Piece, color Color)*Cell{
	return &Cell{
		pos: pos,
		piece: piece,
		color: color,
	}
}
