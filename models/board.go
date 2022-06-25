package models

type Board interface{
	IsPositionFilled(pos *Position)(Piece, bool)
	GetBoardSize()int
}
