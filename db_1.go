package main

import "fmt"

func main() {
	readFirstCommand()

}

func add(string) {

}

func readFirstCommand() {
	var firstCommand = ""
	for {
		fmt.Scanln(&firstCommand)
		if firstCommand == "a" {
			add(readSecondCommand())
			return
		} else if firstCommand == "l" {
			// list
			return
		} else if firstCommand == "f" {
			// find
			return
		} else if firstCommand == "q" {
			// quit
			return
		} else {
			fmt.Println("wrong entry, try again")
		}
	}
}

func readSecondCommand() string {
	var secondCommand = ""
	for {
		fmt.Scanln(&secondCommand)
		if secondCommand == "c" {
			return "c"
		} else if secondCommand == "a" {
			return "a"
		} else if secondCommand == "f" {
			return "f"
		} else {
			fmt.Println("wrong entry, try again")
		}
	}

}
