package piece

import (
	"Chess/models"
)

type Queen struct{
	color models.Color
}

func NewQueen(color models.Color)*Queen{
	return &Queen{
		color: color,
	}
}

func (q *Queen)GetValidMoves(board models.Board, curPos *models.Position)map[models.Position]bool{
	posMap := make(map[models.Position]bool,0)
	// Get Top Moves
	posMap = q.getMoves(posMap, board, curPos,  1, 0)
	// Get Down Moves
	posMap = q.getMoves(posMap, board, curPos,  -1, 0)
	// Get Left Moves
	posMap = q.getMoves(posMap, board, curPos,  0, -1)
	// Get Right Moves
	posMap = q.getMoves(posMap, board, curPos,  0, -1)
	// Get Top Left Diagonal Moves
	posMap = q.getMoves(posMap, board, curPos,  1, -1)
	// Get Top Right Diagonal Moves
	posMap = q.getMoves(posMap, board, curPos,  1, 1)
	// Get Down Left Diagonal Moves
	posMap = q.getMoves(posMap, board, curPos,  -1, -1)
	// Get Down Right DiagonalMoves
	posMap = q.getMoves(posMap, board, curPos,  -1, 1)

	return posMap
}

func (q *Queen)getMoves(posMap map[models.Position]bool, board models.Board, currPos *models.Position, dRow int, dCol int)map[models.Position]bool{
	row := currPos.GetRow() + dRow
	col := currPos.GetCol() + dCol
	for models.ValidatePosition(row , col, board.GetBoardSize()) {
		newPos := models.NewPosition(row,col)
		piece, ok := board.IsPositionFilled(newPos)
		if ok {
			if piece.GetColor() != q.color {
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


func (q *Queen)IsSetNewPiece(newPos *models.Position)bool{
	return false
}

func (q *Queen)GetColor()models.Color{
	return q.color
}

func (q *Queen)GetName()models.PieceName{
	return models.QueenName
}