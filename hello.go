package main

import ("fmt")


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

}