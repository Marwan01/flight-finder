package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("\nEnter your Full Command: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	var commands []string
	commands = strings.Split(input, " ")
	// for _, singleCommand := range commands {

	// 	// good idea: use commad[1], command[2], command[3], concat( rest)
	// }

	var namesOfCities = commands[2:]
	fmt.Println(namesOfCities)
	var firstCommand = commands[0]
	for {
		if firstCommand == "a" {
			// Add(ValidateSecondCommand(secondCommand))
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
func ExecuteFirstCommand(c1 string, c2 string) {
	var firstCommand = c1
	//var secondCommand = c2
	for {
		if firstCommand == "a" {
			// Add(ValidateSecondCommand(secondCommand))
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

func ValidateSecondCommand(c2 string) string {
	var secondCommand = c2
	for {
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

func Add(secondCommand string, cc string, cn string) {
	if secondCommand == "c" {
		addC(cc, cn)
	} else if secondCommand == "a" {
		addA()
	} else {
		addF()
	}
}

func addC(cc string, cn string) {
	var cityCode = cc
	var cityName = cn
	fmt.Scan(&cityCode)
	fmt.Scanln(&cityName)
}

func addA() {
	fmt.Println("yay made it to a, a")

}

func addF() {

}
