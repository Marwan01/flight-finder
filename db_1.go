package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// for _, singleCommand := range commands {

	// 	// good idea: use commad[1], command[2], command[3], concat( rest)
	// }

	// var namesOfCities = commands[2:]
	// fmt.Println(namesOfCities)

	for {
		fmt.Println("\nEnter your Full Command: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		var commands []string
		commands = strings.Split(input, " ")
		var firstCommand = commands[0]
		var secondCommand = commands[1]
		var cc, cn string // city code and city name
		var startcc, endcc string
		var connections int
		if firstCommand == "a" {

			if secondCommand == "c" {
				AddC(cc, cn)
			} else if secondCommand == "a" {
				AddA()
			} else if secondCommand == "f" {
				AddF()
			} else {
				fmt.Println("wrong second command, try again")
			}

		} else if firstCommand == "l" {

			if secondCommand == "c" {
				LoadC()
			} else if secondCommand == "a" {
				LoadA()
			} else if secondCommand == "f" {
				LoadF()
			} else {
				fmt.Println("wrong second command, try again")
			}

		} else if firstCommand == "f" {
			Find(startcc, endcc, connections)
			return
		} else if firstCommand == "q" {
			// quit
			os.Exit(3)
		} else {
			fmt.Println("wrong first command, try again")
		}
	}

}

func AddC(cc string, cn string) {
	var cityCode = cc
	var cityName = cn
	fmt.Scan(&cityCode)
	fmt.Scanln(&cityName)
}

func AddA() {
	fmt.Println("yay made it to a, a")

}

func AddF() {

}

func LoadA() {

}

func LoadC() {

}
func LoadF() {

}

func Find(startcc string, endcc string, connections int) {

}
