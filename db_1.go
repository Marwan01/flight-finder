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
	defer file.Close()
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
		// var aa, an string
		var startcc, endcc string
		var connections int
		cc = commands[2]
		cn = commands[3] + commands[4]
		if firstCommand == "a" {

			if secondCommand == "c" {
				AddC(cc, cn, file)
			} else if secondCommand == "a" {
				AddA(cc, cn, file)
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
	fmt.Fprintln(f, cityCode+" "+cityName)
}

// AddA adds an airport
func AddA(aa string, an string, f *os.File) {
	fmt.Fprintln(f, aa+" "+an)
}

// AddF adds a flight
func AddF(aa, cc1, cc2, p string, f *os.File) {
	fmt.Fprintln(f, aa+" "+cc1+" "+cc2+" "+p)
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
