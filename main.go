package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("\nWelcome to the Log Manager!")
		fmt.Println("Choose an operation:")
		fmt.Println("1. Create Log Directory")
		fmt.Println("2. Change Log Directory")
		fmt.Println("3. Print Current Working Directory")
		fmt.Println("4. List Directory Contents")
		fmt.Println("5. Create Log File")
		fmt.Println("6. Read Log File")
		fmt.Println("7. Write Log Entry")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter your choice (1-8): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			createLogDirectory()
		case 2:
			changeLogDirectory()
		case 3:
			printWorkingDirectory()
		case 4:
			listDirectoryContents()
		case 5:
			createLogFile()
		case 6:
			readLogFile()
		case 7:
			writeLogEntry()
		case 8:
			fmt.Println("Exiting the Log Manager. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Please select a valid option.")
		}
	}
}

func getWorkingDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return cwd
}
