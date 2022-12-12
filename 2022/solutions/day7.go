package solutions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/whcass/Advent-of-Code/2022/entities"
	"github.com/whcass/Advent-of-Code/2022/helpers"
)

type Day7 struct {
	Title      string
	fileSystem []*entities.Directory
}

func NewDay7() *Day7 {
	return &Day7{
		Title:      "Day 7",
		fileSystem: make([]*entities.Directory, 1),
	}
}

func (d Day7) GetTitle() string {
	return d.Title
}

func (d Day7) Run() {
	input := helpers.GetInput("7test")
	sliceData := strings.Split(input, "\n")
	d.ConstructFileSystem(sliceData)
}

func (d Day7) ConstructFileSystem(data []string) {
	currentDirectory := entities.NewDirectory("/", nil)
	directories := make([]*entities.Directory, 1)
	directories[0] = currentDirectory
	d.addToFileSystem(currentDirectory)
	for i := 1; i < len(data); i++ {
		command := strings.Split(data[i], " ")
		// fmt.Printf("[*] command: %v\n", command)
		if command[0] == "$" {
			switch command[1] {
			case "cd":
				if command[2] == ".." {
					currentDirectory = currentDirectory.Parent
				} else {
					currentDirectory = currentDirectory.FindDirectory(command[2])
				}
			case "ls":
				continue
			}
		} else if command[0] == "dir" {

			dir := entities.NewDirectory(command[1], currentDirectory)
			directories = append(directories, dir)
			currentDirectory.AddDirectory(dir)
		} else {
			fileSize, _ := strconv.Atoi(command[0])
			file := entities.NewFile(currentDirectory, command[1], fileSize)
			currentDirectory.AddFile(file)
		}
	}
	for j := 0; j < len(directories); j++ {
		fmt.Printf("[*] Name: %s - Size: %d\n", directories[j].Name, directories[j].Size)
	}
}

func (d *Day7) addToFileSystem(dir *entities.Directory) {
	d.fileSystem = append(d.fileSystem, dir)
}
