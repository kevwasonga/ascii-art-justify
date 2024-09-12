package print

import (
	"fmt"
	"log"
	"strings"
)

// PrintJustify takes a string and a banner map, and prints the justified text by distributing spaces equally among characters.
func PrintJustify(input string, bannerMap map[int][]string) {
	// Get the console width
	cWidth := getConsoleWidth()

	// Remove leading and trailing spaces in the input and split into characters
	trimmedInput := strings.ReplaceAll(input, " ", "") // Ignore internal spaces
	runes := []rune(trimmedInput)

	if len(runes) == 0 {
		fmt.Println("Input is empty or only contains spaces.")
		return
	}

	// Get the total width of all characters (ASCII art)
	totalWidth := getTotalWidth(runes, bannerMap)
	if totalWidth > cWidth {
		log.Fatalln("Ascii-art doesn't fit into console window.")
	}

	// Calculate available space for distribution
	availableSpace := cWidth - totalWidth

	// Calculate the base space between each character and the remainder spaces
	spacePerChar := availableSpace / (len(runes) - 1) // Base number of spaces between each character
	extraSpaces := availableSpace % (len(runes) - 1)  // Extra spaces to distribute

	// Create the banner output line by line
	output := make([]string, 8)
	for i := 0; i < 8; i++ {
		var tmp string

		// Add each character's banner to the output and distribute spaces
		for j, r := range runes {
			if banner, exists := bannerMap[int(r)]; exists {
				tmp += banner[i]
			} else {
				fmt.Printf("Character %c not found in banner map\n", r)
				return
			}

			// Add spaces between characters (but not after the last character)
			if j < len(runes)-1 {
				tmp += strings.Repeat(" ", spacePerChar)

				// Distribute extra spaces if available
				if j < extraSpaces {
					tmp += " "
				}
			}
		}

		output[i] = tmp
	}

	// Print the justified output
	for _, line := range output {
		fmt.Println(line)
	}
}

// getTotalWidth calculates the total width of all characters in the input
func getTotalWidth(runes []rune, bannerMap map[int][]string) int {
	totalWidth := 0
	for _, r := range runes {
		if banner, exists := bannerMap[int(r)]; exists {
			totalWidth += len(banner[0]) // Use the first line to measure width
		}
	}
	return totalWidth
}
