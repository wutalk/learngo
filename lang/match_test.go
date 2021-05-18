package lang_test

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestMatch(t *testing.T) {
	fmt.Println("start test match")
	neSwID := "SBTS20A_ENB_0000_000626_000000"
	bomPath := fmt.Sprintf("files/BOM/FastPass/%s/glbfs/20200506.152539/%s.json", neSwID, neSwID)
	bomPattern := fmt.Sprintf("files/BOM/*/%s/*/*/%s.json", neSwID, neSwID)

	bomFileJSONFound, _ := filepath.Match(bomPattern, bomPath)
	if bomFileJSONFound {
		fmt.Println("found bom json")
	} else {
		fmt.Println("not found bom json")
	}
}

func TestSliceIsOrdered(t *testing.T) {
	l := make([]string, 0, 10)
	for i := 0; i < 5; i++ {
		l = append(l, fmt.Sprintf("apple#%d", i))
	}

	for k, v := range l {
		fmt.Println(k, v)
	}

	fmt.Println("after appended")
	l = append(l, "banana#1")
	l = append(l, "banana#2")
	for k, v := range l {
		fmt.Println(k, v)
	}
}
