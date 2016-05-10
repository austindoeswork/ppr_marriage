// marry.go
package main

import (
	"fmt"
)

func willMarry(offer int, room int, ROOMS map[int]Room) boolean {
	if ROOMS[room].Dowry < offer {
		return true
	} else {
		return false
	}

}

func findBestRoom(name string, ROOMS map[int]Room, TENANTS map[string]Tenant) {
	currentBest := TENANTS[name].Current
	desiredRoom := TENANTS[name].Prefs[current].Room
	offer := TENANTS[name].Prefs[current].Cost
	if willMarry(c)
}

func Marry(ROOMS map[int]Room, TENANTS map[string]Tenant) {
	fmt.Println("running marriage")
}