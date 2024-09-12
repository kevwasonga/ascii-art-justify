package asciiArt

import "os"

func BannerFile() string {
	// Ensure the number of arguments is correct
	if len(os.Args) < 3 || len(os.Args) > 4 {
		return ""
	}

	// If there are 4 arguments (i.e., alignment flag, string, and banner type)
	if len(os.Args) == 4 {
		switch os.Args[3] {
		case "standard":
			return "banners/standard.txt"
		case "shadow":
			return "banners/shadow.txt"
		case "thinkertoy":
			return "banners/thinkertoy.txt"
		default:
			return "invalid bannerfile name"
		}
	}

	// If no banner type is provided (i.e., default to standard)
	if len(os.Args) == 3 {
		return "banners/standard.txt"
	}

	// Return empty string in case of unexpected argument format
	return ""
}
