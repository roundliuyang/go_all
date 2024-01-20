package filepath

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestFibList(t *testing.T) {

	// Calling the Join() function
	fmt.Println(filepath.Join("G", "F", "G"))
	fmt.Println(filepath.Join("G/F", "G"))
	fmt.Println(filepath.Join("gfg", "GFG"))
	fmt.Println(filepath.Join("Geeks", "for", "Geeks"))
	/*
		输出：
			G/F/G
			G/F/G
			gfg/GFG
			Geeks/for/Geeks
	*/
}
