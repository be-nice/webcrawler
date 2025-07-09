package pkg

import (
	"fmt"
	"sort"
)

func PrintOutput(pages map[string]int) {
	pageSlice := []KeyVal{}

	for k, v := range pages {
		pageSlice = append(pageSlice, KeyVal{k, v})
	}

	sort.Slice(pageSlice, func(i, j int) bool {
		return pageSlice[i].v > pageSlice[j].v
	})

	for _, page := range pageSlice {
		fmt.Printf("%s | Visit count %d\n", page.k, page.v)
	}
}
