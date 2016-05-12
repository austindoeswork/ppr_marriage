package main

import ("fmt")


func fscale(total int, needed int, floor int, ROOMS map[int]Room) {

	ftotal := total - (floor*len(ROOMS))
	fneeded := needed - (floor*len(ROOMS))
	var fratio float64 = (float64(fneeded)/float64(ftotal))


	// var ratio float64 = (float64(needed)/float64(total))
	var finalTotal float64 = 0
	for name,room := range ROOMS {
		fmt.Print(name)
		fmt.Print(": ")
		temp := (float64(room.Dowry)-float64(floor))*fratio + float64(floor)
		finalTotal = finalTotal + temp
		fmt.Println(temp)
	}
	fmt.Println(finalTotal)
}


func main(){
	//parse first and generate a ROOMS map and TENANTS map
	ROOMS,TENANTS := Parse("csvs/")

	Marry(ROOMS,TENANTS)
	// fmt.Println(TENANTS)
	for name,tenant := range TENANTS{
		fmt.Print("TENANT: ")
		fmt.Println(name)
		fmt.Print("total: ")
		fmt.Println(tenant.Total)
		fmt.Println(tenant.Prefs)
		
	}
	// fmt.Println(ROOMS)
	for name,tenant := range TENANTS{
		fmt.Print("TENANT: ")
		fmt.Println(name)
		fmt.Print("got his ")
		fmt.Print(tenant.Current)
		fmt.Println(" choice.")
		
	}
	totalOffer := 0
	for num,room := range ROOMS{
		fmt.Print("ROOM: ")
		fmt.Println(num)
		room.PrettyPrint()
		totalOffer += room.Dowry
	}
	fmt.Print("\nTOTAL RENT OFFERED: ")

		fmt.Println(totalOffer)

	fscale(totalOffer, 3420, 300, ROOMS)
	fscale(totalOffer, 3420, 0, ROOMS)
}