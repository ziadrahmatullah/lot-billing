package main

import (
	"fmt"
	"errors"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/lot-billing/restaurant"
	CustomErr "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/lot-billing/errors"
)

func main(){
	result, err := restaurant.LotBilling("car", 8)
	if err != nil{
		switch{
		case errors.Is(err, CustomErr.ErrVehicleIsNotAvailable):
			fmt.Printf("there is some error: %s\n", err)
		default:
			fmt.Printf("unexpected error : %s \n", err)
		}
	}
	fmt.Printf("This is your bill : %d\n", result)
}
