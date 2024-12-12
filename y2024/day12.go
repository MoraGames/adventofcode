package y2024

import "github.com/MoraGames/adventofcode/utils"

type region struct {
	area      int
	perimeter int
	coords    map[utils.MatrixCoord]struct{}
}

func D12P1() {
	day, part = 12, 1
	gardens := utils.SplitMatrixLinesCharacters(ReadFile())
	regions := findRegions(gardens)

	var price int
	for _, region := range regions {
		price += region.area * region.perimeter
	}

	PrintSolution(price)
}

func D12P2() {
	day, part = 12, 2
	gardens := utils.SplitMatrixLinesCharacters(ReadFile())
	regions := findRegions(gardens)

	var price int
	for _, region := range regions {
		price += region.area * findSides(region.coords)
	}

	PrintSolution(price)
}

func findRegions(gardens [][]string) []region {
	var regions []region
	checked := utils.NewCheck[utils.MatrixCoord]()
	for r := 0; r < len(gardens); r++ {
		for c := 0; c < len(gardens[r]); c++ {
			currCoord := utils.MatrixCoord{Row: r, Col: c}
			if exist := utils.CheckExists(checked, currCoord); !exist {
				utils.CheckInsert(checked, currCoord)
				region := region{
					area:      0,
					perimeter: 0,
					coords:    utils.NewCheck[utils.MatrixCoord](),
				}
				defineRegion(gardens, currCoord, checked, &region)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func defineRegion(gardens [][]string, currCoord utils.MatrixCoord, checked map[utils.MatrixCoord]struct{}, region *region) {
	region.area++
	region.perimeter += 4
	utils.CheckInsert(region.coords, currCoord)

	neighbor := []utils.MatrixCoord{
		{Row: currCoord.Row - 1, Col: currCoord.Col},
		{Row: currCoord.Row, Col: currCoord.Col + 1},
		{Row: currCoord.Row + 1, Col: currCoord.Col},
		{Row: currCoord.Row, Col: currCoord.Col - 1},
	}
	for n := 0; n < len(neighbor); n++ {
		if utils.MatrixIsInbound(gardens, neighbor[n]) && gardens[neighbor[n].Row][neighbor[n].Col] == gardens[currCoord.Row][currCoord.Col] {
			if exist := utils.CheckExists(checked, neighbor[n]); !exist {
				utils.CheckInsert(checked, neighbor[n])
				defineRegion(gardens, neighbor[n], checked, region)
			}
			region.perimeter--
		}
	}
}

func findSides(coords map[utils.MatrixCoord]struct{}) int {
	var numSides int
	for coord := range coords {
		n := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row - 1, Col: coord.Col})
		ne := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row - 1, Col: coord.Col + 1})
		e := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row, Col: coord.Col + 1})
		se := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row + 1, Col: coord.Col + 1})
		s := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row + 1, Col: coord.Col})
		sw := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row + 1, Col: coord.Col - 1})
		w := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row, Col: coord.Col - 1})
		nw := utils.CheckExists(coords, utils.MatrixCoord{Row: coord.Row - 1, Col: coord.Col - 1})

		if !n && !e {
			numSides++
		}
		if !e && !s {
			numSides++
		}
		if !s && !w {
			numSides++
		}
		if !w && !n {
			numSides++
		}

		if !ne && n && e {
			numSides++
		}
		if !se && e && s {
			numSides++
		}
		if !sw && s && w {
			numSides++
		}
		if !nw && w && n {
			numSides++
		}
	}
	return numSides
}
