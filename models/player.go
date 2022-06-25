package models

type Player struct {
	user *User
	color 	Color
}

func NewPlayer(user *User, color Color)*Player{
	return &Player{
		user: user,
		color: color,
	}
}

func (p *Player)GetUser()*User{
	return p.user
}

func (p *Player)GetColor()Color{
	return p.color
}
