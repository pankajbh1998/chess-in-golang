package piece

import (
	"Chess/models"
)

// Pyada
type Pawn struct{
	color models.Color
	is1stMove bool
	destinationRow int
	dRow int
}

func NewPawn(color models.Color, destinationRow int, dRow int)*Pawn{
	return &Pawn{
		color: color,
		is1stMove: true,
		destinationRow: destinationRow,
		dRow: dRow,
	}
}

func (p *Pawn)disableFirstMove(){
	p.is1stMove = false
}

func (p *Pawn)isFirstMove() bool {
	return p.is1stMove
}

func (p *Pawn)GetValidMoves(board models.Board, curPos *models.Position)map[models.Position]bool{
	posMap := make(map[models.Position]bool,0)
	// Left side move
	posMap = p.getMovesDiagonal(posMap, board, curPos,  p.dRow, -1)
	// Left side move
	posMap = p.getMovesDiagonal(posMap, board, curPos,  p.dRow, +1)
	// Forward move
	posMap = p.getMovesForward(posMap, board, curPos,  p.dRow)
	// First Double move
	if p.isFirstMove() {
		posMap = p.getMovesForward(posMap, board, curPos,  2*p.dRow)
	}

	return posMap
}

func (p *Pawn)getMovesForward(posMap map[models.Position]bool, board models.Board, currPos *models.Position, dRow int)map[models.Position]bool{
	row := currPos.GetRow() + dRow
	col := currPos.GetCol()
	if models.ValidatePosition(row , col, board.GetBoardSize()) {
		newPos := models.NewPosition(row,col)
		_, ok := board.IsPositionFilled(newPos)
		if !ok {
			posMap[*newPos] = true
		}
	}
	return posMap
}

func (p *Pawn)getMovesDiagonal(posMap map[models.Position]bool, board models.Board, currPos *models.Position, dRow int, dCol int)map[models.Position]bool{
	row := currPos.GetRow() + dRow
	col := currPos.GetCol() + dCol
	if models.ValidatePosition(row , col, board.GetBoardSize()) {
		newPos := models.NewPosition(row,col)
		piece, ok := board.IsPositionFilled(newPos)
		if ok && piece.GetColor() != p.color{
			posMap[*newPos] = true
		}
	}
	return posMap
}

func (p *Pawn)IsSetNewPiece(newPos *models.Position)bool{
	if p.isFirstMove() {
		p.disableFirstMove()
		return false
	}
	return p.isLastMove(newPos)
}

func (p *Pawn)isLastMove(pos *models.Position) bool {
	return p.destinationRow == pos.GetRow()
}

func (p *Pawn)GetColor()models.Color{
	return p.color
}

func (p *Pawn)GetName()models.PieceName{
	return models.PawnName
}