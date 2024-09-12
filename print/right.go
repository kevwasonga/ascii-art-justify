package print

import (
	"fmt"
	"strings"
)

func PrintRight(line string, bannerMap map[int][]string) {
	// Dynamically calculate the console width
	consoleWidth := getConsoleWidth()

	output := make([]string, 8)

	for _, char := range line {
		banner, exists := bannerMap[int(char)]
		if !exists {
			fmt.Printf("Character %c not found in banner map\n", char)
			continue
		}

		for i := 0; i < 8; i++ {
			output[i] += banner[i]
		}
	}

	// Calculate how many spaces to pad the left side for right alignment
	for i := range output {
		padding := consoleWidth - len(output[i])
		if padding > 0 {
			output[i] = strings.Repeat(" ", padding) + output[i]
		}
	}

	// Print the right-aligned output
	for _, outLine := range output {
		fmt.Println(outLine)
	}
}
