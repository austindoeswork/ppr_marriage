// lib.go
package main

import ("fmt"
)

func Hello() {
	fmt.Println("yo nigga")
}

type TenantPref struct {
	Room int
	Cost int
}

type RoomPref struct {
	Name string
	Cost int
}

type Tenant struct {
	Current int
	Prefs []TenantPref
}

type Room struct {
	Current string 
	Dowry int
	Prefs []RoomPref
}

type ByValue []RoomPref

func (s ByValue) Len() int {
    return len(s)
}

func (s ByValue) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByValue) Less(i, j int) bool {
    return s[i].Cost > s[j].Cost
}