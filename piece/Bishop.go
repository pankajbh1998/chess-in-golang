package piece

import (
	"Chess/models"
)
//Camel
type Bishop struct{
	color models.Color
}

func NewBishop(color models.Color)*Bishop{
	return &Bishop{
		color: color,
	}
}

func (b *Bishop)GetValidMoves(board models.Board, curPos *models.Position)map[models.Position]bool{
	posMap := make(map[models.Position]bool,0)
	// Get top Right Moves
	posMap = b.getMoves(posMap, board, curPos,  1, -1)
	// Get down Right Moves
	posMap = b.getMoves(posMap, board, curPos,  -1, -1)
	// Get top Left Moves
	posMap = b.getMoves(posMap, board, curPos,  1, 1)
	// Get down Left Moves
	posMap = b.getMoves(posMap, board, curPos,  -1, 1)

	return posMap
}

func (b *Bishop)getMoves(posMap map[models.Position]bool, board models.Board, currPos *models.Position, dRow int, dCol int)map[models.Position]bool{
	row := currPos.GetRow() + dRow
	col := currPos.GetCol() + dCol
	for models.ValidatePosition(row , col, board.GetBoardSize()) {
		newPos := models.NewPosition(row,col)
		piece, ok := board.IsPositionFilled(newPos)
		if ok {
			if piece.GetColor() != b.color {
				posMap[*newPos] = true
			}
			return posMap
		}

		posMap[*newPos] = true
		row += dRow
		col += dCol
	}
	return posMap
}

func (b *Bishop)GetColor()models.Color{
	return b.color
}

func (b *Bishop)GetName()models.PieceName{
	return models.BishopName
}

func (b *Bishop)IsSetNewPiece(newPos *models.Position)bool{
	return false
}