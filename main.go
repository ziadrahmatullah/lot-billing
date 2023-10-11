// Name : MUhammad Ziad Rahmatullah
// Date : Okt 11, 19:32

package main

import (
	"fmt"
)

var vehicles = map[string][]int{
	"motorcycle" : {3000, 2000, 20000},
	"car" : {7000, 5000, 50000},
}

func calculateBilling(hours int, vehicle string) (result int) {
	if hours>=0 {
		result += vehicles[vehicle][0]
		if  hours > 1{
			result += (hours-1) * vehicles[vehicle][1]
			if hours > 24{
				result += vehicles[vehicle][2]
			}
		}
	}
	return
}

func errorHandlingVehicle(vehicle string) bool{
	if _, ok := vehicles[vehicle]; ok{
		return true
	} else {return false}
}


func main(){
	var vehicle string
	var hours int
	
	for !errorHandlingVehicle(vehicle){
		if vehicle != ""{	
			fmt.Println("Your vehicle is undifiend, Try again")	
		}
		fmt.Print("Enter the vehicle : ")
		fmt.Scanln(&vehicle)
	}

	fmt.Print("How many parking hour : ")
	fmt.Scan(&hours)

	tax := calculateBilling(hours, vehicle)
	fmt.Printf("This is your billing : ", tax)
}