package y2023

import (
	"github.com/MoraGames/adventofcode/utils"
)

type StartingPosition struct {
	coord      Coord
	undershape string
}

var StrtPos StartingPosition

type Coord struct {
	x, y int
}
type Tile struct {
	coord      Coord
	shape      string
	isStrtPos  bool
	isMainLoop bool
	distance   int
	visited    bool
}

func (t *Tile) GetConnections() []Coord {
	var connections []Coord
	switch t.shape {
	case ".":
		connections = []Coord{}
	case "|":
		connections = []Coord{{t.coord.x, t.coord.y - 1}, {t.coord.x, t.coord.y + 1}}
	case "-":
		connections = []Coord{{t.coord.x - 1, t.coord.y}, {t.coord.x + 1, t.coord.y}}
	case "F":
		connections = []Coord{{t.coord.x + 1, t.coord.y}, {t.coord.x, t.coord.y + 1}}
	case "7":
		connections = []Coord{{t.coord.x - 1, t.coord.y}, {t.coord.x, t.coord.y + 1}}
	case "L":
		connections = []Coord{{t.coord.x, t.coord.y - 1}, {t.coord.x + 1, t.coord.y}}
	case "J":
		connections = []Coord{{t.coord.x, t.coord.y - 1}, {t.coord.x - 1, t.coord.y}}
	}
	return connections
}

var OnTest bool = false

func D10P1() {
	day, part = 10, 1
	var lines []string
	if OnTest {
		lines = utils.SplitLines(ReadTestFile())
		StrtPos = StartingPosition{
			coord:      Coord{x: -1, y: -1},
			undershape: "F",
		}
	} else {
		lines = utils.SplitLines(ReadFile())
		StrtPos = StartingPosition{
			coord:      Coord{x: -1, y: -1},
			undershape: "7",
		}
	}

	// log.Printf("file readed.")

	var tilesMap map[Coord]Tile = make(map[Coord]Tile)
	for l := 0; l < len(lines); l++ {
		tiles := utils.Split(lines[l], "")
		for t := 0; t < len(tiles); t++ {
			if tiles[t] == "S" {
				tilesMap[Coord{x: t, y: l}] = Tile{
					coord:      Coord{x: t, y: l},
					shape:      StrtPos.undershape,
					isStrtPos:  true,
					isMainLoop: true,
					distance:   0,
					visited:    true,
				}
				StrtPos.coord = Coord{x: t, y: l}
			} else {
				tilesMap[Coord{x: t, y: l}] = Tile{
					coord:      Coord{x: t, y: l},
					shape:      tiles[t],
					isStrtPos:  false,
					isMainLoop: false,
					distance:   -1,
					visited:    false,
				}
			}
		}
	}

	// log.Printf("TilesMap setted up. Length: %v\n", len(tilesMap))
	// log.Printf("Starting position: %v\n", StrtPos)
	// log.Printf("tilesMap[sp.coord]: %v\n\n", tilesMap[StrtPos.coord])

	maxDistance := 0
	nxtCoord := []Coord{StrtPos.coord}
	for {
		curTile, exist := tilesMap[nxtCoord[0]]
		if !exist {
			// log.Panicf("Tile in TilesMap position not found: %v", StrtPos.coord)
			panic([]any{"Tile in TilesMap position not found: ", StrtPos.coord})
		}
		curTile.visited = true
		curTile.isMainLoop = true

		connections := curTile.GetConnections()
		for c := 0; c < len(connections); c++ {
			nxtTile := tilesMap[connections[c]]
			if !nxtTile.visited {
				nxtTile.distance = curTile.distance + 1
				tilesMap[connections[c]] = nxtTile

				if nxtTile.distance > maxDistance {
					maxDistance = nxtTile.distance
				}

				nxtCoord = append(nxtCoord, connections[c])
			}
		}

		tilesMap[nxtCoord[0]] = curTile

		nxtCoord = nxtCoord[1:]
		if len(nxtCoord) == 0 {
			break
		}

		// log.Printf("New nxtCoord length: %v\n", len(nxtCoord))
		// log.Printf("New nxtCoord: %v\n", nxtCoord)

		// time.Sleep(2 * time.Second)
	}

	// log.Printf("Max distance: %v\n", maxDistance)
	PrintSolution(maxDistance)
}

