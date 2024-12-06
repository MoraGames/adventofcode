package y2023

import (
	"github.com/MoraGames/adventofcode/utils"
)

type Galaxy struct {
	coord Coord
	id    int
}

func D11P1() {
	day, part = 11, 1
	var universe [][]string
	galaxyMap := make(map[int]Galaxy)

	lines := utils.SplitLines(ReadTestFile())
	for l := 0; l < len(lines); l++ {
		ceils := utils.Split(lines[l], "")
		universe = append(universe, ceils)
		var finded bool
		for c := 0; c < len(ceils); c++ {
			if ceils[c] == "#" {
				finded = true
			}
		}
		if !finded { // Expand the universe horizontally
			// fmt.Println("Expanding universe horizontally at line:", l+1)
			universe = append(universe, ceils)
		}
	}
	for c := 0; c < len(universe[0]); c++ { // Expand the universe vertically
		var finded bool
		for l := 0; l < len(universe); l++ {
			if universe[l][c] == "#" {
				finded = true
			}
		}
		if !finded {
			// fmt.Println("Expanding universe vertically at column:", c+1)
			for l := 0; l < len(universe); l++ {
				tempLine := universe[l][:c]
				tempLine = append(tempLine, universe[l][c])
				tempLine = append(tempLine, universe[l][c:]...)
				universe[l] = tempLine
			}
			c++
		}
	}

	// Print the universe
	/*
		for i := 0; i < 3; i++ {
			fmt.Print("      ")
			for c := 0; c < len(universe[0]); c++ {
				colNum := ItoS(c + 1)
				if len(colNum) == 1 {
					colNum = "  " + colNum
				} else if len(colNum) == 2 {
					colNum = " " + colNum
				}
				fmt.Printf("%s", string(colNum[i]))
			}
			fmt.Println()
		}
		for l := 0; l < len(universe); l++ {
			lineNum := ItoS(l + 1)
			if len(lineNum) == 1 {
				lineNum = "  " + lineNum
			} else if len(lineNum) == 2 {
				lineNum = " " + lineNum
			}
			fmt.Printf("%s)  ", lineNum)
			for c := 0; c < len(universe[0]); c++ {
				fmt.Print(universe[l][c])
			}
			fmt.Println()
		}
	*/

	// Map all the galaxies
	var lastID int = -1
	for l := 0; l < len(universe); l++ {
		for c := 0; c < len(universe[l]); c++ {
			if universe[l][c] == "#" {
				galaxyMap[lastID+1] = Galaxy{Coord{c, l}, lastID + 1}
				lastID++
			}
		}
	}

	// Find the shortest path between all pair of galaxies
	var sum int
	for i := 0; i < len(galaxyMap); i++ {
		for j := i + 1; j < len(galaxyMap); j++ {
			pathLength := findShortestPath(galaxyMap[i], galaxyMap[j], universe)
			// fmt.Printf("  x: %d) = %d | now sum %v -> %v\n", j, pathLength, sum, sum+pathLength)
			sum += pathLength
		}
		// fmt.Printf("SP(%d,x) = %d\n", i, sum)
	}
	// fmt.Printf("Sum of all shortest paths: %d\n", sum)
	PrintSolution(sum)
}

func D11P2() {
	day, part = 11, 2
	var universe [][]string
	var lastGalaxyID int = -1
	galaxyMap, expansionMap := make(map[int]Galaxy), make(map[string]int)

	lines := utils.SplitLines(ReadFile())
	for l := 0; l < len(lines); l++ {
		ceils := utils.Split(lines[l], "")
		universe = append(universe, ceils)
		var finded bool
		for c := 0; c < len(ceils); c++ {
			if ceils[c] == "#" {
				galaxyMap[lastGalaxyID+1] = Galaxy{Coord{c, l}, lastGalaxyID + 1}
				lastGalaxyID++
				finded = true
			}
		}
		if !finded { // Expand the universe horizontally
			// fmt.Println("Expanding universe horizontally at line:", l)
			expansionMap["L"+utils.ItoS(l)] = 1000000
		}
	}
	for c := 0; c < len(universe[0]); c++ { // Expand the universe vertically
		var finded bool
		for l := 0; l < len(universe); l++ {
			if universe[l][c] == "#" {
				finded = true
			}
		}
		if !finded {
			// fmt.Println("Expanding universe vertically at column:", c)
			expansionMap["C"+utils.ItoS(c)] = 1000000
		}
	}

	// Find the shortest path between all pair of galaxies
	var sum uint64
	for i := 0; i < len(galaxyMap); i++ {
		for j := i + 1; j < len(galaxyMap); j++ {
			pathLength := findShortestPath2(galaxyMap[i], galaxyMap[j], universe, expansionMap)
			// fmt.Printf("SP(%d->%d) = %d | now sum %v -> %v\n", i, j, pathLength, sum, sum+uint64(pathLength))
			sum += uint64(pathLength)
		}
		// fmt.Printf("SP(%d,x) = %d\n", i, sum)
	}
	// fmt.Printf("Sum of all shortest paths: %d\n", sum)
	PrintSolution(sum)
}

