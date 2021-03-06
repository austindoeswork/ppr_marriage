// marry.go
package main

import (
	"fmt"
)
func marry(name string, offer int, room int, ROOMS map[int]Room, TENANTS map[string]Tenant) {
	oldHusband := ROOMS[room].Current //store old husband to be reassigned
	
	temp := ROOMS[room]
	temp.Current = name
	temp.Dowry = offer
	ROOMS[room] = temp

	if oldHusband != "" {
		tempT := TENANTS[oldHusband]
		tempT.Current++
		TENANTS[oldHusband] = tempT
		findBestRoom(oldHusband, ROOMS, TENANTS)
	}

}

func willMarry(offer int, suitor string, room int, ROOMS map[int]Room, TENANTS map[string]Tenant) (bool) {
	if ROOMS[room].Dowry < offer {
		return true
	} else if  ROOMS[room].Dowry == offer{
		return TENANTS[suitor].Total > TENANTS[ROOMS[room].Current].Total
	} else {
		return false
	}

}

func findBestRoom(name string, ROOMS map[int]Room, TENANTS map[string]Tenant) {
	currentBest := TENANTS[name].Current
	desiredRoom := TENANTS[name].Prefs[currentBest].Room
	offer := TENANTS[name].Prefs[currentBest].Cost

	if willMarry(offer, name, desiredRoom, ROOMS, TENANTS) {	//if they can get married
		marry(name, offer, desiredRoom, ROOMS, TENANTS) //marry them
	} else { //try again with the next best room
		temp := TENANTS[name]
		temp.Current++
		TENANTS[name] = temp

		findBestRoom(name, ROOMS, TENANTS)
	}
}

func Marry(ROOMS map[int]Room, TENANTS map[string]Tenant) {
	fmt.Println("running marriage")
	for name,_ := range TENANTS {
		findBestRoom(name,ROOMS,TENANTS)
	}
}