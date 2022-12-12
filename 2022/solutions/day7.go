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
	input := helpers.GetInput("7")
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
	updateDirSize(d.fileSystem[1])
	total := 0
	for j := 0; j < len(directories); j++ {
		if directories[j].Size < 100000 {
			total += directories[j].Size
		}
	}
	helpers.PrintResult(fmt.Sprintf("Part 1: %d", total))

	// Probably don't need all of this here, but it helps conseptualise it in my brain
	totalUsedSpace := d.fileSystem[1].Size
	totalDiskSpace := 70000000
	totalRequired := 30000000
	totalDiskLeft := totalDiskSpace - totalUsedSpace
	totalDiskNeededToDelete := totalRequired - totalDiskLeft

	sizeOfDirToDelete := 0
	closestDifference := 999999999
	for j := 0; j < len(directories); j++ {
		if directories[j].Size > totalDiskNeededToDelete {
			difference := directories[j].Size - totalDiskNeededToDelete
			if difference < closestDifference {
				sizeOfDirToDelete = directories[j].Size
				closestDifference = difference
			}
		}
	}

	helpers.PrintResult(fmt.Sprintf("Part 2: %d", sizeOfDirToDelete))

}

func updateDirSize(dir *entities.Directory) {
	if len(dir.Directories) == 0 {
		dir.CalculateSize()
	} else {
		for i := 0; i < len(dir.Directories); i++ {
			updateDirSize(dir.Directories[i])
			dir.CalculateSize()
		}
	}
}

func (d *Day7) addToFileSystem(dir *entities.Directory) {
	d.fileSystem = append(d.fileSystem, dir)
}
