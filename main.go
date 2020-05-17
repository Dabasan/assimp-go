package main

import (
	"fmt"

	"github.com/dabasan/assimp-go/assimp"
)

func showInfo() {
	major := assimp.GetVersionMajor()
	minor := assimp.GetVersionMinor()
	revision := assimp.GetVersionRevision()
	fmt.Printf("Version: %v.%v.%v\n\n", major, minor, revision)

	legal := assimp.GetLegalString()
	fmt.Println(legal)
}

func main() {
	showInfo()
}
