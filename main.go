package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/asciiArt"
	"ascii/print"
)

func main() {
	// Check if arguments are provided correctly
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
		return
	}

	// Validate the alignment flag format: --align=<type>
	option := os.Args[1]
	if !strings.HasPrefix(option, "--align=") {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
		return
	}

	// Extract the alignment type
	alignType := strings.TrimPrefix(option, "--align=")
	if alignType != "left" && alignType != "right" && alignType != "center" && alignType != "justify" {
		fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
		return
	}

	// The second argument is the string to print
	text := os.Args[2]

	// The third argument is the banner type (optional)

	// Load the appropriate banner map based on the banner type
	fileName := asciiArt.BannerFile()
	bannerMap, err := asciiArt.LoadBannerMap(fileName)
	if err != nil {
		fmt.Println("Error loading banner map:", err)
		return
	}

	// Process the text, replacing special characters like "\n" and "\t"
	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\\t", "    ")
	lines := strings.Split(text, "\n")

	// Use switch case to determine which alignment to apply
	for _, line := range lines {
		switch alignType {
		case "left":
			print.PrintLeft(line, bannerMap)
		case "right":
			print.PrintRight(line, bannerMap)
		case "center":
			print.PrintCenter(line, bannerMap)
		case "justify":
			print.PrintJustify(line, bannerMap)
		default:
			fmt.Printf("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard\n")
			return
		}
	}
}
