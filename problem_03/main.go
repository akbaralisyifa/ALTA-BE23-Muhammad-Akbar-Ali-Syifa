package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	sort.Ints(dragonHead)
	sort.Ints(knightHeight)

	totalHeight := 0;
	k := 0;

	for _, head := range dragonHead{
		for k < len(knightHeight) && head > knightHeight[k]{
			k++
		}
		if k == len(knightHeight) {
			fmt.Println("knight fall")
			return
		}
		totalHeight += knightHeight[k]
		k++
	}

	fmt.Println(totalHeight)
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	DragonOfLoowater([]int{5, 10}, []int{5})
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})
}