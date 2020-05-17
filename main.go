package main

import (
	"fmt"

	"github.com/dabasan/assimp-go/assimp"
)

func main() {
	major := assimp.GetVersionMajor()
	minor := assimp.GetVersionMinor()
	revision := assimp.GetVersionRevision()
	fmt.Printf("%v.%v.%v\n", major, minor, revision)
}
