package main

import ("fmt")


func main(){
	//parse first and generate a ROOMS map and TENANTS map
	ROOMS,TENANTS := Parse("csvs/")
	// fmt.Println(TENANTS)
	for name,tenant := range TENANTS{
		fmt.Print("TENANT: ")
		fmt.Println(name)
		fmt.Println(tenant.Prefs)
		
	}
	// fmt.Println(ROOMS)
	Marry(ROOMS,TENANTS)
	// fmt.Println(TENANTS)
	for name,tenant := range TENANTS{
		fmt.Print("TENANT: ")
		fmt.Println(name)
		fmt.Println(tenant.Prefs)
		
	}
	fmt.Println(ROOMS)
	for name,tenant := range TENANTS{
		fmt.Print("TENANT: ")
		fmt.Println(name)
		fmt.Print("got his ")
		fmt.Print(tenant.Current)
		fmt.Println(" choice.")
		
	}
	for num,room := range ROOMS{
		fmt.Print("ROOM: ")
		fmt.Println(num)
		room.PrettyPrint()
	}

}