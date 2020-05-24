package main

import (
	"fmt"

	"github.com/dabasan/assimp-go/assimp"
)

func showAssimpInfo() {
	major := assimp.GetVersionMajor()
	minor := assimp.GetVersionMinor()
	revision := assimp.GetVersionRevision()
	fmt.Printf("Version: %v.%v.%v\n", major, minor, revision)
	fmt.Println("==============================")

	legal := assimp.GetLegalString()
	fmt.Println(legal)
	fmt.Println("==============================")
}
func showModelInfo(model_filename string) {
	meshes, err := assimp.ParseFile(model_filename)
	if err != nil {
		panic(err)
	}

	for i, mesh := range meshes {
		fmt.Printf("[Mesh %v]\n", i)
		fmt.Printf("Face num: %v\n", mesh.FaceCount)
	}
	fmt.Println("==============================")
}

func main() {
	showAssimpInfo()
	showModelInfo("model.obj")
}
