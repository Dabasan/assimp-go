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
func showModelInfo(model_filename string) {
	meshes, err := assimp.ParseFile(model_filename)
	if err != nil {
		panic(err)
	}

	for i, mesh := range meshes {
		fmt.Printf("[Mesh %v]\n", i)
		fmt.Printf("Vertex num: %v\n", mesh.VertexCount)
	}
}

func main() {
	//showInfo()
	showModelInfo("model.x")
}
