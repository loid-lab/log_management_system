package main

import (
	"fmt"
	"os"
)

func createLogDirectory() {
	fmt.Println("Enter the name of the log directory you'd like to create:")
	var dirName string
	fmt.Scanln(&dirName)

	err := os.Mkdir("user_profiles", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Directory successfully created.")
}

func changeLogDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Directory:", cwd)

	fmt.Print("Enter the name of the log directory you want to change to: ")
	var newDirName string
	fmt.Scanln(&newDirName)

	err := os.Chdir("user_profiles/new_user")
	if err != nil {
		fmt.Println(err)
		return
	}

	newCwd := getWorkingDirectory()
	fmt.Println("Directory changed successfully. Current Directory:", newCwd)
}

func printWorkingDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Working Directory:", cwd)
}

func listDirectoryContents() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Working Directory:", cwd)

	files, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	fmt.Println("Directory Contents:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
