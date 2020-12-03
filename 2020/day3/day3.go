package day3

import (
	"fmt"
	"strings"
)

func Run() {
	input := ".##.....#....#....#..#.#...#.##\n.###........#.##....#......#..#\n#..#..#.....#...#....#.#.......\n.........#.................#...\n..#.......#.#.......#.......#.#\n.####........#.#..##.........#.\n........#.........#.........#..\n#..##...##....#.....##......#..\n.........#..............#......\n#.........#...##.........#.#...\n..............#........##.....#\n##....#...........#....#.#...#.\n.....#..#.....#...#.#..........\n#.......#...#..##........##..#.\n.#........#.......#............\n.......##.....#.#.#..#.#.......\n..#......#..#....##......#.#...\n.....##....#..#.....#..#.......\n.............#.......#.#....#..\n.................#.#......#....\n.#..#....#..........#.....#.##.\n#.#.#.#.....###.......#.....#..\n#...#..........#..#............\n...#...##.......#.##..#........\n..#...#.#.##...##.........#.#.#\n.....#...#.#....#.#.....#......\n...#...#.#..#...#.....#........\n...........###.#.......#.#.....\n..#..#.#........#.#.......#.#.#\n.#.......#...........#.........\n.#..#...##....#.......###..##..\n#....#.....#....##..#.........#\n#..#.......#...#......#.#....#.\n......##..#..#....#.#........##\n#.....#.........#......#..##..#\n.#..#.......#....#............#\n....#..........#.#...##....#.##\n..#...#.#...#.###.#..#......#..\n#.#...#..............#.......#.\n..##.......#......#....###.....\n......#.......#.#.##.#..##.#...\n.#......#......#.#....#..#.#..#\n....#....#..#...#.....#.#..#...\n.#.....#.#.#..#........#.#.###.\n#..#..#.#......#..#..#....#.#..\n..#.###....#....##...#.........\n...........#..#...........#....\n.................#..........#.#\n.#.#....#..#........#..#.......\n...........##..#...............\n...#.##.........#.........#.#.#\n........#..#....#.#............\n...##...##..................#.#\n...#..##..#...#......#.....#..#\n.##.#..#..#......#......#.....#\n....#.....#....#......##.#.....\n.....#.##....#...#.............\n......#...#.....##....#...#..##\n...#............#.###....##....\n............#.#.#...#.#........\n#.....#..#.#..##...........#.##\n.....#.#.#.#..##......##.#..#..\n.#.##..#.........#......#.....#\n.#.#.#.#.#..#..#........###....\n......##..........#.#.....##..#\n..#...#..#.....#.#.....#.......\n............#....#.............\n........#...#..#....#.#..###...\n#........##....##..............\n.........#.#.#..#..#...#.#.....\n....#............#....####...#.\n##.#.#......#.....#...#....###.\n...#..#..#..#.......#..#.#.#..#\n..#..#....#...#.##..#.........#\n#..#.#....#....#...#..##..#.#..\n...#....#.............#..#.#..#\n..#......#.##...#..............\n#....##.#.#...##......#.....##.\n.#...##...#...####.....##......\n...........#.###....#...#...#..\n##..#..##..#..#.#.#..###.......\n.#...##......#........##..#....\n.#...#...#.....#............#..\n.#.#.#...#.#..#.......#......#.\n.................#..#.#......#.\n#..#####......##.#....#...#....\n........#......#.....##......#.\n....#.#...#...#..#.......#####.\n.....##......#...#.......#....#\n.#....#...#..#...#.#...#..#...#\n....##.........#.#...####.#....\n...##..........#.#.......##....\n.........#......#.....#....#...\n#....##..#......#.....##....#..\n...#.#.............#...#.....#.\n...........#...#.#....#..#.#...\n.......#.#.#.....#..#........#.\n..##.....#..#.....##..#........\n...#.#..........#...#....#.#...\n..#....#......#...#.#...##..#.#\n.#...#..#..#..#.......#........\n.................#.#...........\n...............#......##.....#.\n..##.....###..#....#.........##\n....#.#........#.####.#...#....\n.....#.....##..####..##.......#\n.....####.#...#......#.........\n........#..#......#.....##.....\n...###..#.#..###.......#.......\n...#...##..#..#..#..#.##.......\n..#......##..#.....##..##......\n#.......#.#..#.................\n#.......#...#..###....#.#.#.#..\n#...#.#....##.##.#...........#.\n.#.........#..###..#.........#.\n#...#......#...#.#.........#...\n.#.##..............#.#....#...#\n........#.....#..#.#.#......#..\n............####.#......#......\n......#.#.#...#.#...#.#.....#..\n....#....#...#.....#.......#.#.\n..#....#....###..#....#.....##.\n.................#.....#.#....#\n.#.............#......#.##..#..\n#.....##.......#..#.....#....#.\n#.#......####...##.....#....##.\n.....#.#....#..................\n....#....#..#.#...........##...\n...#.............##......#..#..\n......#..........#...#...##.###\n...##......##.....#......#....#\n........#.#.#...#...#..##......\n......................###....##\n.#.....#..#..#.##.#.#.#....#.##\n.#..............#.....#......#.\n.#...#.##....#.....#.#.#..#..#.\n##...##.......#.....#..###.....\n...#..#.#....#........#........\n....#..............##...#......\n...........#..#.....##.#.#.#...\n#.#.....##..##.#.#........#....\n.........#....#.....#..##.#...#\n...#...#..#..#.####...#.......#\n.....##.....##.....#......#....\n#.##...#....#............#..##.\n.#.##..#...#....#.#......#.....\n..###................#.........\n.#..#..#................#......\n....#..#........#..#....#......\n.#..........###......#...###...\n...........##...#.#..#.........\n...#....#..........#.....#..##.\n..#..#.............#......###..\n#....#....##.....#....#.##.....\n......#.......#..#..........##.\n#..##.#...#.#.........#....#.#.\n...#...#..........#...#..#....#\n...###..#.#......#.##.#####...#\n..#.....#.#..............#..##.\n#..###......#.#..#........#....\n.#.......#.......#.....#.##....\n.#...##..#.......##.....#....##\n..........#.#..#.....#.........\n.......####...#...#.....##.....\n......#.......#.......#..#.#...\n...##....##.#.......#.##......#\n.#...#............#......##....\n#..#..#...#.#........#.........\n.......#.......#.....##.#......\n.#....##...#....#.........#...#\n#.#....#.....##...........#..#.\n.....#......#....#......#.#...#\n.#............#...#.#....#....#\n........##..#..##..##.##....#..\n........................#.#....\n#....#...#.....................\n##.#.............#.....#...##.#\n....##....###.......#..........\n..#.#..#.#...####.....#.....#..\n#.........#.......#......#..##.\n.#.#.............#..#...#...#..\n#..#....#....#..##.........#...\n#.#.....#.##.#...#.##..#.#..##.\n......#......#.###....#..###...\n.##...#.......#.........#.#...#\n..........#...#....#..#....#...\n.....#...#.....#....##....#.#..\n#....#...........#.#...#.......\n.###..#........##..........#...\n....###.##..#...#.#..##......##\n.#...#...........#...........#.\n#......#....#.##.........##..#.\n.#.......#........#......#.#.#.\n.......#..##.........#......#..\n.#..#.....##....##....#.....#..\n#.#.#.....#...#......#.........\n..............#.#.........#.#..\n....#...#.............#.#......\n..##.#............#.#.##....#..\n.....####..........#.#....##..#\n......#.#.........#.......###..\n#....##.#...#.#...........#...#\n.....#...#......#....###...#..#\n#....#..............#...#......\n...#..###...#..........#....#..\n#......#..#.#.#......#..#...#..\n................##......#..#...\n....#..#..#........##..#...#...\n...##.......#.##.#.....##...#.#\n.......#.##.#..#.....#...#.....\n......#........#..#......##.##.\n....................#.....#.#..\n.##....#...#...##...#.........#\n..#...#..#.##..#.#.#......#....\n#....###.#..#..#...#..#...##...\n#.......#.....#.#.......###.#.#\n.#.##...##..#......#....#...#..\n#.....#.......##..#....#.......\n...###...#............#....#..#\n.#....#.#...#..#..#.##.#.#.#...\n#......#.#..#.#.#......#.......\n..#..#....###.#........#..#.#..\n.......#......##.........#.....\n...#...###..#..#.##.#..##......\n.#.......##.......#..#..#.###.#\n.###.#..#.###...........#......\n...#................#.#...##..#\n....#.###....#.......#........#\n.##...#...#..#.....#...#.......\n.#...#..#...........#.#......##\n...##..#.#.#..#.#.#.......#....\n.#.#..#..#.#...........#.......\n..#....#.#.#.#.#..............#\n..##..............##....#.#..#.\n..#....#...##.....###.....#.#.#\n#....#......#..........##......\n.##.#.#......#...##..###..#....\n.#...........#.##.......##..##.\n###.....##...#.##..#...........\n...#.....#...........#..#.....#\n#.........#....#.......#.......\n.#.#...#.###....#..#...........\n.....#.......#.....#.##.#.#.#..\n..##.#.........##.........#..#.\n.......#....#......#.........##\n...##.....#..#.......#..#.#....\n..#...###.......#..#....###....\n.......#...###......#.#.....#.#\n#....#...#.#....#.#..........#.\n........#..#.....#.#.#.........\n......##.......###.......#...#.\n.........#..#..#.......#.......\n#.......#...#.....#.#..#....#..\n.##....#..###.............#....\n#.#...#.......#.....##.#.#....#\n....#....##.#........##........\n...##...#.#.............#...##.\n##....#.....#..#..#......#.....\n#...#.#........##....##......##\n..#...........#..#......###....\n..##..#.....#......#....##.....\n....###.#...#......##......#...\n....#....###...........###.#..#\n..#....#...#.##....#...#.......\n....##...........#............#\n..#.#......#......#.##.#...#..#\n#.###.............#.#.##.#.....\n#....##....#..#.#.#...........#\n...#...................#.......\n.#...#......#.......#.#....#..#\n....#...#..#..#..#.#.....#....#\n..#....#............#..###..##.\n...##...#...........#..#..#.#..\n..#..#..#.........#.........#.#\n...#.#.....#.#..##.........#...\n....#..........................\n....#.....#.#...#.###.........#\n....#.#.......#..#.#.#...#...#.\n.....#...#..#.....##....#.#.#..\n#....###......#..#..........#..\n.#.....#......##.......#...#.##\n...#..#.....#.#.....#.......##.\n............#..#....#...#..#.#.\n..........#.#..#..##...........\n.......#.......#..##...##.....#\n....#...##.#..#...#.#.......#..\n....#.#........#...####...#....\n#.#.............#.............#\n.#.#......#....#..#..#.....##.#\n#..#...........#........#.....#\n#....#....#.#..#.#....#.#...##.\n....##...##...#...#...........#\n...#.#..#....#..#..#..#........\n...#..##..#........#..........#\n#......#.##..##.......#..#.....\n..#...#......#...##.#..........\n.###.#..#..#........####...#...\n#..............#.#.#........#..\n..##....#.......#....##...#..##\n.##...#..#.#.....#..#.......##.\n..#.........##.......#....#..#.\n.#..#...#..##.#..#.....#.......\n.#....#.........#..#...#...##..\n..###..######..#.##.#....#.....\n....#..#.....#.............#..#\n...#....#.......#..#.#.......##\n.....#......#.......#..##...#..\n.##..#....##..##......#...#..#.\n......#......#...#...###.......\n....#.....#.###..##.....#.#.##.\n.......#....#...#..#..#...#.#..\n...####.#...#...#.#...##....#..\n......#.#....#....#.#....##....\n#..##...........####....##.#...\n...#...##.#.......#.#..........\n..#......#..#..#...#......#....\n..###..#.....#..#.#.......#...#\n#........#...##..#...#....#....\n...#.#...#.....#........#...#..\n...#....#.###...#..#...#..##.#.\n.....#..#..#...#...#..#........\n..#......##...............#.#.#\n.#...###.#....##..........#.#..\n"

	slopeMap := generateSlopeMap(input)
	part1(slopeMap)
}

func part1(slopeMap [][]string) {
	x := 0
	y := 0
	treeCount := 0
	for true {
		if y == len(slopeMap)-1 {
			break
		}
		x = (x + 3) % len(slopeMap[0])
		y += 1
		if slopeMap[y][x] == "#" {
			treeCount++
		}
	}

	fmt.Printf("[*] Trees Encountered: %d\n", treeCount)
}

func generateSlopeMap(input string) [][]string {
	inputSplit := strings.Split(input, "\n")
	slopeMap := make([][]string, len(inputSplit)-1)

	for i := 0; i < len(inputSplit)-1; i++ {
		slopeSplit := strings.Split(strings.TrimSpace(inputSplit[i]), "")
		slopeMap[i] = make([]string, len(slopeSplit))
		slopeMap[i] = slopeSplit
	}

	return slopeMap
}