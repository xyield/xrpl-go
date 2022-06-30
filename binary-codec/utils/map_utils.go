package utils

import (
	"fmt"
	"sort"
)

func SortMapByValue(vmap map[string]int) (sortedKeys []string, sortedValues []int, err error) {

	keys := make([]string, 0, len(vmap))

	for key := range vmap {
		keys = append(keys, key)
	}

	fmt.Println("Before Sorting", keys)

	sort.SliceStable(keys, func(i, j int) bool {
		return vmap[keys[i]] < vmap[keys[j]]
	})

	fmt.Println("After Sorting", keys)

	values := make([]int, 0, len(vmap))

	for _, value := range vmap {
		values = append(values, value)
	}

	fmt.Println("Before Sorting", values)

	sort.SliceStable(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println("After Sorting", values)

	return keys, values, nil
}
