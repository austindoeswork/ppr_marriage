package main

import ("fmt")


func main(){
	//parse first and generate a ROOMS map and TENANTS map
	ROOMS,TENANTS := Parse("csvs/")
	fmt.Println(TENANTS)
	fmt.Println(ROOMS)
	Marry(ROOMS,TENANTS)
	fmt.Println(TENANTS)
	fmt.Println(ROOMS)
	for num,room := range ROOMS{
		fmt.Print("ROOM: ")
		fmt.Println(num)
		room.PrettyPrint()
	}

}