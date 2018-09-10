package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("data.txt") // assumes a file already exits
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()
	for {

		fmt.Println("\nEnter your Full Command: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		var commands []string
		commands = strings.Split(text, " ")
		var firstCommand = commands[0]
		if firstCommand == "a" {
			var secondCommand = commands[1]
			var cc, cn string
			cc = commands[2]
			for i := 0; i < len(commands); i++ {
				cn = cn + " " + commands[i]
			}
			if strings.Contains(secondCommand, "c") {
				AddC(cc, cn, file)
			} else if secondCommand == "a" {
				AddA(cc, cn, file)
			} else if secondCommand == "f" {
				aa := commands[2] // airport abbreviation
				dc := commands[3] // departure city
				ac := commands[4] // arrival city
				p := commands[5]  // price
				AddF(aa, dc, ac, p, file)
			} else {
				fmt.Println("wrong second command, try again")
			}

		} else if firstCommand == "l" {
			var secondCommand = commands[1]
			if secondCommand == "c" {
				LoadC(file)
			} else if secondCommand == "a" {
				LoadA(file)
			} else if secondCommand == "f" {
				LoadF(file)
			} else {
				fmt.Println("wrong second command, try again")
			}

		} else if firstCommand == "f" {
			var startcc, endcc string
			var connections string
			startcc = commands[1]
			endcc = commands[2]
			connections = commands[3]
			Find(startcc, endcc, connections)
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
	fmt.Fprintln(f, "c: "+cityCode+" "+cityName)
}

// AddA adds an airport
func AddA(aa string, an string, f *os.File) {
	fmt.Fprintln(f, "a: "+aa+" "+an)
}

// AddF adds a flight
func AddF(aa, cc1, cc2, p string, f *os.File) {
	fmt.Fprintln(f, "f: "+aa+" "+cc1+" "+cc2+" "+p)
}

// LoadC loads cities
func LoadC(f *os.File) {
	f, _ = os.Open("data.txt")
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	var arr []string
	arr = strings.Split(s, "\n")
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "c: ") {
			fmt.Println(arr[i])
		}
	}
}

// LoadA loads airports
func LoadA(f *os.File) {
	f, _ = os.Open("data.txt")
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	var arr []string
	arr = strings.Split(s, "\n")
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "a: ") {
			fmt.Println(arr[i])
		}
	}
}

// LoadF loads flights
func LoadF(f *os.File) {
	f, _ = os.Open("data.txt")
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	var arr []string
	arr = strings.Split(s, "\n")
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "f: ") {
			fmt.Println(arr[i])
		}
	}
}

// Find finds flights according to entries of starting city and
// finish city and number of connections
func Find(startcc string, endcc string, connections string, f *os.File) {
	f, _ = os.Open("data.txt")
	b, _ := ioutil.ReadAll(f)
	s := string(b)
	var arr []string
	arr = strings.Split(s, "\n")
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "f: ") {
			fmt.Println(arr[i])
		}
	
}
