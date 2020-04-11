// https://www.hackerrank.com/challenges/flatland-space-stations/problem

package main

import (
	"fmt"
	"math"
	"sort"
)

func nearestDistance(loc []int32, pos int32) int32 {
	var prevDistance = int32(0)
	var currentDistance = int32(0)
	var minDistance = int32(0)

	for idx, location := range loc {
		currentDistance = int32(math.Abs(float64(location) - float64(pos)))

		if idx == 0 || currentDistance < prevDistance {
			minDistance = currentDistance
		} else {
			break
		}

		prevDistance = currentDistance
	}

	return minDistance
}

func flatlandSpaceStations(n int32, c []int32) int32 {
	if n == int32(len(c)) {
		return 0
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	var maxDistance = int32(-1)

	for i := int32(0); i < n; i++ {
		var distance = nearestDistance(c, i)

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func main() {
	fmt.Println(flatlandSpaceStations(6, []int32{0, 1, 2, 4, 3, 5,}))
	fmt.Println(flatlandSpaceStations(5, []int32{0, 4}))
}
