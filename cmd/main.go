package main

import (
	"fmt"
	utils "github.com/ozoncp/ocp-response-api/internal"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	mappa := map[int]int{
		0: 1,
		1: 1,
		2: 3,
	}

	fmt.Println(utils.ChunkInt(slice, 3))
	fmt.Println(utils.ReverseIntInt(mappa))
}
