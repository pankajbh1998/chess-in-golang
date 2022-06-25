package piece

import (
	"Chess/models"
)

// Hathi
type Rook struct{
	color models.Color
}

func NewRook(color models.Color)*Rook{
	return &Rook{
		color: color,
	}
}

func (r *Rook)GetValidMoves(board models.Board, curPos *models.Position)map[models.Position]bool{
	posMap := make(map[models.Position]bool,0)
	// Get Top Moves
	posMap = r.getMoves(posMap, board, curPos,  1, 0)
	// Get Down Moves
	posMap = r.getMoves(posMap, board, curPos,  -1, 0)
	// Get Right Moves
	posMap = r.getMoves(posMap, board, curPos,  0, 1)
	// Get Left Moves
	posMap = r.getMoves(posMap, board, curPos,  0, -1)

	return posMap
}

func (r *Rook)getMoves(posMap map[models.Position]bool, board models.Board, currPos *models.Position, dRow int, dCol int)map[models.Position]bool{
	row := currPos.GetRow() + dRow
	col := currPos.GetCol() + dCol
	for models.ValidatePosition(row , col, board.GetBoardSize()) {
		newPos := models.NewPosition(row,col)
		piece, ok := board.IsPositionFilled(newPos)
		if ok {
			if piece.GetColor() != r.color {
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

func (r *Rook)GetColor()models.Color{
	return r.color
}

func (r *Rook)GetName()models.PieceName{
	return models.RookName
}

func (r *Rook)IsSetNewPiece(newPos *models.Position)bool{
	return false
}