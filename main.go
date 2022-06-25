package main

import (
	"Chess/models"
	"Chess/piece"
	"Chess/repo"
	"fmt"
)

func main(){
	boardSize := 8
	chessBoard := getNewChessBoard(boardSize)
	players := getPlayers()
	if len(players) < 2 {
		fmt.Println("More than two players are required")
		return
	}
	gameManager := NewGameManager(chessBoard, players)
	RunCommands(gameManager)
}

func getPlayers()[]*models.Player{
	players := make([]*models.Player,0)
	players = append(players, models.NewPlayer(models.NewUser("Pankaj"), models.White))
	players = append(players, models.NewPlayer(models.NewUser("Shivam"), models.Black))
	return players
}

func getNewChessBoard(boardSize int)*repo.ChessBoard{
	chessboard := repo.NewChessBoard(boardSize)
	setBlackPieces(chessboard)
	setWhitePieces(chessboard)
	return chessboard
}

func setBlackPieces(board *repo.ChessBoard){
	color := models.Black
	board.SetPiece(models.NewPosition(0,0),piece.NewRook(color))
	board.SetPiece(models.NewPosition(0,1),piece.NewKnight(color))
	board.SetPiece(models.NewPosition(0,2),piece.NewBishop(color))
	board.SetPiece(models.NewPosition(0,3),piece.NewQueen(color))
	board.SetPiece(models.NewPosition(0,4),piece.NewKing(color))
	board.SetPiece(models.NewPosition(0,5),piece.NewBishop(color))
	board.SetPiece(models.NewPosition(0,6),piece.NewKnight(color))
	board.SetPiece(models.NewPosition(0,7),piece.NewRook(color))

	// For Pawns
	destinationRow := 7
	dRow := 1
	for i:=0;i<board.GetBoardSize();i++ {
		board.SetPiece(models.NewPosition(1,i),piece.NewPawn(color, destinationRow, dRow))
	}
}

func setWhitePieces(board *repo.ChessBoard){
	color := models.White
	board.SetPiece(models.NewPosition(7,0),piece.NewRook(color))
	board.SetPiece(models.NewPosition(7,1),piece.NewKnight(color))
	board.SetPiece(models.NewPosition(7,2),piece.NewBishop(color))
	board.SetPiece(models.NewPosition(7,3),piece.NewQueen(color))
	board.SetPiece(models.NewPosition(7,4),piece.NewKing(color))
	board.SetPiece(models.NewPosition(7,5),piece.NewBishop(color))
	board.SetPiece(models.NewPosition(7,6),piece.NewKnight(color))
	board.SetPiece(models.NewPosition(7,7),piece.NewRook(color))

	// For Pawns
	destinationRow := 0
	dRow := -1
	for i:=0;i<board.GetBoardSize();i++ {
		board.SetPiece(models.NewPosition(6,i),piece.NewPawn(color, destinationRow, dRow))
	}
}
