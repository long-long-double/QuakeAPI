package utils

import (
	"fmt"
	"runtime"
)

func PrintLogo(searchType string) {
	switch runtime.GOOS {
	case "darwin":
		colorPrint(searchType)
	case "windows":
		textPrint(searchType)
	case "linux":
		colorPrint(searchType)
	}
}

func colorPrint(searchType string) {
	if searchType == "fofa" {
		fmt.Println(Blue("___________     _____           _____ __________.___ \n\\_" +
			"   _____/____/ ____\\____     /  _  \\\\______   \\   |\n |    __)/  _ \\   " +
			"__\\\\__  \\   /  /_\\  \\|     ___/   |\n |     \\(  <_> )  |   / __ \\_/   " +
			" |    \\    |   |   |\n \\___  / \\____/|__|  (____  /\\____|__  /____|   |__" +
			"_|\n     \\/                   \\/         \\/           "))
	} else if searchType == "quake" {
		fmt.Println(Blue("________                __             _____ __________.___ " +
			"\n\\_____  \\  __ _______  |  | __ ____   /  _  \\\\______   \\   |\n /  " +
			"/ \\  \\|  |  \\__  \\ |  |/ // __ \\ /  /_\\  \\|     ___/   |\n/   \\_/. " +
			" \\  |  // __ \\|    <\\  ___//    |    \\    |   |   |\n\\_____\\ \\_/____/" +
			"(____  /__|_ \\\\___  >____|__  /____|   |___|\n       \\__>          \\/    " +
			" \\/    \\/        \\/              "))
	}
	fmt.Println(Red("360 Quake API") + "   " + Green("Author:4ra1n"))
}

func textPrint(searchType string) {
	if searchType == "fofa" {
		fmt.Println("___________     _____           _____ __________.___ \n\\_" +
			"   _____/____/ ____\\____     /  _  \\\\______   \\   |\n |    __)/  _ \\   " +
			"__\\\\__  \\   /  /_\\  \\|     ___/   |\n |     \\(  <_> )  |   / __ \\_/   " +
			" |    \\    |   |   |\n \\___  / \\____/|__|  (____  /\\____|__  /____|   |__" +
			"_|\n     \\/                   \\/         \\/           ")
	} else if searchType == "quake" {
		fmt.Println("________                __             _____ __________.___ " +
			"\n\\_____  \\  __ _______  |  | __ ____   /  _  \\\\______   \\   |\n /  " +
			"/ \\  \\|  |  \\__  \\ |  |/ // __ \\ /  /_\\  \\|     ___/   |\n/   \\_/. " +
			" \\  |  // __ \\|    <\\  ___//    |    \\    |   |   |\n\\_____\\ \\_/____/" +
			"(____  /__|_ \\\\___  >____|__  /____|   |___|\n       \\__>          \\/    " +
			" \\/    \\/        \\/              ")
	}
	fmt.Println("Fofa API   Author:4ra1n")
}
