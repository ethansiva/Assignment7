package main

import (
	"assignment7/directory"
	"assignment7/website"
	"fmt"
)

func main(){
	var userInput int 
	fmt.Println("Press 1 for a Static Website\nPress 2 for a File Server for a directory")
	fmt.Scanln(&userInput)
	switch userInput {
	case 1:
		website.StaticFunc()
		break
	case 2:
		directory.DirectoryFunc()
		break
	}
}