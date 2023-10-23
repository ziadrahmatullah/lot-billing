package restaurant

import (
	CustomErr "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/lot-billing/errors"
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

func isAvailableVehicle(vehicle string) (bool, error){
	if _, ok := vehicles[vehicle]; ok{
		return true, nil
	} else {return false, CustomErr.ErrVehicleIsNotAvailable}
}

func LotBilling(vehicle string, hour int) (int, error){
	if ok, err := isAvailableVehicle(vehicle); !ok{
		return 0, err
	}
	return calculateBilling(hour, vehicle), nil
}