package models

import (
	"bufio"
	"os"
	"strings"
)

const BoardSize = 8


func ValidatePosition(row int , col int, boardSize int)bool{
	if row < 0 || col < 0 || row >= boardSize || col >= boardSize {
		return false
	}

	return true
}

func GetString()string{
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	return strings.Replace(str, "\n","",-1)
}

