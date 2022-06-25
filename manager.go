package main

import (
	"Chess/models"
	"Chess/repo"
	"Chess/piece"
	"fmt"
)

type Game struct {
	chessBoard *repo.ChessBoard
	players []*models.Player
	currentPlayer int
	gameState models.GameState
}

func NewGameManager(c *repo.ChessBoard, players []*models.Player)*Game{
	return &Game{
		chessBoard: c,
		players: players,
		currentPlayer: 0,
		gameState: models.NewGame,
	}
}

func (g *Game)SetPiece(curPos *models.Position, newPos *models.Position, piece models.Piece)error{
	validMove := piece.GetValidMoves(g.chessBoard, curPos)
	if val,ok := validMove[*newPos];  !ok || val == false {
		return fmt.Errorf("Invalid Move")
	}
	// Last pos for Pawn
	if piece.IsSetNewPiece(newPos) {
		piece = g.getNewPiece(piece.GetColor())
	}
	// Set Piece to new Pos
	g.chessBoard.RemovePiece(curPos)
	g.chessBoard.SetPiece(newPos, piece)
	return nil
}

func (g *Game)PrintBoard(){
	for i:=0;i<g.chessBoard.GetBoardSize();i++ {
		str := ""
		for j:=0;j<g.chessBoard.GetBoardSize();j++{
			piece, ok := g.chessBoard.IsPositionFilled(models.NewPosition(i,j))
			if ok {
				str = fmt.Sprintf("%v %v%v", str, piece.GetColor(), piece.GetName())
			} else {
				str = fmt.Sprintf("%v --", str)
			}
		}
		fmt.Println(str)
	}
	fmt.Println()
}

func(g *Game) NextTurn(){
	n := len(g.players)
	g.currentPlayer = (g.currentPlayer + 1) % n
}

func (g *Game)GetCurrentPlayer()*models.Player{
	return g.players[g.currentPlayer]
}

func (g *Game)getNewPiece(color models.Color)models.Piece{
	fmt.Printf("What do you want as new Piece: \nPress Q for Queen\nPress N for Knight\nPress B for Bishop\nPress R for Rook")
	for {
		switch models.PieceName(models.GetString()) {
		case models.KnightName:
			return piece.NewKnight(color)
		case models.QueenName:
			return piece.NewQueen(color)
		case models.BishopName:
			return piece.NewBishop(color)
		case models.RookName:
			return piece.NewRook(color)
		default:
			fmt.Println("Wrong Choice")
		}
	}
	return nil
}


func (g *Game)GetBoardSize()int{
	return g.chessBoard.GetBoardSize()
}

