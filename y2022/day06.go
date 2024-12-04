package y2022

import (
	"fmt"
)

func D6P1() {
	file := ReadFile(6)
	for i := 0; i < len(file); i++ {
		section := file[i : i+4]
		//fmt.Printf("Section: %v\n", section)

		var exit bool
		for s1 := 0; s1 < 4 && !exit; s1++ {
			for s2 := s1 + 1; s2 < 4 && !exit; s2++ {
				if s1 != s2 && section[s1] == section[s2] {
					exit = true
				}
			}
		}
		if !exit {
			fmt.Printf("Section[%d:%d] = %v\n", i, i+4, section)
			break
		}
	}
}

func D6P2() {
	file := ReadFile(6)
	for i := 0; i < len(file); i++ {
		section := file[i : i+14]
		//fmt.Printf("Section: %v\n", section)

		var exit bool
		for s1 := 0; s1 < 14 && !exit; s1++ {
			for s2 := s1 + 1; s2 < 14 && !exit; s2++ {
				if s1 != s2 && section[s1] == section[s2] {
					exit = true
				}
			}
		}
		if !exit {
			fmt.Printf("Section[%d:%d] = %v\n", i, i+14, section)
			break
		}
	}
}
