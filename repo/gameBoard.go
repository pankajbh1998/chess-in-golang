package repo

import "Chess/models"

type ChessBoard struct{
	cell map[models.Position]models.Piece
	size int
}

func NewChessBoard(boardSize int)*ChessBoard{
	return &ChessBoard{
		cell: make(map[models.Position]models.Piece, 0),
		size: boardSize,
	}
}

func (g *ChessBoard)IsPositionFilled(pos *models.Position)(models.Piece, bool){
	position, ok := g.cell[*pos]
	return position, ok
}

func (g *ChessBoard) RemovePiece(pos *models.Position){
	delete(g.cell, *pos)
}

func (g *ChessBoard) SetPiece(pos *models.Position, piece models.Piece){
	g.cell[*pos] = piece
}

func (g *ChessBoard) GetBoardSize()int{
	return g.size
}


