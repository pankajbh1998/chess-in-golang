package piece

import (
	"Chess/models"
)

type King struct {
	color models.Color
}

func NewKing(color models.Color)*King{
	return &King{
		color: color,
	}
}

func (k *King)GetValidMoves(b models.Board, curPos *models.Position)map[models.Position]bool{
	posMap := make(map[models.Position]bool,0)
	// Get Top Moves
	posMap = k.getMoves(posMap, b, curPos, 1, 0 )
	// Get Down Moves
	posMap = k.getMoves(posMap, b, curPos, -1, 0)
	// Get Left Moves
	posMap = k.getMoves(posMap, b, curPos, 0, -1)
	// Get Right Moves
	posMap = k.getMoves(posMap, b, curPos, 0, -1)
	// Get Top Left Diagonal Moves
	posMap = k.getMoves(posMap, b, curPos, 1, -1)
	// Get Top Right Diagonal Moves
	posMap = k.getMoves(posMap, b, curPos, 1, 1)
	// Get Down Left Diagonal Moves
	posMap = k.getMoves(posMap, b, curPos, -1, -1)
	// Get Down Right DiagonalMoves
	posMap = k.getMoves(posMap, b, curPos, -1, 1)

	return posMap
}

func (k *King)getMoves(posMap map[models.Position]bool, b models.Board, currPos *models.Position, dRow int, dCol int)map[models.Position]bool{
	row := currPos.GetRow() + dRow
	col := currPos.GetCol() + dCol
	if models.ValidatePosition(row , col, b.GetBoardSize()) {
		newPos := models.NewPosition(row,col)
		piece, ok := b.IsPositionFilled(newPos)
		if ok {
			if piece.GetColor() != k.color {
				posMap[*newPos] = true
			}
			return posMap
		}

		posMap[*newPos] = true
	}
	return posMap
}

func (k *King)GetColor()models.Color{
	return k.color
}

func (k *King)GetName()models.PieceName{
	return models.KingName
}

func (k *King)IsSetNewPiece(newPos *models.Position)bool{
	return false
}
