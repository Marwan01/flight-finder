package main

import (
	"fmt"
)

func main() {
	// fmt.Println("welcome to first database project, Enter command, and click enter at the end of each command")
	// readFirstCommand()
	// fmt.Println()
	// fileHandle, _ := os.Open("data.txt")
	// defer fileHandle.Close()
	// fileScanner := bufio.NewScanner(fileHandle)

	// for fileScanner.Scan() {
	// 	fmt.Println(fileScanner.Text())
	// }

	var cityCode string
	var numberOfInputs int
	numberOfInputs, _ = fmt.Scan(&cityCode)
	fmt.Println(cityCode)
	fmt.Println(numberOfInputs)
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

func add(secondCommand string) {
	if secondCommand == "c" {
		addC()
	} else if secondCommand == "a" {
		addA()
	} else {
		addF()
	}
}

func addC() {
	var cityCode []string
	var cityName = ""
	fmt.Scan(&cityCode)
	fmt.Scanln(&cityName)
}

func addA() {
	fmt.Println("yay made it to a, a")

}

func addF() {

}
