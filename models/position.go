package models

type Position struct{
	row int
	col int
}

func NewPosition(row int, col int)*Position{
	return &Position{
		row: row,
		col: col,
	}
}

func (p *Position)GetRow()int{
	return p.row
}

func (p *Position)GetCol()int{
	return p.col
}

func (p *Position)SetRow(row int){
	p.row = row
}

func (p *Position)SetCol(col int){
	p.col = col
}


