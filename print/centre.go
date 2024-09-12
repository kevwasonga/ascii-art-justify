package print

import (
	"fmt"
	"strings"
)

// PrintCenter centers the given text within the console's width.
func PrintCenter(line string, bannerMap map[int][]string) {
	// Get console width
	consoleWidth := getConsoleWidth()

	// Create an output slice of strings for the 8 lines of the banner
	output := make([]string, 8)

	// Loop through each character in the line and append its corresponding banner to the output
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

	// Center-align each line of the banner by calculating the left padding
	for i := range output {
		padding := (consoleWidth - len(output[i])) / 2
		if padding > 0 {
			output[i] = strings.Repeat(" ", padding) + output[i]
		}
	}

	// Print the center-aligned output
	for _, outLine := range output {
		fmt.Println(outLine)
	}
}
