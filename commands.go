package main

import (
	"Chess/models"
	"fmt"
	"strconv"
	"strings"
)

func RunCommands(manager *Game){
	commands := getCommands()
	manager.PrintBoard()
	for _,command := range commands {
		curPos, newPos := parseCommand(command, manager.GetBoardSize())
		piece, ok := manager.chessBoard.IsPositionFilled(curPos)
		currPlayer := manager.GetCurrentPlayer()
		fmt.Println(currPlayer.GetUser().GetName()+ " " + string(currPlayer.GetColor()) + " "+ command)
		if !ok || piece.GetColor() != currPlayer.GetColor(){
			fmt.Printf("Invalid Move\n")
			continue
		}

		if err := manager.SetPiece(curPos, newPos, piece); err != nil {
			fmt.Println(err)
			continue
		}
		//fmt.Println(i)
		manager.PrintBoard()
		manager.NextTurn()
	}
}

func parseCommand(command string, size int) (cur *models.Position, next *models.Position) {
	cmd := strings.Split(command, " ")
	cCol,_ := strconv.Atoi(fmt.Sprintf("%v",cmd[0][0]-'a'))
	cRow,_ := strconv.Atoi(string(cmd[0][1]))
	nCol,_ := strconv.Atoi(fmt.Sprintf("%v",cmd[1][0]-'a'))
	nRow,_ := strconv.Atoi(string(cmd[1][1]))
	cur = models.NewPosition(size - cRow, cCol)
	next = models.NewPosition(size - nRow, nCol)
	return
}

func getCommands()[]string{
	commands := []string{
		"e2 e4",
		"e7 e5",
		"f1 c4",
		"b8 c6",
		"d1 h5",
		"g8 f6",
		"h5 f7",
		"f8 f7",
		"g7 f7",
		"h8 f7",
		"d8 f7",
		"c6 f7",
		"c4 f7",
		"h8 g8",
		"f2 f4",
		"e5 f4",
		"f7 e8",
	}
	return commands
}