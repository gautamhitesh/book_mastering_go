//error messages : in built and custom
package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func main() {
	//check command line arguments length
	if len(os.Args)==1{
		fmt.Println("Please give one more or arguments")
		os.Exit(1)
	}

	arguments:=os.Args
	var err error = errors.New("An error")
	k:=1
	var n float64

	for err!=nil{ //continue till gotten error
		if k>=len(arguments){ //exit when all arguments are finished and you get error for all
			fmt.Println("None of the arguments are float")
			return
		}
		n, err=strconv.ParseFloat(arguments[k], 64)    //check for float
		k++
	}
	min, max:= n,n
	for i:=2; i<len(arguments); i++ {
		n, err:= strconv.ParseFloat(arguments[i], 64)
		if err == nil{
			if n < min{
				min = n
			}
			if n > max{
				max=n
			}
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}