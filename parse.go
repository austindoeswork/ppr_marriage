// parse.go
package main

import (
	"bytes"
	"encoding/csv"
	// "fmt"
	"log"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"strconv"
)

func sortRooms(ROOMS map[int]Room) {
	for _,room := range ROOMS {	
		sort.Sort(ByValue(room.Prefs))
	}
}

func makeRooms(TENANTS map[string]Tenant) (map[int]Room){

	ROOMS := make(map[int]Room)
	for name,tenant := range TENANTS {
		for _,tPref := range tenant.Prefs {
			var rPref RoomPref
			rPref.Name = name
			rPref.Cost = tPref.Cost

			// fmt.Println(tPref.Room)
			// fmt.Println(rPref)

			if _,exists := ROOMS[tPref.Room]; !exists {
				// fmt.Print("making: ")
				// fmt.Println(tPref.Room)
				var room Room
				ROOMS[tPref.Room] = room
			}

			// fmt.Println(ROOMS[tPref.Room].Prefs)
			
			temp := ROOMS[tPref.Room]
			temp.Prefs = append(temp.Prefs, rPref) //hacky workaroud, add pref to room

			ROOMS[tPref.Room] = temp
		}
	}
	return ROOMS
}

func miniParse(path string) (name string, tenant Tenant){

	buf := bytes.NewBuffer(nil)
	f, err := os.Open(path) //open file
	if err != nil {
			log.Fatal(err)
	}
	io.Copy(buf, f) //read file into buffer
	f.Close()

	s := string(buf.Bytes()) //convert bytes to string
	name = strings.Split(path, "/")[5]
	
	// fmt.Println(name)
	// fmt.Println(s)

	r := csv.NewReader(strings.NewReader(s))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var tPref TenantPref
		tPref.Room,_ = strconv.Atoi(record[0])
		tPref.Cost,_ = strconv.Atoi(record[1])
		tenant.Prefs = append(tenant.Prefs, tPref)
	}
	return
}
 
func Parse(path string) (map[int]Room, map[string]Tenant) {
// func Parse(path string) {
	path = "C:/GoStuff/src/ppr_marriage/csvs/"
	files, _ := ioutil.ReadDir(path)

	TENANTS := make(map[string]Tenant)
	// ROOMS := make(map[int]Room)


    for _, f := range files {
            // fmt.Println(f.Name())
    	name, tenant := miniParse(path + f.Name())
    	// fmt.Println(name)
    	// fmt.Println(tenant)
    	TENANTS[name] = tenant
    }
    ROOMS := makeRooms(TENANTS)
    sortRooms(ROOMS)

	// fmt.Println(TENANTS)
	// fmt.Println(ROOMS)
	return ROOMS,TENANTS

}