func D10P2A() {
	day, part = 10, 2
	var lines []string
	if OnTest {
		lines = utils.SplitLines(ReadTestFile())
		StrtPos = StartingPosition{
			coord:      Coord{x: -1, y: -1},
			undershape: "F",
		}
	} else {
		lines = utils.SplitLines(ReadFile())
		StrtPos = StartingPosition{
			coord:      Coord{x: -1, y: -1},
			undershape: "7",
		}
	}

	// log.Printf("file readed.")

	var tilesMap map[Coord]Tile = make(map[Coord]Tile)
	for l := 0; l < len(lines); l++ {
		tiles := utils.Split(lines[l], "")
		for t := 0; t < len(tiles); t++ {
			if tiles[t] == "S" {
				tilesMap[Coord{x: t, y: l}] = Tile{
					coord:      Coord{x: t, y: l},
					shape:      StrtPos.undershape,
					isStrtPos:  true,
					isMainLoop: true,
					distance:   0,
					visited:    true,
				}
				StrtPos.coord = Coord{x: t, y: l}
			} else {
				tilesMap[Coord{x: t, y: l}] = Tile{
					coord:      Coord{x: t, y: l},
					shape:      tiles[t],
					isStrtPos:  false,
					isMainLoop: false,
					distance:   -1,
					visited:    false,
				}
			}
		}
	}

	// log.Printf("TilesMap setted up. Length: %v\n", len(tilesMap))
	// log.Printf("Starting position: %v\n", StrtPos)
	// log.Printf("tilesMap[sp.coord]: %v\n\n", tilesMap[StrtPos.coord])

	nxtCoord := []Coord{StrtPos.coord}
	for {
		curTile, exist := tilesMap[nxtCoord[0]]
		if !exist {
			// log.Panicf("Tile in TilesMap position not found: %v", StrtPos.coord)
			panic([]any{"Tile in TilesMap position not found: ", StrtPos.coord})
		}
		curTile.visited = true
		curTile.isMainLoop = true

		connections := curTile.GetConnections()
		for c := 0; c < len(connections); c++ {
			nxtTile := tilesMap[connections[c]]
			if !nxtTile.visited {
				nxtTile.distance = curTile.distance + 1
				tilesMap[connections[c]] = nxtTile

				nxtCoord = append(nxtCoord, connections[c])
			}
		}

		tilesMap[nxtCoord[0]] = curTile

		nxtCoord = nxtCoord[1:]
		if len(nxtCoord) == 0 {
			break
		}

		// log.Printf("New nxtCoord length: %v\n", len(nxtCoord))
		// log.Printf("New nxtCoord: %v\n", nxtCoord)

		// time.Sleep(2 * time.Second)
	}

	var newFile string
	for l := 0; l < len(lines); l++ {
		for t := 0; t < len(lines[l]); t++ {
			tm := tilesMap[Coord{x: t, y: l}]
			if tm.isMainLoop {
				newFile += ConvertToUnicode(tm.shape)
			} else {
				newFile += " "
			}
		}
		newFile += "\r\n"
	}

	// os.WriteFile("d10out.txt", []byte(newFile), 0644)
	WriteOutput(newFile, "day10-output.txt")

	// log.Printf("Max distance: %v\n", maxDistance)
	// PrintSolution(maxDistance)
}

func ConvertToUnicode(shape string) string {
	switch shape {
	case "|":
		return "│"
	case "-":
		return "─"
	case "F":
		return "┌"
	case "7":
		return "┐"
	case "L":
		return "└"
	case "J":
		return "┘"
	default:
		return "?"
	}
}

func D10P2B() {
	day, part = 10, 3
	var counter int
	lines := utils.SplitLines(utils.Read("y2023/outputs/day10-output.txt"))
	for l := 0; l < len(lines); l++ {
		tiles := utils.Split(lines[l], "")
		for t := 0; t < len(tiles); t++ {
			if tiles[t] == "o" {
				counter++
			}
		}
	}
	// log.Printf("Counter: %v\n", counter)
	PrintSolution(counter)
}
