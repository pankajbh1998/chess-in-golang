package piece

import (
	"Chess/models"
)

type Knight struct{
	color models.Color
}

func NewKnight(color models.Color)*Knight{
	return &Knight{
		color: color,
	}
}

func (k *Knight)GetValidMoves(board models.Board, curPos *models.Position)map[models.Position]bool{
	posMap := make(map[models.Position]bool,0)
	// Get Top Right Moves
	posMap = k.getMoves(posMap, board, curPos,  2, 1)
	// Get Top Left Moves
	posMap = k.getMoves(posMap, board, curPos,  2, -1)
	// Get Down Right Moves
	posMap = k.getMoves(posMap, board, curPos,  -2, 1)
	// Get Down Left Moves
	posMap = k.getMoves(posMap, board, curPos,  -2, -1)
	// Get Right Top Moves
	posMap = k.getMoves(posMap, board, curPos,  1, 2)
	// Get Right Down Moves
	posMap = k.getMoves(posMap, board, curPos,  -1, 2)
	// Get Left Top Moves
	posMap = k.getMoves(posMap, board, curPos,  1, -2)
	// Get Left Down Moves
	posMap = k.getMoves(posMap, board, curPos,  -1, -2)

	return posMap
}

func (k *Knight)getMoves(posMap map[models.Position]bool, b models.Board, currPos *models.Position, dRow int, dCol int)map[models.Position]bool{
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

func (k *Knight)GetColor()models.Color{
	return k.color
}

func (k *Knight)GetName()models.PieceName{
	return models.KnightName
}

func (k *Knight)IsSetNewPiece(newPos *models.Position)bool{
	return false
}