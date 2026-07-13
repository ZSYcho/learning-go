package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	/*if err := os.Mkdir("Downloads", 0755); err != nil {
		log.Fatal(err)
	}*/

	dir := "Downloads/static/images"
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}

	// RemoveAll removes path and any children it contains, not the path itself
	if err := os.RemoveAll("Downloads"); err != nil {
		log.Fatal(err)
		// we just delete the deepest dir
		// maybe we need recursion
	}
}
