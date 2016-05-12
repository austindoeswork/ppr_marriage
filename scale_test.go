package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func toRooms(in map[int]int) map[int]Room {

	rooms := make(map[int]Room)
	id := 0

	for i := range in {
		for j := 0; j < in[i]; j++ {
			rooms[id] = Room{Current: strconv.Itoa(id), Dowry: i}
			id++
		}
	}

	return rooms
}

func getTotal(rooms map[int]Room) int {
	total := 0
	for i := range rooms {
		total += rooms[i].Dowry
	}

	return total
}

func TestScale(t *testing.T) {
	tst := make(map[int]int)
	tst[380] = 8
	tst[450] = 1

	tests := []struct {
		prices map[int]int
		floor  int
	}{
		{
			prices: map[int]int{380: 8, 450: 1},
			floor:  300,
		},
		{
			prices: map[int]int{300: 8, 700: 1},
			floor:  300,
		},
	}

	for i := range tests {
		rooms := toRooms(tests[i].prices)

		assert.Len(t, rooms, 9)

		total := getTotal(rooms)
		adj := fscale(total, 3420, tests[i].floor, rooms)

		assert.True(t, math.Abs(float64(getTotal(adj)-3420)) < 20)

	}
}
