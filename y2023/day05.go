package y2023

import (
	"sort"

	"github.com/MoraGames/adventofcode/utils"
)

type SmallRange struct {
	strt uint64
	lngt uint64
}

type Range struct {
	dest uint64
	strt uint64
	lngt uint64
}

func D5P1() {
	day, part = 5, 1
	sections := utils.SplitDoubleLines(ReadFile())

	// Configure the maps
	var startingSeeds []uint64
	seeds := utils.SplitSpaces(utils.Split(sections[0], ": ")[1])
	for s := 0; s < len(seeds); s++ {
		startingSeeds = append(startingSeeds, uint64(utils.StoI(seeds[s])))
	}

	var soilMap, fertilizerMap, waterMap, lightMap, temperatureMap, humidityMap, locationMap []Range
	for s := 1; s < len(sections); s++ {
		numLines := utils.SplitLines(utils.Split(sections[s], ":\r\n")[1])
		for n := 0; n < len(numLines); n++ {
			numArgs := utils.SplitSpaces(numLines[n])
			switch s {
			case 1: // Soil
				soilMap = append(soilMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 2: // Fertilizer
				fertilizerMap = append(fertilizerMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 3: // Water
				waterMap = append(waterMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 4: // Light
				lightMap = append(lightMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 5: // Temperature
				temperatureMap = append(temperatureMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 6: // Humidity
				humidityMap = append(humidityMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 7: // Location
				locationMap = append(locationMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			}
		}
	}

	// Print the maps
	/*
		fmt.Printf("Starting seeds: %v\n", startingSeeds)
		fmt.Printf("Soil: %v\n", soilMap)
		fmt.Printf("Fertilizer: %v\n", fertilizerMap)
		fmt.Printf("Water: %v\n", waterMap)
		fmt.Printf("Light: %v\n", lightMap)
		fmt.Printf("Temperature: %v\n", temperatureMap)
		fmt.Printf("Humidity: %v\n", humidityMap)
		fmt.Printf("Location: %v\n", locationMap)
	*/

	var allLocations []uint64
	for ss := 0; ss < len(startingSeeds); ss++ {
		value := startingSeeds[ss]

		// find the soil
		for s := 0; s < len(soilMap); s++ {
			if soilMap[s].strt <= value && value < soilMap[s].strt+soilMap[s].lngt {
				value = value - soilMap[s].strt + soilMap[s].dest
				break
			}
		}
		// find the fertilizer
		for f := 0; f < len(fertilizerMap); f++ {
			if fertilizerMap[f].strt <= value && value < fertilizerMap[f].strt+fertilizerMap[f].lngt {
				value = value - fertilizerMap[f].strt + fertilizerMap[f].dest
				break
			}
		}
		// find the water
		for w := 0; w < len(waterMap); w++ {
			if waterMap[w].strt <= value && value < waterMap[w].strt+waterMap[w].lngt {
				value = value - waterMap[w].strt + waterMap[w].dest
				break
			}
		}
		// find the light
		for l := 0; l < len(lightMap); l++ {
			if lightMap[l].strt <= value && value < lightMap[l].strt+lightMap[l].lngt {
				value = value - lightMap[l].strt + lightMap[l].dest
				break
			}
		}
		// find the temperature
		for t := 0; t < len(temperatureMap); t++ {
			if temperatureMap[t].strt <= value && value < temperatureMap[t].strt+temperatureMap[t].lngt {
				value = value - temperatureMap[t].strt + temperatureMap[t].dest
				break
			}
		}
		// find the humidity
		for h := 0; h < len(humidityMap); h++ {
			if humidityMap[h].strt <= value && value < humidityMap[h].strt+humidityMap[h].lngt {
				value = value - humidityMap[h].strt + humidityMap[h].dest
				break
			}
		}
		// find the location
		for l := 0; l < len(locationMap); l++ {
			if locationMap[l].strt <= value && value < locationMap[l].strt+locationMap[l].lngt {
				value = value - locationMap[l].strt + locationMap[l].dest
				break
			}
		}

		allLocations = append(allLocations, value)
	}
	sort.Slice(allLocations, func(i, j int) bool { return allLocations[i] < allLocations[j] })
	// Logger.WithFields(logrus.Fields{
	// 	"locVal": allLocations[0],
	// }).Info("Found the nearest location")

	PrintSolution(allLocations[0])
}

func D5P2() {
	day, part = 5, 2
	sections := utils.SplitDoubleLines(ReadFile())

	// Configure the seeds map
	var seedMap []SmallRange
	seeds := utils.SplitSpaces(utils.Split(sections[0], ": ")[1])
	for s := 0; s < len(seeds); s += 2 {
		seedMap = append(seedMap, SmallRange{uint64(utils.StoI(seeds[s])), uint64(utils.StoI(seeds[s+1]))})
	}

	// Configure the maps
	var soilMap, fertilizerMap, waterMap, lightMap, temperatureMap, humidityMap, locationMap []Range
	for s := 1; s < len(sections); s++ {
		numLines := utils.SplitLines(utils.Split(sections[s], ":\r\n")[1])
		for n := 0; n < len(numLines); n++ {
			numArgs := utils.SplitSpaces(numLines[n])
			switch s {
			case 1: // Soil
				soilMap = append(soilMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 2: // Fertilizer
				fertilizerMap = append(fertilizerMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 3: // Water
				waterMap = append(waterMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 4: // Light
				lightMap = append(lightMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 5: // Temperature
				temperatureMap = append(temperatureMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 6: // Humidity
				humidityMap = append(humidityMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			case 7: // Location
				locationMap = append(locationMap, Range{uint64(utils.StoI(numArgs[0])), uint64(utils.StoI(numArgs[1])), uint64(utils.StoI(numArgs[2]))})
			}
		}
	}

	// Print the maps
	/*
		fmt.Printf("Starting seeds: %v\n", startingSeeds)
		fmt.Printf("Soil: %v\n", soilMap)
		fmt.Printf("Fertilizer: %v\n", fertilizerMap)
		fmt.Printf("Water: %v\n", waterMap)
		fmt.Printf("Light: %v\n", lightMap)
		fmt.Printf("Temperature: %v\n", temperatureMap)
		fmt.Printf("Humidity: %v\n", humidityMap)
		fmt.Printf("Location: %v\n", locationMap)
	*/

	var minLocation uint64
	locFinded := false
	for sm := 0; sm < len(seedMap); sm++ {
		// Logger.Debug(fmt.Sprintf("SeedMap %v/%v\n", sm+1, len(seedMap)))
		for ss := seedMap[sm].strt; ss < seedMap[sm].strt+seedMap[sm].lngt; ss++ {
			// starting seed
			var value uint64 = uint64(ss)

			// find the soil
			for s := 0; s < len(soilMap); s++ {
				if soilMap[s].strt <= value && value < soilMap[s].strt+soilMap[s].lngt {
					value = value - soilMap[s].strt + soilMap[s].dest
					break
				}
			}
			// find the fertilizer
			for f := 0; f < len(fertilizerMap); f++ {
				if fertilizerMap[f].strt <= value && value < fertilizerMap[f].strt+fertilizerMap[f].lngt {
					value = value - fertilizerMap[f].strt + fertilizerMap[f].dest
					break
				}
			}
			// find the water
			for w := 0; w < len(waterMap); w++ {
				if waterMap[w].strt <= value && value < waterMap[w].strt+waterMap[w].lngt {
					value = value - waterMap[w].strt + waterMap[w].dest
					break
				}
			}
			// find the light
			for l := 0; l < len(lightMap); l++ {
				if lightMap[l].strt <= value && value < lightMap[l].strt+lightMap[l].lngt {
					value = value - lightMap[l].strt + lightMap[l].dest
					break
				}
			}
			// find the temperature
			for t := 0; t < len(temperatureMap); t++ {
				if temperatureMap[t].strt <= value && value < temperatureMap[t].strt+temperatureMap[t].lngt {
					value = value - temperatureMap[t].strt + temperatureMap[t].dest
					break
				}
			}
			// find the humidity
			for h := 0; h < len(humidityMap); h++ {
				if humidityMap[h].strt <= value && value < humidityMap[h].strt+humidityMap[h].lngt {
					value = value - humidityMap[h].strt + humidityMap[h].dest
					break
				}
			}
			// find the location
			for l := 0; l < len(locationMap); l++ {
				if locationMap[l].strt <= value && value < locationMap[l].strt+locationMap[l].lngt {
					value = value - locationMap[l].strt + locationMap[l].dest
					break
				}
			}

			if !locFinded || value < minLocation {
				minLocation = value
				locFinded = true
			}
		}
	}
	// Logger.WithFields(logrus.Fields{
	// 	"locVal": minLocation,
	// }).Info("Found the nearest location")
	PrintSolution(minLocation)
}
