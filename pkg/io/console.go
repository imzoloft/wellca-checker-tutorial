package io

import "fmt"

func ClearConsole() {
	fmt.Print("\033[H\033[2J")
}
