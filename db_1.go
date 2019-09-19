package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Create("data.txt") // create a new file under the current directory
	// and call it data.txt (overwrites if it exits)
	if err != nil {
		log.Fatal("Cannot create file", err) // log errors with file creation if any
	}
	defer file.Close() // close the file, but only after the program is done with it
	for {              // loop to keep asking the user for commands until q is entered.

		fmt.Println("\nEnter your Full Command: ")
		reader := bufio.NewReader(os.Stdin)   // create new reader to read input from terminal
		text, _ := reader.ReadString('\n')    // read until \n (enter key pressed)
		text = strings.TrimSuffix(text, "\n") // golang adds \n to every string. I need to trim it so the last command would not include the characters \n
		var commands []string                 // slice of strings
		commands = strings.Split(text, " ")   // holds strings seperated by spaces in slice
		var firstCommand = commands[0]
		if firstCommand == "a" {
			var secondCommand = commands[1]
			var cc, cn string
			cc = commands[2]
			for i := 3; i < len(commands); i++ {
				cn = cn + " " + commands[i]
			}
			if secondCommand == "c" {
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
			var startcc, endcc string // starting city code, ending city code
			var connections string    // number of connections
			startcc = commands[1]
			endcc = commands[2]
			connections = commands[3]
			Find(startcc, endcc, connections, file)
		} else if firstCommand == "q" {
			fmt.Println("Goodbye!")
			os.Exit(3) // quit with exit code 3
		} else {
			fmt.Println("wrong first command, try again")
		}
	}
}

// AddC adds a city to the data file passed in as a parameter in this format: c: citycode this is the city name
// takes in two strings representing the city code and city name to be written to the file
// and takes in a file pointer to which file to write the data to.
func AddC(cc string, cn string, f *os.File) {
	var cityCode = cc
	var cityName = cn
	fmt.Fprintln(f, "c: "+cityCode+" "+cityName)
}

// AddA adds an airport to the data file passed in as a parameter in this format: a: airport this is the airport name
// takes in two strings representing the airport abbreviation and airport name to be written to the file
// and takes in a file pointer to which file to write the data to.
func AddA(aa string, an string, f *os.File) {
	fmt.Fprintln(f, "a: "+aa+" "+an)
}

// AddF adds a flight to the data file passed in as a parameter in this format: f: airlineAbbreviation startingCityCode arrivalCityCode priceOfTheFlight
// takes in four strings representing the starting city code, arrival city code, and price of the flight to be written to the file
// and takes in a file pointer to which file to write the data to.
func AddF(aa, cc1, cc2, p string, f *os.File) {
	fmt.Fprintln(f, "f: "+aa+" "+cc1+" "+cc2+" "+p)
}

// LoadC take in a file pointer as a parameter and outputs to the console all the cities on that file
func LoadC(f *os.File) {
	f, _ = os.Open("data.txt") // opens file for reading
	b, _ := ioutil.ReadAll(f)  // read all content of the file into a bytes array
	s := string(b)             //convert from []byte to string
	var arr []string           //create a new string array of strings that will have as elements the lines read from the file
	arr = strings.Split(s, "\n")
	var str []string // slice of strings to hold every single space seperated string collected from each line read from the file
	for i := 0; i < len(arr); i++ {
		str = strings.Split(arr[i], " ")
		if str[0] == "c:" { // check if a city, indicated by starting string: "c: " is found
			fmt.Print(str[1] + " ")
			for j := 2; j < len(str); j++ {
				fmt.Print(str[j] + " ")
			}
		}
		fmt.Print("\n")
	}
}

// LoadA take in a file pointer as a parameter and outputs to the console all the airports on that file
func LoadA(f *os.File) {
	f, _ = os.Open("data.txt") // opens file for reading
	b, _ := ioutil.ReadAll(f)  // read all content of the file into a bytes array
	s := string(b)             //convert from []byte to string
	var arr []string           //create a new string array of strings that will have as elements the lines read from the file
	arr = strings.Split(s, "\n")
	var str []string // slice of strings to hold every single space seperated string collected from each line read from the file
	for i := 0; i < len(arr); i++ {
		str = strings.Split(arr[i], " ")
		if str[0] == "a:" { // check if an airport, indicated by starting string: "a: " is found
			fmt.Print(str[1] + " ")
			for j := 2; j < len(str); j++ { // print the rest of the string
				fmt.Print(str[j] + " ")
			}
		}
		fmt.Print("\n")
	}
}

// LoadF take in a file pointer as a parameter and outputs to the console all the flights on that file
func LoadF(f *os.File) {
	f, _ = os.Open("data.txt") // opens file for reading
	b, _ := ioutil.ReadAll(f)  // read all content of the file into a bytes array
	s := string(b)             //convert from []byte to string
	var arr []string           //create a new string array of strings that will have as elements the lines read from the file
	arr = strings.Split(s, "\n")
	var str []string // slice of strings to hold every single space seperated string collected from each line read from the file
	for i := 0; i < len(arr); i++ {
		str = strings.Split(arr[i], " ")
		if str[0] == "f:" { // check if a city, indicated by starting string: "c: " is found
			fmt.Println(str[1] + ": " + str[2] + " ->" + str[3] + " $" + str[4]) // print chunks of flight line
		}
	}
}

// Find finds flights according to entries of starting city and
// arrival city and number of connections. Also takes in a file as input
// and reads the data from it
func Find(startcc string, endcc string, connections string, f *os.File) {
	f, _ = os.Open("data.txt") // opens file for reading
	b, _ := ioutil.ReadAll(f)  // read all content of the file into a bytes array
	s := string(b)             //convert from []byte to string
	var arr []string           //create a new string array of strings that will have as elements the lines read from the file
	arr = strings.Split(s, "\n")
	if connections == "0" { // when we want no connections as in direct flight
		for i := 0; i < len(arr); i++ {
			var aa []string // string array to temporarily hold splits of a string line
			aa = strings.Split(arr[i], " ")
			if aa[0] == "f:" { // if the first line indicates a flight
				if startcc == aa[2] && endcc == aa[3] {
					fmt.Println(startcc + " -> " + endcc + " : " + aa[1] + " $" + aa[4])
				}
			}
		}
	}
	if connections == "1" { // if we want a flight with one connection
		for i := 0; i < len(arr); i++ {
			if strings.Contains(arr[i], "f:") { // if we have a flight line,
				var start []string // string array to temporarily hold splits of a string line
				start = strings.Split(arr[i], " ")
				if startcc == start[2] {
					for j := 0; j < len(arr); j++ {
						var arrival []string
						arrival = strings.Split(arr[j], " ")
						if strings.Contains(arr[j], "f:") {
							if strings.Contains(endcc, arrival[3]) {
								if strings.Contains(start[3], arrival[2]) {
									var finalPrice int                                                                                                                                                                                    // integer denoting final cost of two flights coupled together
									p1, _ := strconv.Atoi(start[4])                                                                                                                                                                       // convert to number
									p2, _ := strconv.Atoi(arrival[4])                                                                                                                                                                     // convert to number
									finalPrice = p1 + p2                                                                                                                                                                                  // add together
									var finalPriceStr string                                                                                                                                                                              // string to hold the final price to be outputted
									finalPriceStr = strconv.Itoa(finalPrice)                                                                                                                                                              // convert final price to string
									fmt.Println(startcc + " -> " + start[3] + " : " + start[1] + " $" + start[4] + "; " + start[3] + " -> " + endcc + " : " + arrival[1] + " $" + arrival[4] + ", for a total cost of $" + finalPriceStr) // output
								}
							}
						}
					}
				}

			}

		}
	}
}