func findShortestPath(galaxy1 Galaxy, galaxy2 Galaxy, universe [][]string) int {
	// Create a queue to store the coordinates and distance
	queue := []struct {
		coord    Coord
		distance int
	}{{
		coord:    galaxy1.coord,
		distance: 0,
	}}

	// Create a visited map to keep track of visited coordinates
	visited := make(map[Coord]bool)
	visited[galaxy1.coord] = true

	// Define the possible moves: up, down, left, right
	moves := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// Perform breadth-first search
	for len(queue) > 0 {
		// Dequeue the front element
		current := queue[0]
		queue = queue[1:]

		// Check if the current coordinate is the target galaxy
		if current.coord == galaxy2.coord {
			return current.distance
		}

		// Explore the neighboring coordinates
		for _, move := range moves {
			next := Coord{current.coord.x + move.x, current.coord.y + move.y}

			// Check if the next coordinate is within the universe boundaries
			if next.x >= 0 && next.x < len(universe[0]) && next.y >= 0 && next.y < len(universe) {
				// Check if the next coordinate is not visited
				if !visited[next] {
					// Enqueue the next coordinate with an increased distance
					queue = append(queue, struct {
						coord    Coord
						distance int
					}{coord: next, distance: current.distance + 1})

					// Mark the next coordinate as visited
					visited[next] = true
				}
			}
		}
	}

	// If no path is found, return -1
	return -1
}

func findShortestPath2(galaxy1 Galaxy, galaxy2 Galaxy, universe [][]string, expansions map[string]int) int {
	// Create a queue to store the coordinates and distance
	queue := []struct {
		coord    Coord
		distance int
	}{{
		coord:    galaxy1.coord,
		distance: 0,
	}}

	// Create a visited map to keep track of visited coordinates
	visited := make(map[Coord]bool)
	visited[galaxy1.coord] = true

	// Define the possible moves: up, down, left, right
	moves := []struct {
		coord     Coord
		direction string
	}{{Coord{0, 1}, "down"}, {Coord{0, -1}, "up"}, {Coord{1, 0}, "left"}, {Coord{-1, 0}, "right"}}

	// Perform breadth-first search
	for len(queue) > 0 {
		// Dequeue the front element
		current := queue[0]
		queue = queue[1:]

		// Check if the current coordinate is the target galaxy
		if current.coord == galaxy2.coord {
			return current.distance
		}

		// Explore the neighboring coordinates
		for _, move := range moves {
			next := Coord{current.coord.x + move.coord.x, current.coord.y + move.coord.y}

			// Check if the next coordinate is within the universe boundaries
			if next.x >= 0 && next.x < len(universe[0]) && next.y >= 0 && next.y < len(universe) {
				// Check if the next coordinate is not visited
				if !visited[next] {
					// fmt.Printf("Current: %v | Next: %v | Move: %v\n", current.coord, next, move)

					var distanceVariation int = 1
					if move.direction == "left" || move.direction == "right" {
						// Check if the next coordinate is a vertical expansion
						if value, exist := expansions["C"+utils.ItoS(next.x)]; exist {
							// fmt.Println("Finded a vertical expansion at", next.x)
							distanceVariation = value
						}
					} else if move.direction == "up" || move.direction == "down" {
						// Check if the next coordinate is an horizontal expansion
						if value, exist := expansions["L"+utils.ItoS(next.y)]; exist {
							// fmt.Println("Finded an horizontal expansion at", next.y)
							distanceVariation = value
						}
					}

					// Enqueue the next coordinate with an increased distance
					queue = append(queue, struct {
						coord    Coord
						distance int
					}{coord: next, distance: current.distance + distanceVariation})

					// Mark the next coordinate as visited
					visited[next] = true
				}
			}
		}
	}

	// If no path is found, return -1
	return -1
}

// 9'226'330 too low
// 9918828 correct
