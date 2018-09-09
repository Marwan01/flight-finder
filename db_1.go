package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
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
		var secondCommand = commands[1] // need fix for when q is entered ( dont have second command, will complain)
		var cc, cn string               // city code and city name
		var startcc, endcc string
		var connections int
		cc = commands[2]
		cn = commands[3] + commands[4]
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

// AddC adds a city
func AddC(cc string, cn string, f *os.File) {
	var cityCode = cc
	var cityName = cn
	fmt.Println(cityCode)
	fmt.Println(cityName)
	defer file.Close()
	fmt.Fprintln(file, cityCode+" "+cityName)
}

// AddA adds an airport
func AddA(f *os.File) {
	fmt.Println("yay made it to a, a")

}

// AddF adds a flight
func AddF() {

}

// LoadC loads cities
func LoadC() {

}

// LoadA loads airports
func LoadA() {

}

// LoadF loads flights
func LoadF() {

}

// Find finds flights according to entries of starting city and
// finish city and number of connections
func Find(startcc string, endcc string, connections int) {

}
