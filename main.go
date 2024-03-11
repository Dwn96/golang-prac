package main

import (	
	"errors"
	"godemo/iterations"
	"godemo/person"
)


func intDivision(numerator int, denominator int) (int, int, error) {
	if(denominator == 0) {
		return 0,0, errors.New("Division by zero encountered")
	}
	val := numerator/denominator
	remainder := numerator % denominator
	return val, remainder, nil
}


func main () {
	iterations.MapIterate()
	iterations.ArrIterate()
	person.GetFakeUser()
}